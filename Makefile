build-dev-image:
	podman build -t golang-dev -f .devcontainer/Dockerfile.dev