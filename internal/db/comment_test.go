//go:build integration
// +build integration

package db

import (
	"context"
	"testing"

	"github.com/QQuinn03/go-rest-api-course/internal/comment"
	"github.com/stretchr/testify/assert"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})

		assert.NoError(t, err)

		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "slug", newCmt.Slug)

	})
	//test delete comment
	t.Run("test delete comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "new_slug",
			Author: "qqqqq",
			Body:   "new_bodyqqqq",
		})

		assert.NoError(t, err)

		err = db.DeleteComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		_, err = db.GetComment(context.Background(), cmt.ID)
		assert.Error(t, err)

	})

	t.Run("test update comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "new_slug",
			Author: "qq",
			Body:   "new_body",
		})
		assert.NoError(t, err)

		newCmt, err := db.UpdateComment(context.Background(), cmt.ID, comment.Comment{
			Slug:   "new_slug",
			Author: "qq",
			Body:   "new_body_letstry",
		})
		assert.NoError(t, err)

		updatedCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.Equal(t, newCmt, updatedCmt)

	})

	t.Run("test get comment", func(t *testing.T) {

		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "new_slug_new",
			Author: "qq_new_qq",
			Body:   "new_body_new_body",
		})
		assert.NoError(t, err)

		retrivedCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.Equal(t, cmt, retrivedCmt)
	})

}
