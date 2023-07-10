FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
EXPOSE 8089
CMD ["/docker-gs-ping"]

