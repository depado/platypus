<h1 align="center">Platypus</h1>
<h2 align="center">
  <img src="img/platypus.jpg" alt="mascot" height="350px">

  [![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)[![forthebadge](https://forthebadge.com/images/badges/uses-badges.svg)](https://forthebadge.com)

  ![Go Version](https://img.shields.io/badge/Go%20Version-latest-brightgreen.svg)
  [![Go Report Card](https://goreportcard.com/badge/github.com/Depado/platypus)](https://goreportcard.com/report/github.com/Depado/platypus)
  [![Build Status](https://drone.depa.do/api/badges/Depado/platypus/status.svg)](https://drone.depa.do/Depado/platypus)
  [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Depado/platypus/blob/master/LICENSE)
  [![Say Thanks!](https://img.shields.io/badge/Say%20Thanks-!-1EAEDB.svg)](https://saythanks.io/to/Depado)

  Very simple mock server that doesn't do much
</h2>

> That external API you're relying on is supposed to send back a duck.
> Unfortunately sometimes things don't go as expected and you're dealing with a
> platypus instead. It's ok.

## Introduction

Platypus is a very simple mock server to abstract external services. It supports
CORS which is disabled by default but fully configurable.

## Install

### From source

You can install platypus by using directly the go command line:

```bash
$ go install github.com/Depado/platypus
```

Or if you wish to have the version information and such:

```
$ git clone git@github.com:Depado/platypus.git
$ make install
```

### Binary release

You can also download the latest release for your system on the
[release page](https://github.com/Depado/platypus/releases).

### Docker image

You can run platypus directly using docker using the packages found
[on this page](https://github.com/Depado/platypus/pkgs/container/platypus).

Example:

```bash
docker run -it --rm -v (pwd):/data ghcr.io/depado/platypus:latest --mock mocks/gerrit.yml
```

## Usage

<h2 align="center">
  <img src="img/screenshot.png" alt="screenshot" height="400px">
</h2>

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

When starting up, platypus will display the endpoints with the associated
method(s) as well as the response information (type of response, code, name,
etc)

## Configuration

### mock.yml

To define how platypus should behave you should create a `mock.yml` file where
you're going to start platypus. Alternatively you can place the yaml file
anywhere and name it as you like if you provide the `--mock <path>` flag when
starting platypus.

This file contains the various definitions of your endpoints like so:

```yaml
---
endpoints:
- path: /hello/world
  get:
    responses:
    - code: 200
      body: |
        {"hello": "world"}

- path: /with/parameter/:param
  get:
    responses:
    - code: 200
      body: |
        {"hello": "world"}
```

Each endpoint is defined by a string which can accept query params in the
same form as [gin](https://github.com/gin-gonic/gin). Then you can define the
HTTP methods it should handle (`get`, `post`, `put`, `patch`, `delete`, `head`,
`options`). Each of these method can have a specific behavior.

## Methods

Each method associated to an endpoint can be customized to add validation
behavior or to output more information when it gets called.

```yaml
endpoints:
- path: /hello/world
  get:
    dump: [headers,body,host]
    validate:
      headers:
        present: [Authorization]
        absent: [Forbidden, X-Forbidden]
        match:
          X-Exact-Value: "exact value"
      body:
        contains: "interesting value"
   responses:
    - code: 200
      body: |
        {"hello": "world"}
```

### Dump

`dump` defines what should be shown in the terminal when a request calls this
specific endpoint. Values can be `header`, `body`, `host`, `proto` or any
combination of these. The special option `all` can also be used to group all
these.

<h2 align="center">
  <img src="img/dump.png" alt="screenshot" height="400px">
</h2>

### Validate

The `validate` key can be used to validate incoming requests to match special
rules. This is especially useful to quickly see whether or not a request is
supposed to fail with the non-mocked API. This is entirely optional as it may
not be the desired behavior anyway.

- `headers`
  - `present`: List of headers the incoming request must have
  - `absent`: List of headers the incoming request must not have
  - `match`: Header/Value pairs the incoming request must have
- `body`
  - `contains`:



### Responses

Every method in a path can have multiple responses. These responses are sent
in a random manner when the endpoint is called. This allows to simulate
unexpected behaviors on the mock side. If you wish you can add a `ratio` keyword
which will tell platypus how often this response should be sent back. For
example the following example will send back a `500` error 1% of the time:

```yaml
endpoints:
- path: /failure/maybe
  get:
    responses:
    - code: 200
    - code: 500
      ratio: 1
```

The ratio keyword is a percentage. If no ratio is provided, then the odds are
distributed equally between all the possible responses. So if there's only one
response provided, it will always be sent.

Note that the above yaml is equivalent to this one:

```yaml
endpoints:
- path: /failure/maybe
  get:
    responses:
    - code: 200
      ratio: 99
    - code: 500
      ratio: 1
```

### Headers

An additional field `headers` can be added to a specific method.
This header will then be considered as mandatory and, if absent from the request,
the endpoint will return a `400 Bad Request` with a specific JSON output.

Note that this check will happen before picking up the response, so it will
always return this error before returning any other response.

Headers in the responses can be defined directly in the response like so:

```yaml
endpoints:
- path: /failure/maybe
  get:
    responses:
    - code: 200
      headers:
      - name: Content-Type
        value: application/json; charset=utf-8
    - code: 500
      ratio: 1
```

To make things simpler regarding headers, a specific key `preset` can be used.
The `json` preset will set the `Content-Type` header with the correct value. So
the example above can be rewritten like so:

```yaml
endpoints:
- path: /failure/maybe
  get:
    responses:
    - code: 200
      preset: json
    - code: 500
      ratio: 1
```

### Empty method endpoint

If a method is defined within an endpoint but doesn't have responses or anything
attached, then a `200` will be sent back. This can be achieved like so:

```yaml
- path: /maybe
  post: {}
  patch: {}
  get:
    responses:
    - code: 200
      preset: json
    - code: 500
      ratio: 1
```

When starting up, platypus will be able to generate two simple endpoints for the
POST and PATCH methods. Both will simply return a `200` and return.

## Attributions

Platypus drawing by
[rawpixel](https://www.rawpixel.com/image/562293/duck-billed-platypus-shade-drawing)
under the [CC0 license](https://creativecommons.org/publicdomain/zero/1.0/)