# Your Kastor Eve module

This starter was supplied by `kastor-eve` through `kastor new`.

Replace the example MCP URL in `kastor.hcl`, then validate and build:

```sh
kastor validate
kastor build
cd gen/eve/assistant
npm install
npx eve dev
```

Set `AI_GATEWAY_API_KEY` for local model access. Edit the Kastor source and
rebuild instead of editing generated files.
