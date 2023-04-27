GO_ENV=test go test ./...  -cover -covermode=count -coverpkg=./... -coverprofile=cover.out -v
go tool cover -func cover.out
go tool cover -html=cover.out -o cover.html