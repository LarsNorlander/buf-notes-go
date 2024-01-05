package main

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	notesv1 "github.com/LarsNorlander/buf-notes-go/gen/larsnorlander/notes/v1"
	"github.com/LarsNorlander/buf-notes-go/gen/larsnorlander/notes/v1/notesv1connect"
	"log"
	"net/http"
)

func main() {
	client := notesv1connect.NewNoteServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)
	res, err := client.ListNotes(context.Background(), connect.NewRequest(&notesv1.ListNotesRequest{}))
	if err != nil {
		log.Fatalln(err)
	}
	for _, note := range res.Msg.Notes {
		fmt.Println(note.Title)
		fmt.Println(note.Body)
		fmt.Println(note.CreatedTime.AsTime())
		fmt.Println("---")
	}
}
