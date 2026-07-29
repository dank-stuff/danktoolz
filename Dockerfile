FROM golang:1.26-alpine AS build

WORKDIR /app
COPY . .

RUN apk add --no-cache wget make sed git

RUN go install github.com/a-h/templ/cmd/templ@v0.3.1020
RUN make generate
RUN make build

FROM alpine:latest AS run

WORKDIR /app
COPY --from=build /app/danktoolz .

EXPOSE 8080

CMD ["./danktoolz"]
