
# Golang app demo

the app consumes a JSON payload from https://www.data.gov/, populates a database and displays the database contents on a web page.

## run application

1. docker-compose up
1. ssh to 'backend' conainer
1. inside the conainer:

```
cd src/backend/
go run .
```
after the server is running:
- import data to the database `http://localhost:4200/import`
- query data from the database `http://localhost:4200/export`

