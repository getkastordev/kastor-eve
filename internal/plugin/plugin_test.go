package plugin

import (
	"context"
	"testing"

	protocol "github.com/weirdGuy/kastor/protocol/v1"
)

func TestGenerateModuleWithoutAgents(t *testing.T) {
	response, err := (Handler{}).Generate(context.Background(), &protocol.GenerateRequest{
		Module: &protocol.Module{},
		Target: &protocol.Target{Name: "typescript", Type: "codegen"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(response.Files) != 0 {
		t.Fatalf("files = %#v", response.Files)
	}
}

func TestValidateRejectsStdio(t *testing.T) {
	response, err := (Handler{}).Validate(context.Background(), &protocol.ValidateRequest{
		Target: &protocol.Target{Name: "typescript", Type: "codegen"},
		Module: &protocol.Module{
			Tools:      []*protocol.Tool{{Name: "fetch", Source: &protocol.ToolSource{Kind: "mcp", URI: "mcp://fetch/get"}}},
			MCPServers: []*protocol.MCPServer{{Name: "fetch", Transport: "stdio", Command: "uvx"}},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(response.Diagnostics) != 1 || response.Diagnostics[0].Addr != "mcp_server.fetch" {
		t.Fatalf("diagnostics = %#v", response.Diagnostics)
	}
}
