# Todo List API
This is a coding task to implement a web service that allows a user to manage a todo list. The user can:
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
- [GET] /todos/search?q=search_term -> Searches for todo entries with the given search term

## Running locally
To test the program locally, make sure you have Go installed on your machine and PostgreSQL database server.
- clone this repo
- from the terminal, run `go build` in the root directory
- make sure you have postgres running locally and you have set up the database
- execute the built binary and test it using curl or your favorite API testing tool. 

Make sure to set appropriate env variables for the database connection to be made successfully, check the `.env.example` file for the required variables

