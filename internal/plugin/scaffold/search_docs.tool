tool "search_docs" {
  description = "Search product documentation"

  param "query" {
    type = string
  }

  returns {
    type = string
  }

  source {
    kind = "mcp"
    uri  = "mcp://docs/search"
  }
}
