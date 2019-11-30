[![Codacy Badge](https://api.codacy.com/project/badge/Grade/a7c64154be87401a9b35a1a51bf48a9d)](https://www.codacy.com/manual/jgodara/golang-todo-app?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=realbucksavage/golang-todo-app&amp;utm_campaign=Badge_Grade)

# Simple todo API

The GitHub repo for my [blog series](https://www.sarvika.com/2019/11/21/build-scalable-web-application-in-go/) "Building a scalable web application in Go"

## Downloading and running

Assuming that you have `$GOPATH` set-up

```shell
$ go get -u github.com/realbucksavage/golang-todo-app
$ cd $GOPATH/src/github.com/realbucksavage/golang-todo-app
$ dep ensure
$ docker-compose up
```

## APIs

### `GET /api/todos/` : List all todo items

```shell
$ curl --silent localhost:8080/api/todos/ | jq .
```

```json
[
  {
    "completed": true,
    "title": "Todo #1",
    "id": 1
  }
]
```

### `GET /api/todos/:id`: List single todo item

```shell
$ curl --silent localhost:8080/api/todos/1 | jq .
```

```json
{
  "completed": false,
  "title": "Todo #1",
  "id": 1
}
```

### `POST /api/todos/` : Create todo item

```shell
$ curl --silent -X POST localhost:8080/api/todos/ -d '{"title": "Todo #1"}'  | jq .
```

```json
{
  "completed": false,
  "title": "Todo #1",
  "id": 1
}
```

### `PATCH /api/todos/:id` : Update a single todo item

```shell
$ curl --silent -X PATCH localhost:8080/api/todos/1 -d '{"completed": true}'  | jq .
```

```json
{
  "completed": true,
  "title": "Todo #1",
  "id": 1
}
```

```shell
$ curl --silent -X PATCH localhost:8080/api/todos/2 -d '{"title": "Changed title #2"}'  | jq .
```

```json

{
  "completed": false,
  "title": "Changed title #2",
  "id": 2
}
```

### `DELETE /api/todos/:id` : Delete single todo item

```shell
$ curl --silent -X DELETE localhost:8080/api/todos/1  | jq .
```
```shell
$ curl --silent localhost:8080/api/todos/ | jq .
```

```json
[
  {
    "completed": false,
    "title": "Changed title #2",
    "id": 2
  }
]
```