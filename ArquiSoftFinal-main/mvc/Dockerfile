FROM golang:1.17

ADD . /backend

WORKDIR /backend
RUN go mod init mvc
RUN go mod tidy
RUN go build -o mvc .
RUN chmod +x /backend

ENTRYPOINT ["/backend/mvc"]