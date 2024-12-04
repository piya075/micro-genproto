
### Build Docker

docker build -t auto-proto:latest . -f DockerFile

docker run -d --name auto-proto -v ${PWD}/proto:/proto auto-proto

### Run Docker & Cli in Docker run cmd

protoc --proto_path=. --micro_out=. --go_out=. proto/user/user.proto
