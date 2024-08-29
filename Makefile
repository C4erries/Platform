.PHONY: build
build:
		go build -v .order-service/cmd
		go build -v .product-service/cmd
		go build -v .user-service/cmd
	
.PHONY: proto
proto:
		protoc -I order-service/proto order-service/proto/order.proto --go_out=./order-service/gen/ --go_opt=paths=source_relative --go-grpc_out=./order-service/gen/ --go-grpc_opt=paths=source_relative
		protoc -I product-service/proto product-service/proto/product.proto --go_out=./product-service/gen/ --go_opt=paths=source_relative --go-grpc_out=./product-service/gen/ --go-grpc_opt=paths=source_relative
		protoc -I user-service/proto user-service/proto/user.proto --go_out=./user-service/gen/ --go_opt=paths=source_relative --go-grpc_out=./user-service/gen/ --go-grpc_opt=paths=source_relative

.DEFAULT_GOAL := build