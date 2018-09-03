# HTTP Server Timing Example

## Start the server

`go run cmd/server/main.go`

## Run 

Without `Server-Timing` header:

`curl --insecure --raw -v https://localhost:8443`

With `Server-Timing` header:

`curl --insecure --raw -v https://localhost:8443?perf`

## TODO

- Explain the `HTTP Trailers` thingy issue. The `Server-Timing` should be a `Trailer` field instead of a `Header` field, but the browsers (Firefox, Chrome, Edge, etc) does not support trailer fields, and looks there is no plan to support for now.

## How to generate priv and cert for SSL/TLS

### Generate private key (.key)

```bash
# Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out server.key 2048

# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out server.key
```

### Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)

```bash
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```
