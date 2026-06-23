FROM golang:1.25-alpine AS build

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o portfolio .


FROM alpine:latest

WORKDIR /app

COPY --from=build /app/portfolio .

COPY templates ./templates
COPY components ./components
COPY static ./static

EXPOSE 8090

CMD ["./portfolio"]