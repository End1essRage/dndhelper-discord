run:
	go run cmd/main.go -t=${BOT_TOKEN}
build:
	podman build -t end1essrage/dndhelper-discord .
run_container:
	podman run --env ENVIRONMENT="DEV" --env TOKEN=${BOT_TOKEN} end1essrage/dndhelper-discord