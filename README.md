# Arvigo Backend

This repository contains the backend implementation for the Arvigo app, which is an application for trying out products before making a purchase.

## Project Description

The Arvigo backend is built using the Go programming language and the Echo framework. It provides API endpoints for user authentication and user-related operations. The backend utilizes JWT for authentication, MySQL for data storage, and includes database migration capabilities.

## Getting Started

To get started with the Arvigo backend, follow these steps:

1. Clone the repository: `git clone https://github.com/Arvigo-Capstone/arvigo-backend.git`
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

## Project Structure
```bash
├── CHANGELOG.md
├── Dockerfile
├── Makefile
├── README.md
├── constant
│   ├── date.go
│   └── role.go
├── datastruct
│   ├── address.go
│   ├── brand.go
│   ├── category.go
│   ├── location.go
│   ├── marketplace.go
│   ├── others file ...
├── go.mod
├── go.sum
├── main.go
├── middleware
│   └── auth.go
├── pkg
│   └── database
│       ├── conn.go
│       └── migration
│           ├── 001_indonesian.sql
│           ├── 002_datas.sql
│           └── 003_seeders.sql
├── repository
│   ├── auth_repository.go
│   ├── controller.go
│   ├── location_repository.go
│   └── others file ...
├── route
│   ├── auth.go
│   ├── location.go
│   └── user.go
└── utils
    ├── array.go
    ├── converter.go
    ├── response.go
    ├── string.go
    ├── response.go
    ├── others file ...
```

## Branch Naming Convention
This project aims to demonstrate the application of Trunk-Based Development (TBD) methodology in software development. TBD promotes a streamlined and collaborative approach to version control and continuous integration, enabling teams to deliver high-quality software with shorter lead times.

For more information about Trunk-Based Development, please visit [https://trunkbaseddevelopment.com/](https://trunkbaseddevelopment.com/).

When naming branches, it is recommended to use the following prefixes:

1. `feature/`: Use this prefix for adding new features.
2. `hotfix/`: Use this prefix for fixing critical issues or bugs in production services.
3. `bugfix/`: Use this prefix for resolving non-critical issues or bugs.
4. `refactor/`: Use this prefix when making changes to the folder structure or code without altering functionality.

Additionally, when preparing a branch for release, please include the following release tag:

- `release-arvigo-backend-*`: Use this tag for releases related to the API version specified in the main.go file.
Ensure that the version in the release tag corresponds to the API version mentioned in the main.go file.

## License

This project is licensed under the MIT License.


