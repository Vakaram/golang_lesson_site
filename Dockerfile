FROM golang:latest
RUN mkdir /appsite
ADD . /appsite/
WORKDIR /appsite
RUN go build -o /my_site/main .
CMD ["/appsite/my_site/main"]
