## This is a Golang code examples repo I used for learning.

### TCP Client Server

- Features:

  - starts a TCP server on port 8080 by default
  - prints data the client has sent it
  - responds back to client with timestamped text message confirming data reception
  - the client currently sends a single hardcoded msg and quits, so quite rudimentary to not say useless, but this is just learning. For more advanced client functionality one could just use telnet or netcat instead of course :)
- Installation:
  - go install github.com/tomsucho/go-examples/tcp-client-server/tcp-client@latest
  - go install github.com/tomsucho/go-examples/tcp-client-server/tcp-server@latest
- Usage:
  - the above will build executable and store it under your `$PATH` so you then you can start each with just typing `tcp-server` or `tcp-client`.
