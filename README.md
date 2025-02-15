# GoBuildSomething

GoBuildSomething is a minimal CLI tool built with Go that helps you to quickly scaffold a new project.

Simply provide a project name and optional path (defaults to Desktop) to get started with an already initialized Git repository

**Why?**

I just wanted to automate the process of creating new projects, I got tired of doing mkdir, cd and git init multiple times everyday 🥲

## **Getting Started**

### **Prerequisites**

- [Go](https://go.dev/dl/) installed on your system.
- A project idea 🧠

### **Installation**

1. Clone the repository

2. Run the program:

  ```bash
  cd GoBuildSomething && go run main.go
  ```

  or

  ```bash
  go build && go install
  GoBuildSomething
  ```

### **Help**

To get help on how to use the program, run:

```bash
go run main.go --help
```

or

```bash
GoBuildSomething -h
```

> [!Note]
> You can edit the `config.json` file to change the default path. It will be located at `~/go/bin`.
