# Go - Swagger

## Install

### swagger-editor

docker pull swaggerapi/swagger-editor
docker run --rm -p 80:8080 swaggerapi/swagger-editor

### go-swagger

brew tap go-swagger/go-swagger
brew install go-swagge

---

## Generate

swagger generate spec -o ./swagger.json