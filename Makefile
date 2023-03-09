PHONY: generate
generate:
		if not exist pkg\note_v1 mkdir pkg\note_v1
		protoc 	--proto_path=api\note_v1 --go_opt=paths=source_relative --go_out=pkg\note_v1  --go-grpc_opt=paths=source_relative --go-grpc_out=pkg\note_v1 api\note_v1\note.proto