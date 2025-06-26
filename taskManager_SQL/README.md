# 📝 Task Manager API in Go with MySQL (Dockerized)

This project is a simple Task Manager API built using **Go (Golang)** and **MySQL**, with MySQL running in a **Docker container**.

It allows you to:
- Add new tasks
- View all tasks
- Get a task by ID
- Mark a task as complete
- Delete a task



```bash
docker run --name gofr-mysql \
  -e MYSQL_ROOT_PASSWORD=root123 \
  -e MYSQL_DATABASE=test_db \
  -p 3306:3306 \
  -d mysql:8.0.30