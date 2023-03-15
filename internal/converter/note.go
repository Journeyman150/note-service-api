package converter

import (
	"database/sql"

	"github.com/Journeyman150/note-service-api/internal/model"
	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ToNoteInfo(noteInfo *desc.NoteInfo) *model.NoteInfo {
	return &model.NoteInfo{
		Title:  noteInfo.GetTitle(),
		Text:   noteInfo.GetText(),
		Author: noteInfo.GetAuthor(),
		Email:  noteInfo.GetEmail(),
	}
}

func ToDescNoteInfo(noteInfo *model.NoteInfo) *desc.NoteInfo {
	return &desc.NoteInfo{
		Title:  noteInfo.Title,
		Text:   noteInfo.Text,
		Author: noteInfo.Author,
		Email:  noteInfo.Email,
	}
}

func ToGetNoteResponse(getNoteResponse *desc.GetNoteResponse) *model.GetNoteResponse {
	updatedAt := sql.NullTime{
		Time:  getNoteResponse.GetUpdatedAt().AsTime(),
		Valid: getNoteResponse.GetUpdatedAt().IsValid(),
	}

	return &model.GetNoteResponse{
		Id:        getNoteResponse.GetId(),
		Info:      ToNoteInfo(getNoteResponse.GetNoteInfo()),
		CreatedAt: getNoteResponse.GetCreatedAt().AsTime(),
		UpdatedAt: updatedAt,
	}
}

func ToDescGetNoteResponse(getNoteResponse *model.GetNoteResponse) *desc.GetNoteResponse {
	var updatedAt *timestamppb.Timestamp
	if getNoteResponse.UpdatedAt.Valid {
		updatedAt = timestamppb.New(getNoteResponse.UpdatedAt.Time)
	}

	return &desc.GetNoteResponse{
		Id:        getNoteResponse.Id,
		NoteInfo:  ToDescNoteInfo(getNoteResponse.Info),
		CreatedAt: timestamppb.New(getNoteResponse.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToDescGetListNoteResponse(getListNoteResponse []*model.GetNoteResponse) *desc.GetListNoteResponse {
	res := make([]*desc.GetNoteResponse, 0, len(getListNoteResponse))
	for i := range getListNoteResponse {
		res = append(res, ToDescGetNoteResponse(getListNoteResponse[i]))
	}

	return &desc.GetListNoteResponse{Notes: res}
}

func ToUpdateNoteRequest(updateNoteRequest *desc.UpdateNoteRequest) *model.UpdateNoteRequest {
	title := sql.NullString{String: updateNoteRequest.GetTitle().GetValue()}
	if updateNoteRequest.GetTitle() != nil {
		title.Valid = true
	}

	text := sql.NullString{String: updateNoteRequest.GetText().GetValue()}
	if updateNoteRequest.GetText() != nil {
		text.Valid = true
	}

	author := sql.NullString{String: updateNoteRequest.GetAuthor().GetValue()}
	if updateNoteRequest.GetAuthor() != nil {
		author.Valid = true
	}

	email := sql.NullString{String: updateNoteRequest.GetEmail().GetValue()}
	if updateNoteRequest.GetEmail() != nil {
		email.Valid = true
	}

	return &model.UpdateNoteRequest{
		Id:     updateNoteRequest.GetId(),
		Title:  title,
		Text:   text,
		Author: author,
		Email:  email,
	}
}

func ToDescUpdateNoteRequest(updateNoteRequest *model.UpdateNoteRequest) *desc.UpdateNoteRequest {
	var title, text, author, email *wrapperspb.StringValue

	if updateNoteRequest.Title.Valid {
		title = &wrapperspb.StringValue{Value: updateNoteRequest.Title.String}
	}
	if updateNoteRequest.Text.Valid {
		text = &wrapperspb.StringValue{Value: updateNoteRequest.Text.String}
	}
	if updateNoteRequest.Author.Valid {
		author = &wrapperspb.StringValue{Value: updateNoteRequest.Author.String}
	}
	if updateNoteRequest.Email.Valid {
		email = &wrapperspb.StringValue{Value: updateNoteRequest.Email.String}
	}

	return &desc.UpdateNoteRequest{
		Id:     updateNoteRequest.Id,
		Title:  title,
		Text:   text,
		Author: author,
		Email:  email,
	}
}
