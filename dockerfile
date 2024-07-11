# Use the official Golang image as the base image
FROM golang:1.18-alpine


WORKDIR /app


COPY go.mod go.sum ./


RUN go mod download

COPY . .


RUN go build -o hack4bengal

# Expose port 8080 to the outside world
EXPOSE 8080

CMD ["./"]
