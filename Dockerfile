FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
COPY *.go ./
RUN go build /docker-gs-ping
EXPOSE 8089
CMD ["/docker-gs-ping"]

