test:
	go test -cover ./... -coverprofile=.output/cover.out && \
	go tool cover -html=.output/cover.out -o .output/cover.html

run:
	go run ./...
