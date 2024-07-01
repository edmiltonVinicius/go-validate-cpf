FROM golang:1.22 AS build

RUN mkdir /app
WORKDIR /app
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY go.mod go.sum ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/api ./
RUN addgroup -S nonroot && adduser -S nonroot -G nonroot
USER nonroot

EXPOSE 8080
CMD [ "./api" ]