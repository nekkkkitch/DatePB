authpb:
	protoc --proto_path=Auth/proto --go_out=Auth/pb --go-grpc_out=Auth/pb Auth/proto/auth.proto