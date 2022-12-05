# Status

Go wrapper for [Plain Text Status](https://status.plaintext.sh)

## Instructions
1. Clone the repository
2. Run `go build -o status main.go`
3. Move the `status` binary to `/usr/local/bin`

There is also a pre-built binary, which will require going into `Security & Privacy` settings to let it run.

## Usage
```shell
status slack aws pagerduty github
```
If the binary is executed with no arguments, all services will be displayed.