prepare-env= \
	PORT=2565 \
	DATABASE_URL=postgres://root:root@localhost:5432/assessment-db?sslmode=disable	\
	AUTH_TOKEN="November 10, 2009"	\

start:
	${prepare-env} \
	go run server.go

check:
	staticcheck ./...

test:
	go test -v ./...

cover:
	go test -v ./... -cover -coverprofile=coverage/cover.out && \
	go tool cover -html=coverage/cover.out -o coverage/coverage.html