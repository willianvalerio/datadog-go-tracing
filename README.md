# datadog-go-tracing
This code is an example of manual and auto tracing with Datadog.

To run it, open two terminals and run

`go run main.go`

And in another terminal

`go run another-service/main.go`

Do some requests and in Datadog, notice the traces:

`curl localhost:3001/manual`

`curl localhost:3001/auto`