# stringsvc5

This example demonstrates how to use Go kit to implement a [Gin](https://github.com/gin-gonic/gin) HTTP service.

Run the Gin HTTP service:

```bash
# start http service on :8080 port
go run .
```

Request to server:

```bash
curl -XPOST -d '{"s": "Go-Kit"}' http://localhost:8080/uppercase
# {"v":"GO-KIT"}

curl -XPOST -d '{"s": "Go-Kit"}' http://localhost:8080/count
# {"v":6}
```
