# Health Check Tester

This is a simple container which can be used when testing other applications using [TestContainers](https://testcontainers.com).

This container will start up and respond to requests on the following endpoints:
- `/healthz`
- `/readyz`
- `/livez`

