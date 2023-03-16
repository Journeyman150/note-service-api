package main

import (
	"context"
	"log"
	"time"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const address = "localhost:50051"

func main() {
	con, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteV1Client(con)

	//create
	res, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		NoteInfo: &desc.NoteInfo{
			Title:  "Title",
			Text:   "Text",
			Author: "Author",
			Email:  "Email",
		},
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Note created")
	log.Println("Note Id:", res.GetId())
	log.Println()

	//get
	id := int64(2)
	get, err := client.GetNote(context.Background(), &desc.GetNoteRequest{
		Id: id,
	})
	if err != nil {
		log.Println(err.Error())
	}
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Note with Id =", id, "received")
	log.Println("Title:", get.GetNote().GetNoteInfo().GetTitle())
	log.Println("Author:", get.GetNote().GetNoteInfo().GetAuthor())
	log.Println("Text:", get.GetNote().GetNoteInfo().GetText())
	log.Println("Created at:", get.GetNote().GetCreatedAt().AsTime().In(loc).Format(time.UnixDate))
	if get.GetNote().GetUpdatedAt().IsValid() {
		log.Println("Updated at:", get.GetNote().GetUpdatedAt().AsTime().In(loc).Format(time.UnixDate))
	} else {
		log.Println("Updated at: Note has never been updated")
	}
	log.Println()

	//getList
	getList, err := client.GetListNote(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Println(err.Error())
	}
	loc, err = time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Println(err.Error())
	}
	if len(getList.Notes) != 0 {
		log.Println("All notes received")
		for i, note := range getList.GetNotes() {
			log.Println("Note", i+1)
			log.Println("Title:", note.GetNoteInfo().GetTitle())
			log.Println("Author:", note.GetNoteInfo().GetAuthor())
			log.Println("Text:", note.GetNoteInfo().GetText())
			log.Println("Created at:", note.GetCreatedAt().AsTime().In(loc).Format(time.UnixDate))
			if note.GetUpdatedAt().IsValid() {
				log.Println("Updated at:", note.GetUpdatedAt().AsTime().In(loc).Format(time.UnixDate))
			} else {
				log.Println("Updated at: Note has never been updated")
			}
			log.Println()
		}
	} else {
		log.Println("No notes found")
	}

	//update
	newTitle := "updated title"
	newText := "updated text"
	_, err = client.Update(context.Background(), &desc.UpdateNoteRequest{
		Id: 3,
		UpdateNoteInfo: &desc.UpdateNoteInfo{
			Title: wrapperspb.String(newTitle),
			Text:  wrapperspb.String(newText),
		},
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Note updated")
	log.Println()

	//delete
	_, err = client.DeleteNote(context.Background(), &desc.DeleteNoteRequest{
		Id: 1,
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Note deleted")
	log.Println()
}
