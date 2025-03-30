# --- Build Stage ---
    FROM golang:latest AS builder

    WORKDIR /app
    COPY . .
    
    RUN go mod tidy
    RUN go build -o chainnews_api_gateway
    
    # --- Runtime Stage ---
    FROM ubuntu:latest
    
    WORKDIR /root/
    COPY --from=builder /app/chainnews_api_gateway .
    
    EXPOSE 8000
    CMD ["./chainnews_api_gateway"]
    