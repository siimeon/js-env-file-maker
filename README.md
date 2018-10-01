# Javascript environment file maker

This simple go script creates environment file from operating system environment variables.
Environment file uses `window.env` to store variables.

Example env file with environment `APP_EXAMPLE: "example string"` and `APP_FOO: "bar"` variables

```
window.env = { EXAMPLE: "example string", FOO: "bar" }
```

## How to use

To build linux binary run docker-compose commands

```
docker-compose build build-linux
docker-compose run build-linux
```

or build with your local go by running go build command in root of project

```
go build -o bin/js-env-file-maker
```
