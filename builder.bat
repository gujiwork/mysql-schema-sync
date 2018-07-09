go build -o ./mysql-schema-sync.exe
set GOOS=linux
set GOARCH=amd64
go build -o ./mysql-schema-sync