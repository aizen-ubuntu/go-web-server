# go-web-server


## main.go

This is a HTTP server that returns a the HTTP headers from a request.

## secure-server.go

This is a HTTP server that returns the HTTP headers from a request.

### Caveats

- The certificate and private key filenames are harcoded in the source code.
- WARNING: Make sure your .gitignore exclude sensitive files like the key or even better, do not keep the files in your local repo.
- These servers and the whole content of the repo was done for pure testing purposes and with time constraints.