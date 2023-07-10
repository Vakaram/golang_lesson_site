FROM golang:latest
RUN mkdir /appsite
ADD . /appsite/
WORKDIR /appsite
RUN go build main.go .
CMD ["chmod", "+x","/main.go"]
