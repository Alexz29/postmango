FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .
CMD ["/app/main", "-f", "fixtures/fixture.json", "-p", "6868", "-h", "0.0.0.0"]
