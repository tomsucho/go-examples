## This is a Golang code examples repo I used for learning.

### TCP Client Server

- Features:

  - starts a TCP server on port 8080 by default
  - prints data the client has sent it
  - responds back to client with timestamped text message confirming data reception
  - the client currently sends a single hardcoded msg and quits after getting server ACK message back. When the `-pipe` flag added it alternatively can read data from Linux pipe insteaf of using the hardcoed text message e.g.:
    ```
    11:52 $ echo "Testing!" | tcp-client -pipe
    2024/01/03 11:52:22 Waiting for reply..
    Server response: 2024-01-03T11:52:22.535248794+01:00      ==> SERVER ACKed!
    ```
- Installation:
  - go install github.com/tomsucho/go-examples/tcp-client-server/tcp-client@latest
  - go install github.com/tomsucho/go-examples/tcp-client-server/tcp-server@latest
- Usage:
  - the above will build executable and store it under your `$PATH` so you then you can start each with just typing `tcp-server` or `tcp-client`.
