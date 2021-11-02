test/services:
	go test `go list ./services/... | grep -v mocks | grep -v migrations` -cover
test/pkg:
	go test `go list ./pkg/... | grep -v mocks` -cover
mocks:
	cd ./services/user
	mockery --dir ./services/user --output ./services/user/mocks --all --case snake
	cd ..