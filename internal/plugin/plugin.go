package plugin

import (
	"context"
	"fmt"

	"github.com/getkastordev/kastor-eve/internal/build"
	"github.com/getkastordev/kastor-eve/internal/eve"
	protocol "github.com/weirdGuy/kastor/protocol/v1"
)

const Source = "github.com/getkastordev/kastor-eve"

var Version = "0.1.0-dev"

type Handler struct{}

func (Handler) Metadata(context.Context) (protocol.Metadata, error) {
	return protocol.Metadata{
		Protocol: protocol.Version,
		Source:   Source,
		Version:  Version,
		Kinds:    []protocol.Kind{protocol.KindCodegen},
		Capabilities: protocol.Capabilities{
			CredentialSchemes: []string{protocol.SchemeEnv},
		},
	}, nil
}

func (Handler) Validate(_ context.Context, request *protocol.ValidateRequest) (*protocol.ValidateResponse, error) {
	if request == nil || request.Module == nil || request.Target == nil {
		return nil, fmt.Errorf("eve: validate request is incomplete")
	}
	var diagnostics []protocol.Diagnostic
	if request.Target.Type != string(protocol.KindCodegen) {
		diagnostics = append(diagnostics, diagnostic(request.Target.Addr(), "target must have type codegen", ""))
	}
	for key := range request.Target.Config {
		diagnostics = append(diagnostics, diagnostic(request.Target.Addr(), "unsupported config attribute", key))
	}
	referenced := referencedServers(request.Module)
	for _, server := range request.Module.MCPServers {
		if !referenced[server.Name] {
			continue
		}
		if server.Transport == "stdio" {
			diagnostics = append(diagnostics, diagnostic(server.Addr(), "stdio transport is not supported", "Eve connections require an HTTP endpoint"))
		}
		auth, ok := server.AuthFor(request.Target.Addr())
		if !ok {
			continue
		}
		scheme, _, err := protocol.ParseCredentialRef(auth.Ref)
		if err == nil && scheme != protocol.SchemeEnv {
			diagnostics = append(diagnostics, diagnostic(server.Addr(), "unsupported credential scheme", scheme+"://"))
		}
	}
	return &protocol.ValidateResponse{Diagnostics: diagnostics}, nil
}

func (Handler) Generate(_ context.Context, request *protocol.GenerateRequest) (*protocol.GenerateResponse, error) {
	if request == nil || request.Module == nil || request.Target == nil {
		return nil, fmt.Errorf("eve: generate request is incomplete")
	}
	files, err := (eve.Generator{}).Generate(&build.Job{Module: request.Module, Target: request.Target})
	if err != nil {
		return nil, err
	}
	return &protocol.GenerateResponse{Files: files}, nil
}

func diagnostic(addr, summary, detail string) protocol.Diagnostic {
	return protocol.Diagnostic{Severity: protocol.SeverityError, Addr: addr, Summary: summary, Detail: detail}
}

func referencedServers(module *protocol.Module) map[string]bool {
	referenced := map[string]bool{}
	for _, tool := range module.Tools {
		if tool.Source == nil || tool.Source.Kind != "mcp" {
			continue
		}
		server, _, err := protocol.ParseMCPURI(tool.Source.URI)
		if err == nil {
			referenced[server] = true
		}
	}
	return referenced
}
