docker build -f deployments/Dockerfile -t ghcr.io/kirimi/rb-bot:latest --platform=linux/amd64 .
docker push ghcr.io/kirimi/rb-bot --all-tags