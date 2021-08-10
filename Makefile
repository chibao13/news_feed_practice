# START: compile
compile:
	protoc api/v1/*.proto api/v1/userfriend/*.proto \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.
# END: compile