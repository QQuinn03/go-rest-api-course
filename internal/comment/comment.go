package comment

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

// the interface for comment service
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	// GetCommentBySlug(Slug string) ([]Comment, error)
	// PostComment(comment Comment) (Comment, error)
	UpdateComment(context.Context, Comment) error
	DeleteComment(context.Context, string) error
	// GetAllComments() ([]Comment, error)
	CreateComment(context.Context, Comment) (Comment, error)
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retriveing a comment...")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

// func (s *Service) GetCommentBySlug(slug string) ([]Comment, error) {
// 	var comment []Comment
// 	if result := s.DB.Find(&comment).Where("slug=?", slug); result != nil {
// 		return []Comment{}, result.Error
// 	}
// 	return comment, nil
// }

// // add new comment in a database

// func (s *Service) PostComment(comment Comment) (Comment, error) {
// 	if result := s.DB.Save(&comment); result.Error != nil {
// 		return Comment{}, result.Error
// 	}
// 	return comment, nil
// }

//// UpdateComment - updates a comment by ID with new comment info
func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	// comment, err := s.GetComment(ID)
	// if err != nil {
	// 	return Comment{}, err
	return ErrNotImplemented
}

// 	if result := s.DB.Model(&comment).Updates(newComment); result.Error != nil {
// 		return Comment{}, result.Error
// 	}

// 	return comment, nil

// }

func (s *Service) DeleteComment(ctx context.Context, ID string) error {
	// if result := s.DB.Delete(&Comment{}, ID); result.Error != nil {
	// 	return result.Error
	// }
	// return nil
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
