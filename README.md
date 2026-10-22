# reinchat

> A self-hosted chat harness for LLMs, written in Go, on top of any OpenAI-compatible API.

![CI](https://github.com/Naw3/reinchat/actions/workflows/ci.yml/badge.svg)
![Go version](https://img.shields.io/github/go-mod/go-version/Naw3/reinchat)
![License](https://img.shields.io/github/license/Naw3/reinchat)
![Status](https://img.shields.io/badge/status-work%20in%20progress-orange)

**Status:** early development.

## Why

Testing a model the way people actually use it takes more than a raw API call: a chat UI, system prompts,
skills and tool calling. reinchat provides that layer as a single self-hosted binary. It talks only to the
standard OpenAI-compatible API, so the same harness works with Ollama, llama.cpp, vLLM, Unsloth Studio or any
custom model server, local or remote.

## Features

What works today:

- **Chat Completions client**: sends a message to any OpenAI-compatible endpoint (`/v1/chat/completions`) and
  prints the response, from the command line.

## Quick start

### Requirements

- Go 1.27 or later
- An OpenAI-compatible endpoint, for example [Ollama](https://ollama.com) running locally

### Run

```bash
git clone https://github.com/Naw3/reinchat.git
cd reinchat
```

Create a `.env.dev` file at the root of the repository with the variables described in
[Configuration](#configuration), then:

```bash
go run ./cmd/reinchat
```

The program sends a test message to the model and prints the raw JSON response.

## Configuration

reinchat reads three variables from `.env.dev` (the file is required for now):

- **`BASE_URL`**: the base URL of your backend's OpenAI-compatible API, up to and including `/v1`, without
  `/chat/completions`. It depends on the backend and the port it listens on: check its documentation or its
  settings. For example, Ollama listens on `http://localhost:11434/v1` by default.
- **`API_KEY`**: the key your backend expects, sent as a `Bearer` token. Leave it empty if the backend needs no key
  (local servers usually do not).
- **`MODEL`**: the model name exactly as your backend reports it. Most backends list the available names at
  `<BASE_URL>/models`; the format differs from one backend to another (for example `name:tag` on Ollama, a
  repository path on others).

## Development

```bash
go build ./...
go vet ./...
go test -race ./...
```

The same checks run in CI on every push and pull request. The race detector (`-race`) needs cgo and a C
compiler; without one, run `go test ./...`.

## License

[MIT](LICENSE)
