mkdir -p output
gofmt -w .
go mod tidy
go test -gcflags all=-l -cover -race -coverprofile=output/cover.txt ./... && go tool cover -func=output/cover.txt && go tool cover -html=output/cover.txt
