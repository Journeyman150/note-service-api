package main

import (
	"log"
	"context"
	"google.golang.org/grpc"
	desc "github.com/Journeyman150/note-service-api/pkg/note_v1"
)

const address = "localhost:50051"

func main() {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteV1Client(con)
	res, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{
		Title:  "Some title",
		Text:   "Some text",
		Author: "Some author",
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Client created")
	log.Printf("Client Id: %d\n\n", res.Id)

	id := int64(1)
	get, err := client.GetNote(context.Background(), &desc.GetNoteRequest{
		Id:  id,
	})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Client with Id =", id, "received")
	log.Println("Title:", get.Title)
	log.Println("Author:", get.Author)
	log.Println("Text:", get.Text)
}
