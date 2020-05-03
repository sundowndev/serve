# http-server

Production ready HTTP server for static file serving

## Installation

```
$ curl ...
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

## License

This project is MIT Licensed.