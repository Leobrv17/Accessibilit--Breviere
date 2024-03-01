FROM golang:latest
WORKDIR /app
COPY . .
RUN go get github.com/lib/pq
EXPOSE 8080
CMD ["go", "run", "cmd/main.go"]