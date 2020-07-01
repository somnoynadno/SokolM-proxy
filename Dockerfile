FROM golang:alpine
LABEL maintainer="Alexander Zorkin"

ADD . /app/
WORKDIR /app

RUN go build -o main .
CMD ["/app/main"]