# serve

![action](https://img.shields.io/endpoint.svg?url=https://actions-badge.atrox.dev/sundowndev/serve/badge?ref=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/sundowndev/serve)](https://goreportcard.com/report/github.com/sundowndev/serve)
![release](https://img.shields.io/github/release/SundownDEV/serve.svg)

Lightweight, simple, production ready HTTP server for static file serving.

## Installation

```
# Download the archive
$ curl -L "https://github.com/sundowndev/serve/releases/download/v1.1.0/serve_$(uname -s)_$(uname -m).tar.gz" -o serve.tar.gz

# Extract the binary
$ tar xfv serve.tar.gz

# Run the software
$ ./serve -h

# You can install it globally
$ mv ./serve /usr/bin/serve
```

If the installation fails, it probably means your OS/arch is not suppored.

Please check the output of echo "$(uname -s)_$(uname -m)" in your terminal and see if it's available on the [GitHub release page](https://github.com/sundowndev/serve/releases).

## Usage

Here's the help message :

```
$ serve -h

Usage of http_server:
  -addr string
    	The address to listen to. Defaults to 0.0.0.0 (default "0.0.0.0")
  -port string
    	The port to listen to. Defaults to 80 (default "80")
  -prefix string
    	The URI prefix path. Default is / (default "/")
  -root string
    	the directory to serve files from. Defaults to the current dir (default ".")
```

Launch a HTTP server to serve files in the current directory :

```
$ serve -root . -port 8080 
```

### Docker image

Launch a HTTP server to serve files in the `public` directory :

```
$ docker run -it --rm -v $PWD/public:/app/public:ro -p 80:80 sundowndev/serve -root public
```

#### Using compose

```yml
version: '3.7'

services:
    serve:
      container_name: serve
      restart: on-failure
      image: sundowndev/serve:latest
      command: "-root ./public"
      environment:
        - GIN_MODE=release
      volumes:
        - ./fixtures:/app/public:ro
      ports:
        - "80:80"
```

## License

This project is MIT Licensed.
