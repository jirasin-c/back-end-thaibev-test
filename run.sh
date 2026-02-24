export $(cat .env.dev | grep -v "#" | awk "/=/ {print $1}")

go run ./cmd/api/main.go