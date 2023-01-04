prepare-env= \
	PORT=2565 \
	DATABASE_URL=postgres://root:root@localhost:5432/assessment-db?sslmode=disable

start:
	docker-compose -f docker-compose.dev.yml up -d && \
	${prepare-env} \
	go run server.go

stop:
	docker-compose -f docker-compose.dev.yml down

check:
	staticcheck ./...

test:
	go test -v ./...

cover:
	go test -v ./... -cover -coverprofile=coverage/cover.out && \
	go tool cover -html=coverage/cover.out -o coverage/coverage.html

deploy:
	docker-compose up --build

deploy-down:
	docker-compose down

integration:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit --exit-code-from test_app
