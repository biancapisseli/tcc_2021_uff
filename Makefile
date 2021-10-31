test:
	go test `go list ./... | grep -v vendor` -cover
mocks:
	cd ./internal/user
	mockery --dir ./internal/user --output ./internal/user/mocks --all --case snake
	cd ..