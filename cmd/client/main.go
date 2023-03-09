package main

import (
	"context"
	"log"

	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteV1Client(con)

	//create
	res, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "Some title",
		Text:   "Some text",
		Author: "Some author",
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Note created")
	log.Println("Note Id:", res.GetId())
	log.Println()

	//get
	id := int64(1)
	get, err := client.GetNote(context.Background(), &desc.GetNoteRequest{
		Id: id,
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Note with Id =", id, "received")
	log.Println("Title:", get.GetTitle())
	log.Println("Author:", get.GetAuthor())
	log.Println("Text:", get.GetText())
	log.Println()

	//getList
	getList, err := client.GetListNote(context.Background(), &desc.GetListNoteRequest{})
	if err != nil {
		log.Println(err.Error())
	}
	if len(getList.Notes) != 0 {
		log.Println("All notes received")
		for i, note := range getList.Notes {
			log.Println("Note", i+1)
			log.Println("Title:", note.GetTitle())
			log.Println("Author:", note.GetAuthor())
			log.Println("Text:", note.GetText())
			log.Println()
		}
	} else {
		log.Println("No notes found")
	}

	//update
	newTitle := "updated title"
	newText := "updated text"
	_, err = client.Update(context.Background(), &desc.UpdateNoteRequest{
		Id:    1,
		Title: &newTitle,
		Text:  &newText,
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
