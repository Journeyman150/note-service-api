gRPC service for CRUD operations with notes\
Service also handles HTTP requests through grpc-gateway to support RESTful JSON API

Creating, updating, or deleting a Note object will also create or delete a Log object in the database. These operations are transactional.
See models of Note and Log entities in internal/model folder.

Service uses Note object input validation. The rules are defined in the api/note_v1/note.proto file.

All api methods covered by tests (see internal/app/api/note_v1 folder).
