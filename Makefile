up:
	@ echo "Building latest version"
	@ docker build -t random-joke:latest .
	@ docker compose up -d

down:
	@ docker compose down

test:
	go test random-joke/handler -cover -count=1
	go test random-joke/model -cover -count=1
	go test random-joke/repository/external -cover -count=1
	go test random-joke/usecase -cover -count=1