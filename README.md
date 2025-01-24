# CRUD-GO

This project is a sample Golang application implementing CRUD operations (Create, Read, Update, Delete) using a PostgreSQL database. It demonstrates how to build a RESTful API to manage user entities.

## Project Structure

The project is organized as follows:

- **main.go**: Entry point of the application, which sets up the HTTP server and routes.
- **go.mod** and **go.sum**: Dependency management files for Go.
- **postgres_connection.go**: Handles the connection to the PostgreSQL database.
- **docker-compose.yaml**: Configuration file for Docker Compose to orchestrate services.
- **scripts/**: Contains Bash scripts for managing Docker containers.
- **src/**: Directory containing the application's main source code.

## Dependencies

The main dependencies of the project are:

- [Gin Gonic](https://github.com/gin-gonic/gin): A high-performance HTTP web framework for Go.
- [pq](https://github.com/lib/pq): PostgreSQL driver for Go.
- [godotenv](https://github.com/joho/godotenv): Loads environment variables from a `.env` file.

## Prerequisites

- [Go](https://golang.org/dl/): Ensure that Go is installed on your machine.
- [Docker](https://www.docker.com/get-started): Required to run the development environment.
- [Docker Compose](https://docs.docker.com/compose/install/): For orchestrating Docker services.

## Setup and Execution

1. **Clone the repository:**

   ```bash
   git clone https://github.com/jphalexandrino/CRUD-GO.git
   cd CRUD-GO
   ```

2. **Start Docker containers:**

   Use the `up.sh` script to start the containers in detached mode:

   ```bash
   ./scripts/up.sh
   ```

   This will start the services defined in the `docker-compose.yaml` file, including the PostgreSQL database.

3. **Run the application:**

   With the services running, start the Go application:

   ```bash
   go run main.go
   ```

   The application will be available at `http://localhost:8080`.

## API Endpoints

The API exposes the following endpoints to manage users:

- **POST /users**: Creates a new user.
- **GET /users**: Retrieves all users.
- **GET /users/{id}**: Retrieves a specific user by ID.
- **PUT /users/{id}**: Updates user information by ID.
- **DELETE /users/{id}**: Deletes a user by ID.

## Docker Scripts

In the `scripts/` directory, there are Bash scripts to simplify Docker container management:

- **up.sh**: Starts containers in detached mode.
- **stop.sh**: Stops the containers without removing them.
- **down.sh**: Stops and removes the containers.

## Contribution

Contributions are welcome! Feel free to open issues or pull requests for improvements or bug fixes.
