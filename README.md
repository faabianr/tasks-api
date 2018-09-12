# Simple Go Rest API + Dockerfile

## Overview
This is a simple Rest API built with Go. It shows how to create a minimal API and how to deploy it within a docker container.

## Docker
### Build the image
```
$ docker build -t people-api -f build/Dockerfile .
```

### Run the container
```
$ docker run -d -p 8001:8001 tasks-api
```
