authpb:
	protoc --proto_path=Auth/proto --go_out=Auth/pb --go-grpc_out=Auth/pb Auth/proto/auth.proto
profilerpb:
	protoc --proto_path=Profiler/proto --go_out=Profiler/pb --go-grpc_out=Profiler/pb Profiler/proto/profiler.proto