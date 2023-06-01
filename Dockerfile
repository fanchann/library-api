FROM golang:1.20 AS builder

WORKDIR /library_api
COPY . /library_api/
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /library_api/library_app /library_api/main.go


FROM alpine:latest

WORKDIR /library_api
COPY --from=builder /library_api .
ENTRYPOINT [ "/library_api/library_app" ]