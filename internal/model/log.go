package model

import "time"

type Log struct {
	Id        int64     `db:"id"`
	NoteId    int64     `db:"note_id"`
	Msg       string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
}
