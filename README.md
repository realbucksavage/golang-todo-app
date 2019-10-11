# Todos App

A simple todos API implemented in [Revel](https://github.com/revel/revel).

## Running

```shell
$ go get github.com/revel/cmd/revel
$ revel run github.com/realbucksavage/golang-todos-app
```

## APIs

### `GET` [/api/todos](/api/todos) : List all todo items

### `GET` [/api/todos/:id](/api/todos/5) : List single todo item

### `PUT` [/api/todos/:id](/api/todos/5) : Update a single todo item

```json
{
  "Title": "Changed Title",
  "Completed": true
}
```

### `POST` [/api/todos](/api/todos) : Create todo item

```json
{
  "Title": "Some title",
  "Completed": false
}
```

### `DELETE` [/api/todos/:id](/api/todos/5) : List single todo item
