.ONESHELL:

protogen:
	cd api
	protoc \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=. --grpc-gateway_opt paths=source_relative \
		fees.proto

run:
	docker-compose build
	docker-compose up

mock:
	mockgen -destination pkg/fees/usecases/mocks/ethRepositoryMock.go \
		github.com/rsmarincu/glassnode/pkg/fees/usecases ETHRepository
	mockgen -destination pkg/fees/grpc/mocks/feesServiceMock.go \
    		github.com/rsmarincu/glassnode/pkg/fees/grpc FeesService