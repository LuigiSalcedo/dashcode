# dashcode
Backend. Dashcode is a project to manage code and programming tasks.

## Requirements
- MariaDB or MySQL on your device.
- Go (Golang) compiler.

## Configurations
- Compile the file **server.go** using
```
go build server.go
```
This will create an executable file.

- Load database schema

on schema folder you will found a file named "schema.sql". Load it on your database manager
and execute it to create the database schema.

## Executing
To run the server you have to run the next command on the command console.
```
server <user> <password> dashcode
```
user and password are your database user access credentials. "dashcode" is the default database name used by the schema.sql file.

If nothing went wrong the server should start on http://localhost:8080.

## Endpoints
This is the list of endpoints of the server:
| Path | Method | Request format | Request description | Response format | Response description |
--------|-------|----------------|---------------------|-----------------|--------------------------
| /login | GET | JSON | email: string, password: string |JSON | JWT to access to others endpoint or error message |
| /register | POST | JSON | id: int, name: string, email: string, password: string | JSON | error message or null with 201 status code |