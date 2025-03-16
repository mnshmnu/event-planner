FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum  ./

RUN go mod download

COPY . .

RUN go build -o build/server cmd/*

FROM alpine

COPY --from=build /app/build/server /server

EXPOSE 8081

ENTRYPOINT [ "/server" ]
