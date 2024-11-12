# Health Check Tester

This is a simple container which can be used when testing other applications using [TestContainers](https://testcontainers.com).

This container will start up and respond to requests on the following endpoints:
- `/healthz`
- `/readyz`
- `/livez`
- `/testcontainersz`

## Usage
When starting the Docker container, you can specify 3 different environment variables to control the behavior of the container:
- `HEALTHZ=TRUE|FALSE` - If set to `TRUE`, the `/healthz` endpoint will return a 200 status code. If set to `FALSE`, the `/healthz` endpoint will return a 500 status code.
- `READYZ=TRUE|FALSE` - If set to `TRUE`, the `/readyz` endpoint will return a 200 status code. If set to `FALSE`, the `/readyz` endpoint will return a 500 status code.
- `LIVEZ=TRUE|FALSE` - If set to `TRUE`, the `/livez` endpoint will return a 200 status code. If set to `FALSE`, the `/livez` endpoint will return a 500 status code.

The default value for each of these environment variables is `TRUE`.

The `/testcontainersz` endpoint will always respond with a 200 and can be monitored by TestContainers for pod readiness in lieu of the healthz endpoint.

```golang
req := testcontainers.ContainerRequest{
    Image:        "healthcheck-tester:latest",
    ExposedPorts: []string{"80/tcp"},
    WaitingFor:   wait.ForHttp("/testcontainersz"),
}
```
