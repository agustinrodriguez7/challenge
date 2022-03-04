FROM golang:1.17
COPY . /go/src/github.com/user/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice
WORKDIR /go/src/github.com/user/agustinrodriguez7/vidflex-challenge/src/orders-challenge-microservice
RUN go get ./
RUN go build
CMD ["./orders-challenge-microservice"]
EXPOSE 8080