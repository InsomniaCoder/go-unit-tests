# Test Command

go test .
go test -cover ./...

## Generate report 
go test -cover -coverprofile=c.out ./...
go tool cover -html=c.out -o coverage.html 