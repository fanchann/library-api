
## Library Book API

### Framework 
- GIN
### Database
- Mysql
----

## Installation
run the command `go mod download` to download all the required dependencies.
<br>

## Configuration
- Before running the API, you need to configure the connection to the MySQL database. Here are the configuration steps:

- Open the `/pkg/environments/.env` file in the pkg directory.

- Modify the following configuration values according to your MySQL database settings:
```
export DB_DRIVER=mysql
export DB_AUTH_USERNAME=<your username here>
export DB_AUTH_PASSWORD=<your password here>
export DB_NAME=<database target>
export DB_URL=<database url>
export DB_PORT=<database port>
```
- Create table with migrate
```
migrate -database 'mysql://<DB_AUTH_USERNAME>:<DB_AUTH_PASSWORD>@tcp(<DB_URL>:<DB_PORT>)/<DB_NAME>' -path /pkg/database/migrations up
```

- Run API with 
```
go run main.go
```
<br>

## Test With POSTMAN

Open Postman and create a new collection or open an existing collection where you want to import the API specification.

- Go to the "Collections" tab on the left sidebar and select the desired collection.

- Click on the three-dot menu icon (⋮) next to the collection's name and choose "Edit."

- In the collection editor, click on the "API" tab at the top.

- Within the "API" tab, click on the "Import" button.

- In the import options, select the "Import From File" option.

- Browse and select the `/models/library.yaml`

- Postman will analyze the file and attempt to import the specifications. If any issues are found, it will display a summary of the problems encountered during the import process.

- Review the import summary and make any necessary adjustments or fixes as suggested by Postman.

- Click on the "Import" button to complete the import process.