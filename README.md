# tryp

tryp is a minimal wrapper for the Google Maps Platform Distance Matrix API.

## Install

```bash
go get github.com/jamiekieranmartin/tryp
```

Otherwise you can download the binary from [Releases](https://github.com/jamiekieranmartin/tryp/releases)

## Usage

### CLI

```bash
tryp -key "my super secret key"
```

### Golang SDK

```go
// make new client
client, err := tryp.NewClient("my super secret key")
if err != nil {
  panic(err)
}

// get distance matrix
response, err := client.Get(tryp.Request{
  Origins:      []string{"Southbank, Brisbane"},
  Destinations: []string{"Fortitude Valley, Brisbane"},
})
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## CLI flags

### `-key`

Google Maps Platform API key. Get one here: https://developers.google.com/maps/documentation/distance-matrix/get-api-key

```bash
trpy -key "my super secret key"
```

### `-config`

Path a custom configuration file. Defaults to `./config.toml`.

```bash
tryp -config "./path/to/my/file.toml"
```

### `-out`

Output to JSON file. Defaults to none.

```bash
trpy -key "my super secret key" -out "./result.json"
```

## TOML Configuration

```toml
# config.toml
key = "my super secret key"

[request]
origins = ["Southbank, Brisbane"]
destinations = ["Fortitude Valley, Brisbane"]
```
