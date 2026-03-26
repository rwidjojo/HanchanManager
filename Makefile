DATABASE_URL="postgres://mahjong:mahjong@host.containers.internal:5432/mahjong?sslmode=disable"

build-dev-image:
	podman build -t golang-dev -f .devcontainer/Dockerfile.dev

migrate-up:
	migrate -path ./migrations -database "$(DATABASE_URL)" up
