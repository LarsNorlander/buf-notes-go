package main

import (
	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"context"
	notesv1 "github.com/LarsNorlander/buf-notes-go/gen/larsnorlander/notes/v1"
	"github.com/LarsNorlander/buf-notes-go/gen/larsnorlander/notes/v1/notesv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net/http"
)

func main() {
	notes := &NotesService{}

	mux := http.NewServeMux()
	path, handler := notesv1connect.NewNoteServiceHandler(notes)

	reflector := grpcreflect.NewStaticReflector(notesv1connect.NoteServiceName)

	mux.Handle(path, handler)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	log.Println("starting server...")
	_ = http.ListenAndServe(":8080", h2c.NewHandler(mux, &http2.Server{}))
}

type NotesService struct {
	notes []*notesv1.Note
}

func (ns *NotesService) CreateNote(
	_ context.Context,
	req *connect.Request[notesv1.CreateNoteRequest],
) (
	*connect.Response[notesv1.CreateNoteResponse],
	error,
) {
	req.Msg.Note.CreatedTime = timestamppb.Now()
	ns.notes = append(ns.notes, req.Msg.Note)
	return connect.NewResponse(&notesv1.CreateNoteResponse{Note: req.Msg.Note}), nil
}

func (ns *NotesService) ListNotes(
	_ context.Context,
	_ *connect.Request[notesv1.ListNotesRequest],
) (
	*connect.Response[notesv1.ListNotesResponse],
	error,
) {
	return connect.NewResponse(&notesv1.ListNotesResponse{Notes: ns.notes}), nil
}
