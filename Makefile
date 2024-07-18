gen:
	protoc --proto_path=. --go_out=config --go_opt=paths=source_relative \
		--go-grpc_out=config --go-grpc_opt=require_unimplemented_servers=false --go-grpc_opt=paths=source_relative config.proto
	protoc --proto_path=. --go_out=library --go_opt=paths=source_relative \
		--go-grpc_out=library --go-grpc_opt=require_unimplemented_servers=false --go-grpc_opt=paths=source_relative library.proto
	protoc --proto_path=. --go_out=transaction --go_opt=paths=source_relative \
		--go-grpc_out=transaction --go-grpc_opt=require_unimplemented_servers=false --go-grpc_opt=paths=source_relative transaction.proto
	protoc-go-tags --dir=config
	protoc-go-tags --dir=library
	protoc-go-tags --dir=transaction




	