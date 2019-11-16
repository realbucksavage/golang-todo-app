# Todos App

## APIs

### `GET` [/api/todos](/api/todos) : List all todo items

### `GET` [/api/todos/:id](/api/todos/5) : List single todo item

### `PATCH` [/api/todos/:id](/api/todos/5) : Update a single todo item

```json
{
  "title": "Changed Title",
  "completed": true
}
```

### `POST` [/api/todos](/api/todos) : Create todo item

```json
{
  "title": "Some title",
  "completed": false
}
```

### `DELETE` [/api/todos/:id](/api/todos/5) : Delete single todo item
