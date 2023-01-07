FROM golang:1.19-alpine as build
WORKDIR /app
COPY . .
RUN go build  ./cmd/twitter

FROM alpine:3.17
COPY --from=build /app/twitter .
COPY --from=build /app/migrations ./migrations
CMD ["./twitter"]
