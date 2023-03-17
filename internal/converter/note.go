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

func ToNote(note *desc.Note) *model.Note {
	updatedAt := sql.NullTime{
		Time:  note.GetUpdatedAt().AsTime(),
		Valid: note.GetUpdatedAt().IsValid(),
	}

	return &model.Note{
		Id:        note.GetId(),
		Info:      ToNoteInfo(note.GetNoteInfo()),
		CreatedAt: note.GetCreatedAt().AsTime(),
		UpdatedAt: updatedAt,
	}
}

func ToDescNote(note *model.Note) *desc.Note {
	var updatedAt *timestamppb.Timestamp
	if note.UpdatedAt.Valid {
		updatedAt = timestamppb.New(note.UpdatedAt.Time)
	}

	return &desc.Note{
		Id:        note.Id,
		NoteInfo:  ToDescNoteInfo(note.Info),
		CreatedAt: timestamppb.New(note.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToDescListNote(notes []*model.Note) []*desc.Note {
	res := make([]*desc.Note, 0, len(notes))
	for i := range notes {
		res = append(res, ToDescNote(notes[i]))
	}

	return res
}

func ToUpdateNoteInfo(noteInfo *desc.UpdateNoteInfo) *model.UpdateNoteInfo {
	title := sql.NullString{String: noteInfo.GetTitle().GetValue()}
	if noteInfo.GetTitle() != nil {
		title.Valid = true
	}

	text := sql.NullString{String: noteInfo.GetText().GetValue()}
	if noteInfo.GetText() != nil {
		text.Valid = true
	}

	author := sql.NullString{String: noteInfo.GetAuthor().GetValue()}
	if noteInfo.GetAuthor() != nil {
		author.Valid = true
	}

	email := sql.NullString{String: noteInfo.GetEmail().GetValue()}
	if noteInfo.GetEmail() != nil {
		email.Valid = true
	}

	return &model.UpdateNoteInfo{
		Title:  title,
		Text:   text,
		Author: author,
		Email:  email,
	}
}

func ToDescUpdateNoteInfo(noteInfo *model.UpdateNoteInfo) *desc.UpdateNoteInfo {
	var title, text, author, email *wrapperspb.StringValue

	if noteInfo.Title.Valid {
		title = &wrapperspb.StringValue{Value: noteInfo.Title.String}
	}
	if noteInfo.Text.Valid {
		text = &wrapperspb.StringValue{Value: noteInfo.Text.String}
	}
	if noteInfo.Author.Valid {
		author = &wrapperspb.StringValue{Value: noteInfo.Author.String}
	}
	if noteInfo.Email.Valid {
		email = &wrapperspb.StringValue{Value: noteInfo.Email.String}
	}

	return &desc.UpdateNoteInfo{
		Title:  title,
		Text:   text,
		Author: author,
		Email:  email,
	}
}
