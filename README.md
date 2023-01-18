# Echo

A Docker image that serves an HTTP server, responding with these when hit:
- Request headers
- Request path
- URL queries
- User agent

## Usage

How to run:
```
docker run -p 8000:80 dewadg/echo
```

Make a request:
```
curl 'http://localhost:8000/' 
```
