# EIP-Server

A small EtherNet/IP (CIP) test server written in Go. It uses the [gologix](https://github.com/danomagnum/gologix) library to start a Logix‑compatible endpoint that exposes a predictable set of tags—ideal for integration tests or client demos.

---

## Quick start (local)

```bash
# clone the repo
git clone https://github.com/united-manufacturing-hub/eip-server && cd eip-server

# run with defaults (booltags=10, inttags=10, path="1,0")
make run

# or build a binary
make build
./eip-server -booltags 32 -inttags 64 -path "1,0"
```

The server listens on **TCP/44818** (the EtherNet/IP standard port). Use any compatible client (e.g., PyLogix, gologix, Studio 5000) to browse the tags.

---

## Run in Docker

```bash
docker run -p 44818:44818 ghcr.io/united-manufacturing-hub/eip-server:latest \
           -booltags 32 -inttags 64 -path "1,0"
```

### docker-compose example

```yaml
eip-server:
  image: ghcr.io/united-manufacturing-hub/eip-server:latest
  ports:
    - "44818:44818"
  command: ["-booltags", "32", "-inttags", "64", "-path", "1,0"]
```

---

## CLI flags

| Flag        | Default | Description                            |
| ----------- | ------- | -------------------------------------- |
| `-booltags` | `10`    | How many `Bool<N>` tags to auto-create |
| `-inttags`  | `10`    | How many `Int<N>` tags to auto-create  |
| `-path`     | `1,0`   | CIP routing path (e.g. slot number)    |

---

## Tags exposed by default

The program creates a deterministic dataset on startup (arrays have length 10 unless noted):

| Name              | Go type   | Logix type | Example value   |
| ----------------- | --------- | ---------- | --------------- |
| `Bool1…BoolN`     | `bool`    | `BOOL`     | `true / false`  |
| `Int1…IntN`       | `int16`   | `INT`      | `123`           |
| `testbyte`        | `byte`    | `SINT`     | `0x01`          |
| `testint8`        | `int8`    | `SINT`     | `-16`           |
| `testint16`       | `int16`   | `INT`      | `3`             |
| `testint32`       | `int32`   | `DINT`     | `12 345`        |
| `testint64`       | `int64`   | `LINT`     | `12 345 678`    |
| `testuint`        | `uint16`  | `UINT`     | `123`           |
| `testuint32`      | `uint32`  | `UDINT`    | `1 234`         |
| `testuint64`      | `uint64`  | `LWORD`    | `1 234 567`     |
| `testfloat32`     | `float32` | `REAL`     | `543.21`        |
| `testfloat64`     | `float64` | `LREAL`    | `10 238.21`     |
| `teststring`      | `string`  | `STRING`   | `"Hello World"` |
| `…array` variants | slices    | `ARRAY[]`  | see code        |

