# Go with Dabatase

## run MySQL with Docker
```
docker run -p 3306:3306 -e MYSQL_USER=golang -e MYSQL_PASSWORD=golang -e MYSQL_ROOT_PASSWORD=golang -e MYSQL_DATABASE=wilbur mysql:latest
```