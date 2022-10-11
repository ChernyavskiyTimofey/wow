FROM golang:1.19.1
WORKDIR /app
COPY . .
RUN go mod download
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /client ./cmd/client
ADD ./cmd/client/config.json /config.json
CMD ["/client", "-path=/config.json"]