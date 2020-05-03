FROM golang:1.14-alpine as go_builder

LABEL maintainer="Sundowndev" \
  org.label-schema.name="http-server" \
  org.label-schema.description="Production ready HTTP server for static file serving." \
  org.label-schema.url="https://github.com/sundowndev/http-server" \
  org.label-schema.vendor="Sundowndev" \
  org.label-schema.schema-version="1.0"

WORKDIR /app

COPY . .

RUN go get -v -t -d ./...

RUN go build -v -o serve .

FROM golang:1.14-alpine

WORKDIR /app

COPY --from=go_builder /app/serve .

EXPOSE 80

ENTRYPOINT ["/app/serve"]