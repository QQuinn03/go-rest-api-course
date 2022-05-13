package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/QQuinn03/go-rest-api-course/internal/comment"
)

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func convertCommentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID:     c.ID,
		Slug:   c.Slug.String,
		Author: c.Author.String,
		Body:   c.Body.String,
	}
}

func (d *Database) GettComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var cmt CommentRow
	row := d.Client.QueryRowxContext(ctx,
		`SELECT id, slug,body, author from comments where id=$1`,
		uuid)
	err := row.Scan(&cmt.ID, &cmt.Slug, &cmt.Body, &cmt.Author)
	if err != nil {
		return comment.Comment{}, err
	}
	return convertCommentRowToComment(cmt), nil

}
func (d *Database) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {

	cmtrow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}

	insertRow, err := d.Client.NamedQueryContext(ctx, `INSERT INTO comments(id,slug,body,author) values(:id,:slug,:body,:author`, cmtrow)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert comment: %w", err)
	}
	if err := insertRow.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close table: %w", err)
	}
	return cmt, nil

}
func (d *Database) DeleteComment(ctx context.Context, uuid string) error {
	_, err := d.Client.NamedQueryContext(ctx, `DELETE FROM comments where is=$1`, uuid)
	if err != nil {
		return fmt.Errorf("failto delete the comment:%v", err)

	}
	return nil

}

func (d *Database) UpdateComment(ctx context.Context, uuid string, cmt comment.Comment) (comment.Comment, error) {
	cmtRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}
	updateRow, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE comments SET
		 id =:id,
		 slug =:slug,
		 body =:body,
		 author =:author`, cmtRow)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("fail to update comments:%w", err)
	}
	if err := updateRow.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close table: %w", err)
	}
	return convertCommentRowToComment(cmtRow), nil

}
