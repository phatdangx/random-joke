up:
	@ echo "Building latest version"
	@ docker compose up --build -d

down:
	@ docker compose down

test:
	go test random-joke/handler -cover -count=1
	go test random-joke/model -cover -count=1
	go test random-joke/repository/external -cover -count=1
	go test random-joke/usecase -cover -count=1

clean:
	@ docker image prune -f
