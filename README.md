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

The server will start on the port specified in config.yaml (default is :8080).

## Testing Endpoints
Access http://localhost:8080 to check the home page.
Access http://localhost:8080/about to check the about page.
Access an undefined route like http://localhost:8080/undefined to test the custom 404 handler.


## Dependencies

Gorilla Mux: A powerful URL router and dispatcher for Golang.
YAML.v2: A YAML support for the Go language.


## License
This project is licensed under the MIT License - see the LICENSE file for details.