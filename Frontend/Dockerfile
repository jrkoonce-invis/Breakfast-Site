FROM golang:alpine

WORKDIR /build

COPY . .

RUN go build -o main .

WORKDIR /dist/client

RUN cp /build/main .

EXPOSE 9090

CMD ["/dist/main"]