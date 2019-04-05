Docker Executor
===============

Docker Executor is able to execute a command inside a currently running docker using an http endpoint.

The command is defined via the environment variable `DOCKER_EXECUTOR_CMD`

The container in which the command should be run can be defined by partial name using `DOCKER_EXECUTOR_IDENTIFIER`

By default the service runs on port 8080, but this can be changed using `DOCKER_EXECUTOR_PORT`

The output of the command will be returned in the http response.


Example Configuration
=====================

An example configuration can be seen in `Dockerfile` and `Makefile`

This setup will add the binary to a docker image called `docker-executor:latest`

When run using the `Makefile` it will run on port `8080`. When called it will execute a `php -v` on all running containers that have `php` in their name.