mkdir artifacts
go test -coverprofile ./artifacts/coverage.out ./...
go tool cover -html=./artifacts/coverage.out -o ./artifacts/coverage.html
start ./artifacts/coverage.html