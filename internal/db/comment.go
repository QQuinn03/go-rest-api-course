package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/QQuinn03/go-rest-api-course/internal/comment"
	uuid "github.com/satori/go.uuid"
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

func (d *Database) getComment(
	ctx context.Context,
	uuid string,
) (comment.Comment, error) {
	var cmtRow CommentRow
	_, err := d.Client.ExecContext(ctx, "SELECT pg_sleep(3)")
	if err != nil {
		return comment.Comment{}, err
	}

	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, slug,body,author
		FROM comments
		WHERE id = $1`,
		uuid,
	)
	err = row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)
	if err != nil {
		return comment.Comment{},
			fmt.Errorf("error fetching the comment uuid")
	}

	return convertCommentRowToComment(cmtRow), nil

}

func (d *Database) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.ID = uuid.NewV4().String()
	postRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
	}
	/* NamedQueryContext Loop through rows using only one struct*/
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO comments
	 (id, slug, author,body)
	 VALUES
	 (:id, :slug, :author, :body)`,
		postRow,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert comment: %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return cmt, nil

}

func (d *Database) DeleteComment(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE from comments where
		 id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("deletion fails: %w", err)
	}
	return nil

}

func (d *Database) UpdateComment(ctx context.Context,
	id string,
	cmt comment.Comment,
) (comment.Comment, error) {

	cmtRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE comments SET
		 slug = :slug,
		 body = :body,
		 author = :author
		 where id = :id`,
		cmtRow,
	)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("updates fails: %w", err)
	}
	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("close fails: %w", err)
	}

	return convertCommentRowToComment(cmtRow), nil
}
