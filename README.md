# Todo App API

![Go](https://img.shields.io/badge/Go-1.22-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

This is a well-structured and documented GoLand application that implements a Todo App API. It provides features for managing tasks, including adding, retrieving, updating, and deleting them.

## Installation

Follow these steps to install and run the application:

1. **Clone the repository**:
    ```sh
    git clone https://github.com/cherrycutter/todo-app-api.git
    ```
2. **Install Dependencies**:
    ```sh
    go mod download
    ```
3. **Run the application**:
    ```sh
    make build && make run
    ```
4. **Make Migrations** (if the application is launched for the first time):
    ```sh
    make migrate
    ```

## Usage

The API endpoints for managing tasks are designed to follow RESTFUL principles:

1. **Retrieve all todos**:
    ```http
    GET /todos
    ```

2. **Create a new todo**:
    ```http
    POST /todos
    ```

3. **Retrieve a specific todo by its ID**:
    ```http
    GET /todos/:id
    ```

4. **Update todo fields with the provided ID**:
    ```http
    PATCH /todos/:id
    ```

5. **Delete a todo with the provided ID**:
    ```http
    DELETE /todos/:id
    ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.