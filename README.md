# Todo API
This is a web service that allows a user to:
- Create a todo list
- Add items to the todo list
- Delete an item from the todo list
- Fetch all items in the todo list
- Search the list of items in the todo list for a given search text.

## Requirements
- Go >= 1.16
- PostgreSQL server 13

## Endpoints
- [GET] /todos -> Fetches all todo list items
- [GET] /todos/:id -> Fetches a particular todo item by id
- [POST] /todos -> Create todo
- [GET] /todos/search?q=search_term -. Searches for todo entries with the given search term
