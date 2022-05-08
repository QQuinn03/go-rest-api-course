package comment

//Persistence Layer responsibility is to provide an interface for accessing the database
import (
	"context"
	"errors"
	"fmt"
)

//the struct on which all our logic will be
//built on top of

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

//is the struct on which all our logic will be built on
//top of
type Service struct {
	Store Store
}

//return a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

//a representation of the comment structure of our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// theis interface
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	PostComment(context.Context, Comment) (Comment, error)
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retriveing a comment...")
	cmt, err := s.Store.GetComment(ctx, id) //triggrt getComment from database comment
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

//// UpdateComment - updates a comment by ID with new comment info
func (s *Service) UpdateComment(ctx context.Context, ID string, UpdateCmt Comment) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, ID, UpdateCmt)
	if err != nil {
		fmt.Println("error updating comment")
		return Comment{}, err
	}
	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, ID string) error {
	return s.Store.DeleteComment(ctx, ID)

}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	insertedCmt, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}
	return insertedCmt, nil
}
