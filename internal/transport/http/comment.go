package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/QQuinn03/go-rest-api-course/internal/comment"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type CommentService interface {
	PostComment(context.Context, comment.Comment) (comment.Comment, error)
	GettComment(ctx context.Context, ID string) (comment.Comment, error)
	DeleteComment(ctx context.Context, ID string) error
	UpdateComment(ctx context.Context, ID string, newCmt comment.Comment) (comment.Comment, error)
}

type Response struct {
	message string
}

//uuid schema
type PostCommentRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Author string `json:"author" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

func convertPostCommentToComment(c PostCommentRequest) comment.Comment {
	return comment.Comment{
		Slug:   c.Slug,
		Author: c.Author,
		Body:   c.Body,
	}

}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {

	var cmt PostCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	validate := validator.New()
	err := validate.Struct(cmt)
	if err != nil {
		http.Error(w, "not a validate comment", http.StatusBadRequest)
		return
	}

	convertedComment := convertPostCommentToComment(cmt)
	postedComment, err := h.Service.PostComment(r.Context(), convertedComment)
	if err != nil {
		log.Print(err)
		return
	}
	if err := json.NewEncoder(w).Encode(postedComment); err != nil {
		panic(err)
	}

}
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cmt, err := h.Service.GettComment(r.Context(), id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}

}
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteComment(r.Context(), id); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{message: "successfully deleted"}); err != nil {
		panic(err)
	}

}
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newCmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&newCmt); err != nil {
		return
	}

	newCmt, err := h.Service.UpdateComment(r.Context(), id, newCmt)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(newCmt); err != nil {
		panic(err)
	}

}
