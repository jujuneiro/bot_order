WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o app
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/internal/HTTP/http ./internal/HTTP/http
   # Создаём папку template и копируем config.yaml
RUN mkdir -p template
COPY template/config.yaml ./template/config.yaml

EXPOSE 8081
ENTRYPOINT ["./app"]
