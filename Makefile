PHONY: generate
generate:
		if not exist pkg\note_v1 mkdir pkg\note_v1
		protoc 	--proto_path=api\note_v1 --go_out=pkg\note_v1 --go_opt=paths=import --go-grpc_out=pkg\note_v1 --go-grpc_opt=paths=import api\note_v1\note.proto
		move /Y pkg\note_v1\github.com\Journeyman150\note-service-api\pkg\note_v1\* pkg\note_v1
		rmdir /s /q pkg\note_v1\github.com
