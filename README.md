# Status

Go wrapper for [Plain Text Status](https://status.plaintext.sh)

## Installation
### Build From Source
1. Clone the repository
2. Run `go build -o status main.go`
3. Move the `status` binary to `/usr/local/bin`

### Pre-built Binary
1. [Download the binary from GitHub](https://github.com/cassamajor/status/releases)
2. Decompress the archive
3. Move the `status` binary to `/usr/local/bin`
4. Go into `Privacy & Security` settings to let it run.

## Usage
You can query a single service
```shell
status aws
```
Multiple services
```shell
status slack aws pagerduty github
```
Or print the status of all services
```shell
status
```