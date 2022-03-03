FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go env -w GOPRIVATE=github.com/agustinrodriguez7
RUN go build

EXPOSE 8080

CMD [ "/orders-challenge-microservice" ]
