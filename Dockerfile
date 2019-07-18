FROM golang

WORKDIR /app/simple-grpc

COPY ./ ./

RUN go build -o ./bin/server ./src/server/server.go

CMD ["./bin/server"]