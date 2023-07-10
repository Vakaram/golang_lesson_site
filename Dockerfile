FROM golang:latest
RUN mkdir /appsite
ADD . /appsite/
WORKDIR /appsite
RUN go build /my_site_app
CMD ["chmod", "+x","/my_site_app"]
