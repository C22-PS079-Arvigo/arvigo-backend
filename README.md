### Arvigo Backend

This repository contains the backend implementation for the Arvigo app, which is an application for trying out products before making a purchase.

## Project Description

The Arvigo backend is built using the Go programming language and the Echo framework. It provides API endpoints for user authentication and user-related operations. The backend utilizes JWT for authentication, MySQL for data storage, and includes database migration capabilities.

## Getting Started

To get started with the Arvigo backend, follow these steps:

1. Clone the repository: `git clone <repository-url>`
2. Install the dependencies: `go mod download`
3. Set up the environment variables by creating a `.env` file (refer to .env section below).
4. Run the application: `make run-local`

Make sure you have a MySQL database available and update the database connection details in the `.env` file.

## Environment Variables

The following environment variables are required to run the Arvigo backend:

- `PORT`: The port on which the server will listen.
- `DB_USER`: The username for the MySQL database.
- `DB_PASSWORD`: The password for the MySQL database.
- `DB_HOST`: The host address of the MySQL database.
- `DB_PORT`: The port number of the MySQL database.
- `DB_NAME`: The name of the MySQL database.
- `JWT_SECRET`: The secret key for JWT token generation and validation.

Make sure to set these variables in the `.env` file before running the application.

## License

This project is licensed under the MIT License.