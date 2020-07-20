FROM golang:alpine

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .

WORKDIR /dist/server

RUN cp /build/main .

EXPOSE 3000

CMD ["/dist/main"]