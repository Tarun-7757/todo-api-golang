# 📝 Todo API in Golang

A simple RESTful API for managing todo tasks, built using Go.

## 🚀 Features

- Create, Read, Update, and Delete (CRUD) todos
- Clean and simple folder structure
- Written in idiomatic Go

## 📁 Project Structure

todo-api-golang/
├── main.go          # Entry point
├── models/          # Data models
├── handlers/        # HTTP request handlers
├── routes/          # Route initialization
├── go.mod           # Module file
├── go.sum           # Dependencies
└── README.md        # Project documentation

## 🛠️ Requirements

- Go 1.18+
- Git

## 🧪 Running Locally

# Clone the repo
git clone https://github.com/Tarun-7757/todo-api-golang.git

# Navigate into the project
cd todo-api-golang

# Run the server
go run main.go

## 🧵 Sample Endpoints (Assumed)

GET    /todos         → Get all todos  
POST   /todos         → Create new todo  
PUT    /todos/{id}    → Update a todo  
DELETE /todos/{id}    → Delete a todo  

## 📬 Example Request (POST /todos)

{
  "title": "Learn Go",
  "description": "Practice building REST APIs",
  "completed": false
}

## ✅ Future Improvements

- Add persistent storage (PostgreSQL, MongoDB, etc.)
- Authentication & Authorization
- Dockerize the app
- Add unit tests

## 🧑‍💻 Author

Tarun Sharma  
GitHub: https://github.com/Tarun-7757

## 📄 License

This project is licensed under the MIT License. See LICENSE file for details.
