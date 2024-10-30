![Linting](https://github.com/UpsDev42069/BM_Search_Engine/actions/workflows/lint.yml/badge.svg)

# BM_Search_Engine

This is a search engine that searches for code language keywords. The project is written in Go for the backend and in Svelte for the frontend.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them:

- **Go**: You need to have Go installed on your machine. You can download it from [here](https://golang.org/dl/).
- **Node.js and npm**: You need to have Node.js and npm installed. You can download them from [here](https://nodejs.org/).

### Installing

A step-by-step series of examples that tell you how to get a development environment running.

1. **Clone the repository**:
   ```sh
   git clone https://github.com/UpsDev42069/BM_Search_Engine.git
   cd BM_Search_Engine/backend
   ```
2. **Install dependencies**:

   ```sh
   install docker
   install docker-compose
   ```

3. **Run docker-compose**:
   ```sh
   docker-compose up --build
   ```

### Running the Tests

Explain how to run the automated tests for this system.

#### Backend Tests

1. **Navigate to the backend directory**:

   ```sh
   cd backend
   ```

2. **Run the tests**:

   ```sh
   go test ./...
   ```

#### Frontend Tests

WORK IN PROGRESS

### Deployment

[The frontend is deployed here:](13.79.97.206:8069)

[The backend is deployed here:](13.79.97.206:8080)

### Built With

- [Go](https://golang.org/) - The backend programming language
- [Svelte](https://svelte.dev/) - The frontend framework

### API Documentation

To see the swagger api ui
run the project

```sh
 cd backend
 go run main.go
```

and go to http://localhost:8080/swagger/index.html
