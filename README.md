# Web-server

This is a simple Go web server project that uses Gorilla Mux for routing and YAML for configuration. The project demonstrates how to structure a Go web application, load configuration from a file, and apply middleware for logging, authentication, and error handling.

## Project Structure
myproject/
├── cmd/
│ └── main.go
├── internal/
│ ├── config/
│ │ └── config.go
│ └── handler/
│ └── handlers.go
├── configs/
│ └── config.yaml
├── go.mod
└── go.sum



## Configuration

The configuration file `config.yaml` is located in the `configs` directory. It includes settings for the server port and database connection.

```yaml
# configs/config.yaml
server:
  port: ":8080"
database:
  host: "localhost"
  port: 5432
  user: "dbuser"
  password: "dbpass"
  dbname: "mydb"


## Handlers
Handlers for different routes are defined in internal/handler/handlers.go.

HomeHandler: Handles requests to the root path.
AboutHandler: Handles requests to the /about path.
NotFoundHandler: Handles requests to undefined paths and returns a 404 error.

##Middleware
Middleware functions for logging, authentication, and error handling are defined in cmd/myapp/main.go.

loggingMiddleware: Logs the details of each request.
authenticationMiddleware: Checks for an authorization token in the header.
errorHandlingMiddleware: Handles panics and returns a 500 error.


## Usage


Prerequisites: 
    Go (version 1.16 or higher)


##  nstallation
Clone the repository:
    git clone https://github.com/realakinolaisrael/Web-server.git
    cd Web-server


## Install dependencies:
        go mod tidy


##  Running the Server
To run the server, execute the following command:
       go run cmd/main.go
