

## Minimal Golang Project

[![Go Report Card](https://goreportcard.com/badge/github.com/blankdots/minimalgo)](https://goreportcard.com/report/github.com/blankdots/minimalgo)
![Go Unit Tests](https://github.com/blankdots/minimalgo/workflows/Go/badge.svg)
![Go style check](https://github.com/blankdots/minimalgo/workflows/Multilinters/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/blankdots/minimalgo/badge.svg?branch=main)](https://coveralls.io/github/blankdots/minimalgo?branch=main)

An example of a minimal Golang project that contains an web application built with [fiber](https://github.com/gofiber/fiber).

At the same time the project exemplifies:
* fiber server;
* logging formatting;
* unit tests;
* github-actions and [coveralls](https://coveralls.io/github/blankdots/minimalgo) integration;


### Install and Run

Installation can be done:
* cloning repository:
```
$ git clone git@github.com:blankdots/minimalgo.git
$ cd minimalgo
$ pip install .
```

After install the application can be started like: `$ minimalgo`

#### Tests and Documentation

In order to run the tests: `$ go test -v ./...` in the root directory of the git project.

###  Structure

```
.
├── api
│   ├── api.go
│   └── health_test.go
├── coverage.txt
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── minimalgo
├── README.md
└── utils
    └── utils.go
```

### License

`minimal` go and all it sources are released under `Apache License 2.0`.
