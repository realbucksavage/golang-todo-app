# Build stage
FROM golang:1.13-alpine AS build

ENV GOPATH=/go
ENV APPNAME=github.com/realbucksavage/golang-todo-app
RUN mkdir -p $GOPATH/src/$APPNAME
COPY . $GOPATH/src/$APPNAME

WORKDIR $GOPATH/src/$APPNAME
RUN go mod vendor && go build -o $GOPATH/todos .

# Deployment Stage
FROM alpine:3.7
EXPOSE 8080

WORKDIR /app
COPY --from=build /go/todos /app/
ENTRYPOINT ./todos

