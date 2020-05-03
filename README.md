# http-server

Production ready HTTP server for static file serving

## Installation

```
# Download the archive
$ curl -L "https://github.com/sundowndev/http-server/releases/download/v2.0.6/http-server_$(uname -s)_$(uname -m).tar.gz" -o http-server.tar.gz

# Extract the binary
tar xfv http-server.tar.gz

# Run the software
./serve -h

# You can install it globally
mv ./serve /usr/bin/serve
```

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

```
$ docker run -it -v $PWD/public:/app/public:ro -p 80:80 sundowndev/http-server
```

#### Using compose

```
version: '3.7'

services:
    serve:
      container_name: http-server
      restart: on-failure
      build:
        context: .
        dockerfile: Dockerfile
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