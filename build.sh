export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build -o bin/gintest cmd/gintest/main.go 
go build -o bin/migrate cmd/gintest/migrate.go
docker build --build-arg ENV=example -t=wright1992/gintest:latest .