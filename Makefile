test/internal:
	go test `go list ./internal/... | grep -v mocks | grep -v migrations` -cover
test/pkg:
	go test `go list ./pkg/... | grep -v mocks` -cover
mocks:
	cd ./internal/user
	mockery --dir ./internal/user --output ./internal/user/mocks --all --case snake
	cd ..