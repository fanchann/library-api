
## Library Book API

### Framework 
- GIN
### Database
- Mysql
----

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