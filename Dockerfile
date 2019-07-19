FROM golang

WORKDIR /app/simple-grpc

COPY ./go.mod ./

RUN go mod download

COPY ./src ./src

RUN go build -o ./bin/server -tags dev ./src/server/server.go

RUN go build -o ./bin/client -tags dev ./src/client/client.go