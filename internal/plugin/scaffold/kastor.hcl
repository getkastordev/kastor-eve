kastor {
  required_plugins {
    eve = {
      source  = "github.com/getkastordev/kastor-eve"
      version = "~> 0.1"
    }
  }
}

model "fast" {
  provider = "anthropic"
  id       = "claude-sonnet-4-5"
}

target "eve" {
  type   = "codegen"
  plugin = "eve"
  output = "./gen/eve"
}

mcp_server "docs" {
  transport = "http"
  url       = "https://example.com/mcp"
}
