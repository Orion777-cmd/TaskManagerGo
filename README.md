# TaskManagerGo

TaskManager API is a simple RESTful API for managing tasks built with Golang, Gin, and the Bun ORM with a PostgreSQL database.

## Features

- Create, Read, Update, and Delete tasks.
- Retrieve a list of tasks.
- Task details include title, description, due date, and status.

## Prerequisites

Make sure you have the following installed on your machine:

- [Golang](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)

## Getting Started

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/yourusername/TaskManager.git
   cd TaskManager

2 **Database Setup:

Create a PostgreSQL database for TaskManager.
Update the database connection details in database.go with your PostgreSQL credentials.

3. **Install Dependencies:

    go mod tidy

4.. **run the application

    go run main.go

The API will be available at http://localhost:8080.

**API Endpoints
GET /tasks: Get a list of all tasks.
GET /tasks/:id: Get details of a specific task.
POST /tasks: Create a new task.
PUT /tasks/:id: Update an existing task.
DELETE /tasks/:id: Delete a task.