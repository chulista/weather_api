#My hello world
FROM golang:1.14

WORKDIR /go/src/github.com/chulista/weather_api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["cmd"]