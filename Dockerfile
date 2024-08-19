FROM golang:latest

# Install PostgreSQL client to have access to psql
RUN apt-get update && apt-get install -y postgresql-client

# Verify installation
RUN psql --version

ENV GOPATH=/

COPY ./ ./

# Make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# Build Go app
RUN go mod download
RUN go build -o todo-app ./cmd/main.go

CMD ["./todo-app"]
