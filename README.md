# lifxprotocol-go

This repository provides the [LIFX LAN protocol](https://lan.developer.lifx.com/) definitions and a Go-based code generator for working with LIFX devices on the local network.

It includes:

- `cmd/protocol-gen` — a code generator that produces Go types and packet structures.
- `cmd/protocol-gen/src/protocol.yml` — the upstream protocol definition from the official [LIFX repo](https://github.com/LIFX/public-protocol).
- `cmd/protocol-gen/src/protocol_commit.txt` — the upstream commit the current protocol definition was sourced from.
- `gen/` — the generated Go files ready to be imported by other projects.

## Usage

To update the protocol and regenerate code:

```bash
make regenerate
```

This will:

- Clone the latest upstream lifx-protocol repo.
- Check if the protocol definition has changed.
- If changed, copy the new protocol.yml and commit hash.
- Regenerate Go code under gen/.

### Requirements

- Go 1.24+
- Git
- Internet connection (to fetch latest upstream repo)

### License

This project is MIT licensed. See LICENSE for details.
