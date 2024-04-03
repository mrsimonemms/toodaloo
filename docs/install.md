# Installation

This project is written in Golang, so the application is contained in a single
binary.

## Docker

A Docker image is maintained at `ghcr.io/mrsimonemms/toodaloo`. By default, the
`latest` tag is the most recent tag created. You may specify a version, in the
format `vX.Y.Z`.

```shell
docker run \
  -it \
  --rm \
  -v /path/to/dir:/data \
  ghcr.io/mrsimonemms/toodaloo
```

A Docker volume exists at `/data` and the environment variables are configured
to point to that directory.

## Go

```shell
go install github.com/mrsimonemms/toodaloo@latest
```

## GitHub releases

Download the desired binary from the [GitHub releases page](https://github.com/mrsimonemms/toodaloo/releases)
and install to your system.
