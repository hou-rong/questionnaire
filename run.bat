set GOARCH=amd64
set GOOS=linux
go build -v -o questionnaire -ldflags="-extld=$CC"