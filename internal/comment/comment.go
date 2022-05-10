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
	StoreInterface Store
}

//return a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		StoreInterface: store,
	}
}

//a representation of the comment structure of our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// this interface defines all methods that our service needs in order to operate
type Store interface {
	getComment(context.Context, string) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	PostComment(context.Context, Comment) (Comment, error)
}

//Service layer does not need to know how implemntation details of how comments get retrived from repository layer
func (s *Service) getComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retriveing a comment...")

	cmt, err := s.StoreInterface.getComment(ctx, id) //triggrt getComment from database comment
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

//// UpdateComment - updates a comment by ID with new comment info
func (s *Service) UpdateComment(ctx context.Context, ID string, UpdateCmt Comment) (Comment, error) {
	cmt, err := s.StoreInterface.UpdateComment(ctx, ID, UpdateCmt)
	if err != nil {
		fmt.Println("error updating comment")
		return Comment{}, err
	}
	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, ID string) error {
	return s.StoreInterface.DeleteComment(ctx, ID)

}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	insertedCmt, err := s.StoreInterface.PostComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}
	return insertedCmt, nil
}
