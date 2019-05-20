# platypus
Very simple mock server that doesn't do much

[![Build Status](https://drone.depa.do/api/badges/Depado/platypus/status.svg)](https://drone.depa.do/Depado/platypus)

## Introduction

Platypus is a very simple mock server to abstract external services. It supports
CORS which is disabled by default but fully configurable. Platypus also allows
to respond with weighted responses.

## Usage

```
Platypus is a very simple mock server

Usage:
  platypus [flags]
  platypus [command]

Available Commands:
  help        Help about any command
  version     Show build and version

Flags:
      --conf string                   configuration file to use
  -h, --help                          help for platypus
      --log.format string             one of text or json (default "text")
      --log.level string              one of debug, info, warn, error or fatal (default "info")
      --log.line                      enable filename and line in logs
      --mock string                   file to mock from (default "mock.yml")
      --server.cors.all               defines that all origins are allowed
      --server.cors.enable            enable CORS
      --server.cors.expose strings    array of exposed headers
      --server.cors.headers strings   array of allowed headers (default [Origin,Authorization,Content-Type])
      --server.cors.methods strings   array of allowed method when cors is enabled (default [GET,PUT,POST,DELETE,OPTION,PATCH])
      --server.cors.origins strings   array of allowed origins (overwritten if all is active)
      --server.host string            host on which the server should listen (default "127.0.0.1")
      --server.mode string            server mode can be either 'debug', 'test' or 'release' (default "release")
      --server.port int               port on which the server should listen (default 8080)

Use "platypus [command] --help" for more information about a command.
```