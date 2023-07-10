FROM golang:latest
RUN mkdir /appsite
ADD . /appsite/
WORKDIR /appsite
RUN CGO_ENABLED=0 GOOS=linux go build /my_site_app
CMD ["chmod", "+x","/my_site_app"]
