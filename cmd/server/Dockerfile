FROM golang:1.19.1
WORKDIR /app
COPY . .
RUN go mod download
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server ./cmd/server
ADD ./cmd/server/config.json /config.json
EXPOSE 8080
CMD ["/server", "-path=/config.json"]