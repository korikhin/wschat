# WebSocket Chat

For _now_ its just a simple echo server with ~~Web~~Sockets.

## Usage

First start the server:

```sh
go run . server
```

Then launch the client in another terminal:

```sh
go run . client
```

After that you can enter a message and hit enter on the client side and the server will echo the message. For example:

```
      Client      |   Server
------------------|------------
 > Hello          | > Hello
 > Echo: Hello    |
```
