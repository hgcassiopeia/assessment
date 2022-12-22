prepare-env= \
	PORT=2565 \
	DATABASE_URL=postgres://root:root@localhost:5432/assessment-db?sslmode=disable	\
	AUTH_TOKEN="November 10, 2009"	\

start:
	${prepare-env} \
	go run server.go

check:
	staticcheck ./...