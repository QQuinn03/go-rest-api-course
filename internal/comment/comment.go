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
//this interface will be implemented by repository layer(sql/datavase layer)
type Store interface {
	GettComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
}

//Service layer does not need to know how implemntation details of how comments get retrived from repository layer
//The GetComment outside is not the same as GetComment inside, inside one is calling method GetComment in the interface
//which will be implemented by mthod GetComment in database layer
func (s *Service) GettComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retriveing a comment...")

	cmt, err := s.StoreInterface.GettComment(ctx, id) //triggrt getComment from database comment
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	fmt.Println("creating a comment...")
	cmt, err := s.StoreInterface.PostComment(ctx, cmt) //triggrt getComment from database comment
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}
	return cmt, nil

}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	fmt.Println("deleting a comment...")
	err := s.StoreInterface.DeleteComment(ctx, id)
	if err != nil {
		fmt.Println("failing to delete the comment", err)
		return err
	}
	return nil
}

func (s *Service) UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error) {
	fmt.Println("Updating comment")
	updatedCmt, err := s.StoreInterface.UpdateComment(ctx, id, cmt)
	if err != nil {
		fmt.Println("failing to update the comment", err)
		return Comment{}, err
	}
	return updatedCmt, nil

}
