# ascii-art-go
ASCII Art in Golang

_Future improvements can be found marked as **TODO** comments for now. In the future, move them to github issues._

# How to run

Once you run, the service will be available on port 8080 in your localhost. To test, I recommend using postman. Send a request to `localhost:8080/artify` with a form-data containing an image in the `file` key to get the art.

## Docker compose

The easiest way to run is by using docker compose. It will spin up the service in port 8080 and also run API test collection against the running service, producing the json report in `resources/test/newman/report`.

```bash
docker-compose up
```

## Running the binaries (Linux and MacOS)

Alternatively Look for the files in `bin`. Run the one corresponding to your environment.

```bash
./ascii-art-<os>
```

# Scaling considerations

## Independence from cloud provider

Since the api is containerized, it could be ported to another cloud provider, specially if you're using Kubernetes. By creating a helm chart, the deployment process would be even smoother and allow for a GitOps approach.

## How to make it more scalable 

Currently, the solution uses multithreading to improve performance while traversing the image pixel space. The downside of using this approach is that we rely on memory to retain the canvas that would be returned later on. One way to overcome this would be to use gRPC to stream the "pixels" (position and character) to the client as they were ready - each client would build its own canvas space, reducing the memory cost from the server.

## Risks of upgrading the approach

Although the usage of design patterns would provide a more adaptable code, currently there's little unit test coverage - paramount to generate safety for refactoring. Also, I'd remove the endpoint implementation from the main.go, os that reading the routes and understanding the general api structure would be easier (therefore, more maintainable).