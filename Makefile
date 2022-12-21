prepare-env= \
	PORT=2565 \

start:
	${prepare-env} \
	go run server.go

check:
	staticcheck ./...