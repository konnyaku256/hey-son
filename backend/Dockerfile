# build stage
FROM golang:1.13-rc-alpine as build-stage
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o /heyson

# production stage
FROM alpine:3.11 as production-stage
RUN apk add alsa-utils
COPY audio .
COPY --from=build-stage /heyson .
EXPOSE 3000
CMD ["./heyson"]