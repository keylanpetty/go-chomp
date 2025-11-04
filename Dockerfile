FROM golang:1.22 AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Disable VCS stamping to avoid "error obtaining VCS status" 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -buildvcs=false -o /out/chomp .

# ---------- Runtime stage ----------
FROM gcr.io/distroless/base-debian12:nonroot
WORKDIR /app
COPY --from=builder /out/chomp /app/chomp

# Interactive TTY app; no ports exposed
USER nonroot:nonroot
ENTRYPOINT ["/app/chomp"]
