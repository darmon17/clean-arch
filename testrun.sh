# go test ./features/... -coverprofil=cover.out && go tool cover -html=cover.out


# go test -v -coverprofile cover.out ./features/...
# go tool cover -html=cover.out -o cover.html
# open cover.html


go test -coverprofile=coverage.out ./features/... ;    go tool cover -html=coverage.out