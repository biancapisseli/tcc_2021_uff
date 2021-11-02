test/services:
	go test `go list ./services/... | grep -v mocks | grep -v migrations` -cover
test/pkg:
	go test `go list ./pkg/... | grep -v mocks` -cover
mocks:
	mockery --dir ./ --output ./mocks --all --case snake