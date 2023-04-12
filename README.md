<h1>gRPC service for notes with REST API support</h1>
This is a gRPC service for CRUD operations with notes.<br>
Service also handles HTTP requests through grpc-gateway to support REST JSON API.

<h2>Features</h2>
Creating, updating or deleting a Note object will also create or delete a Log object in the database. These operations are <u>transactional</u>.
<br>
The objects structure for requests and responses defined in the <a href=https://github.com/Journeyman150/note-service-api/blob/task6/api/note_v1/note.proto>api/note_v1/note.proto</a> file.
See models of Note and Log entities in <a href=https://github.com/Journeyman150/note-service-api/tree/task6/internal/model>internal/model</a> folder.
<br><br>
Service uses input data validation of Note object. Rules defined in the <a href=https://github.com/Journeyman150/note-service-api/blob/task6/api/note_v1/note.proto>api/note_v1/note.proto</a> file.
<br><br>
All api methods covered by tests (see <a href=https://github.com/Journeyman150/note-service-api/tree/task6/internal/app/api/note_v1>internal/app/api/note_v1</a> folder).