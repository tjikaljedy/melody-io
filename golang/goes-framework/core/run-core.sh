export MONGO_URL="mongodb://localhost:27017/eventstore"
export NATS_URL="nats://localhost:4222"

go run ./core/cmd/server/main.go