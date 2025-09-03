# My First Go Web Server

A simple HTTP server built with Go's standard library to learn the basics of the language.

## Features

- `GET /`: Returns a "Hello World" JSON message.
- `GET /greet/{name}`: Returns a personalized greeting in JSON.

## Prerequisites

- Go 1.16 or higher installed.

## Getting Started

1.  **Clone this repository:**
    ```bash
    git clone <your-repo-url>
    cd my-first-go-server
    ```

2.  **Run the application:**
    ```bash
    go run main.go
    ```
    The server will start on `http://localhost:8080`.

## Testing the API

- Visit `http://localhost:8080` in your browser.
- Visit `http://localhost:8080/greet/YourName` to get a personalized message.
- Or use `curl` from the command line:
  ```bash
  curl http://localhost:8080
  curl http://localhost:8080/greet/World
