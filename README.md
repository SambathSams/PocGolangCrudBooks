# Project Name

A Go-based application.

---

## 🚀 Getting Started

### Prerequisites

Make sure you have the following installed:

- [Go](https://golang.org/dl/) (latest stable version recommended)
- `swag` CLI (for Swagger documentation generation)

Install `swag` if you don't have it:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

---

## 📦 Installation

Clone the repository and install dependencies:

```bash
go get ./...
```

---

## ▶️ Running the Project (Development)

To start the project in the development environment:

```bash
go run main.go
```

---

## 📄 Swagger Documentation

This project uses Swagger for API documentation.

After making changes to Swagger annotations in the code, regenerate the Swagger documentation by running:

```bash
swag init
```

This will update the Swagger UI files.

---

## 📌 Notes for Developers

- Always run `swag init` after modifying Swagger comments.
- Ensure dependencies are up to date with `go get ./...`.
- Follow Go best practices and formatting:

```bash
go fmt ./...
```