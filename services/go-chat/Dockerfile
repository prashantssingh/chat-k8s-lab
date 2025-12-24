# services/go-chat/Dockerfile
FROM golang:1.25-alpine AS build
WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/go-chat ./cmd/server

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=build /out/go-chat /go-chat
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/go-chat"]
