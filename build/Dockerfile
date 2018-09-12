# BUILD: docker build -t people-api -f build/Dockerfile .
# RUN: docker run -d -p 8001:8001 tasks-api
FROM golang:1.10 AS builder
LABEL maintainer="Fabian Rivera <fabian.riveram@gmail.com>"
COPY ./app ./

# Installing required dependencies
RUN go get -d -v github.com/apsdehal/go-logger
RUN go get -d -v github.com/gorilla/mux

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM scratch
COPY --from=builder /app ./
EXPOSE 8001
ENTRYPOINT ["./app"]