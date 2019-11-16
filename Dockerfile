FROM golang:1.13-alpine
EXPOSE 8080

ENV GOPATH=/go
RUN mkdir -p $GOPATH/src/github.com/realbucksavage/todos
COPY . $GOPATH/src/github.com/realbucksavage/todos

WORKDIR $GOPATH/src/github.com/realbucksavage/todos
RUN go build -o todos .

CMD ["/go/src/github.com/realbucksavage/todos/todos"]
