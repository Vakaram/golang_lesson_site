FROM golang:latest
RUN mkdir /appsite
ADD . /appsite/
WORKDIR /appsite
RUN go build main.go
CMD ["/appsite/my_site/main.go"]
