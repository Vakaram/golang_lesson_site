FROM golang:latest
RUN mkdir /appsite
ADD . /appsite/
WORKDIR /appsite
CMD ["/appsite/main"]
