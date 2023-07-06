FROM golang:latest
RUN mkdir /appsite
ADD . /appsite/
WORKDIR /appsite
RUN go build my_site/main.go
CMD ["/appsite/my_site/main.go"]
