FROM golang:1.17-alpine3.16 AS build
WORKDIR /app
COPY go.mod go.sum /app/
RUN  go mod download
COPY . /app
RUN go build -o /app /app/...

## Deploy
FROM alpine:3.16
WORKDIR /app
COPY --from=build /app/api /app/api
ENV BACKEND_PORT=8081
EXPOSE ${BACKEND_PORT}
ENTRYPOINT ["/app/api"]