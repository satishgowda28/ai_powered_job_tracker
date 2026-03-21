# Project: AI-Powered Job Tracker

## Overview

This project is an AI-powered job application tracker. It consists of a Go backend and is intended to have a React/TypeScript frontend in the future.

The backend is a web service built with the Fiber framework. It provides a RESTful API for user authentication and other job-tracking-related functionalities. The backend is designed with a clean architecture, separating concerns into different layers like handlers, services, repositories, and a database layer.

## Backend Details

*   **Language:** Go
*   **Framework:** Fiber
*   **Database:** PostgreSQL (using `pgx` driver)
*   **Authentication:** JWT-based authentication with password hashing using `argon2id`.
*   **API Endpoints:**
    *   `GET /health`: Health check
    *   `POST /auth/register`: User registration
    *   `POST /auth/login`: User login
*   **Live Reloading:** The project uses `air` for live reloading during development.

## Frontend Details

The `FrontEnd` directory is currently empty, but the `README.md` suggests that a React + TypeScript frontend is planned.

## Building and Running

To run the backend with live reloading, you can use the `air` tool.

1.  **Install air:**
    ```bash
    go install github.com/cosmtrek/air@latest
    ```
2.  **Run the backend:**
    From the `Backend` directory, run:
    ```bash
    air
    ```
    This will build the application and run it. The application will automatically restart when you make changes to the Go source files.

Alternatively, you can build and run the application manually:

1.  **Build the application:**
    From the `Backend` directory, run:
    ```bash
    go build -o ./tmp/main .
    ```
2.  **Run the application:**
    ```bash
    ./tmp/main
    ```

## Development Conventions

*   The backend follows a structured project layout, separating different concerns into packages like `handlers`, `services`, `repositories`, etc.
*   Database migrations are managed using `goose`. The migration files are located in `Backend/db/migration`.
*   SQL queries are defined in `.sql` files in `Backend/db/queries` and are used by `sqlc` to generate type-safe Go code.
