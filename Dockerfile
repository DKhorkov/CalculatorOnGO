FROM golang:alpine
LABEL authors="dkhorkov"

WORKDIR /app
COPY . .
RUN go build -o calculator ./src/main/main.go
CMD ["./calculator"]