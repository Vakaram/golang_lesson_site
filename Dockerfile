FROM golang:latest
RUN mkdir /appsite
ADD . /appsite/
WORKDIR /appsite
RUN go build -o main .
CMD ["/appsite/my_site/main"]
