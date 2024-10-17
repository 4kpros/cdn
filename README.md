# Simple and well-structured API for multiple projects

This repository contains a simple API written in Go, designed to be reusable and easily integrated into multiple projects. The API is built with a focus on performance, scalability, and maintainability. 

It provides a clean and consistent interface for accessing and managing data, making it an ideal choice for a variety of applications.


# Key features

**- Good performance:** The API is designed to handle a high volume of requests with low latency.

**- Scalable:** The API can be easily scaled to accommodate increasing workloads.

**- Well-structured:** The API is well-documented and easy to use, with a consistent and intuitive interface.

**- Reusable:** The API can be easily integrated into multiple projects, reducing development time and effort.

# To get started with the API, follow these steps:

### 1. Requirements

  - Make installed for shortcuts

  - Docker installed

  - Rename .env.example to ```app.env```

  Others information such configurations are on ```app.env```

### 2. Clone the repository

```go
git clone https://github.com/your-username/go-api.git
```

```go
cd go-api/
```

The entry point of the project is `cmd/` folder. In this folder the is a `main.go` file.

### 3. Install dependencies

```go
make install
```

### 4. Run the API

```go
make build
```

```go
make run
```

API docs with openAPI v3.1(latest) is on 
```go
/api/v1/docs
```

If you want to scan vulnerabilities(security issues)
```go
make scan
```

# Features

- [x] Images


# Contributing

We welcome contributions to this project. If you have any ideas or improvements, please feel free to open an issue or pull request.

We believe that this API can be a valuable tool for developers who need to build high-performance, scalable, and maintainable applications. We encourage you to try it out and let us know what you think.
