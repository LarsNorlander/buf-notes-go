# Example Buf Server/Client

For this example to work, you'll need to first generate some code through buf, then you could build the example 
server and client.

```shell
make buf-generate
```

You can then start run `cmd/server/main.go` and use a tool such as Postman to populate the list of notes. The server
supports Reflection to make using tools easier.

Running `cmd/client/main.go` will print out all the notes to the terminal.