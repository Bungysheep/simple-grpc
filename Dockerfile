FROM golang

WORKDIR /app/simple-grpc

COPY ./go.mod ./

RUN go mod download

COPY ./src ./src

RUN go build -o ./bin/server ./src/server/server.go

RUN go build -o ./bin/client ./src/client/client.go