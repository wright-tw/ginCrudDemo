# run all unit test
go test ./... -v -count=1 -failfast 

# binding env 
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

# build binary file
go build -o bin/gintest cmd/gintest/main.go 
go build -o bin/migrate cmd/migrate/main.go

# build docker image
docker build --build-arg ENV=example -t=wright1992/gintest:latest .