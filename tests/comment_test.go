//go:build e2e
// +build e2e

package tests

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetComment(t *testing.T) {
	t.Run("can get comment with correct id", func(*testing.T) {
		client := resty.New()
		resp, err := client.R().
			//SetQueryString("f9a838b2-e80c-473e-ba5b-78e02fddc148").
			//SetHeader("Accept", "application/json").
			Get("http://localhost:8080/api/v1/comment/f9a838b2-e80c-473e-ba5b-78e02fddc148")

		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
	})
	t.Run("cannot get comment without correct id", func(*testing.T) {
		client := resty.New()
		resp, err := client.R().
			//SetQueryString("f9a838b2-e80c-473e-ba5b-78e02fddc148").
			//SetHeader("Accept", "application/json").
			Get("http://localhost:8080/api/v1/comment/f9a838b2-e80c-473e-ba5b-78e02fddc138")

		assert.Equal(t, err, nil)
		assert.Equal(t, 500, resp.StatusCode())
	})
}
func TestUpdateComment(t *testing.T) {
	t.Run("can update comment with correct id", func(*testing.T) {
		client := resty.New()
		resp, err := client.R().
			//SetQueryString("f9a838b2-e80c-473e-ba5b-78e02fddc148").
			SetHeader("Authorization", "bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.gKF3xmqVhxwnbKxDdBby1iCUXsKEsMc9UUYel-tk3Do6gOpthLdggBvTuaCTTvu__9d9S3uESxtl3QSEotRlzA").
			SetHeader("Accept", "application/json").
			SetBody(`{"slug": "/", "author": "QQisagoodcoder", "body": "testing this"}`).
			Put("http://localhost:8080/api/v1/comment/f9a838b2-e80c-473e-ba5b-78e02fddc148")

		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
	})
}

// func TestDeleteComment(t *testing.T) {
// 	t.Run("can delete comment", func(*testing.T) {
// 		client := resty.New()
// 		resp, err := client.R().
// 			SetHeader("Authorization", "bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.gKF3xmqVhxwnbKxDdBby1iCUXsKEsMc9UUYel-tk3Do6gOpthLdggBvTuaCTTvu__9d9S3uESxtl3QSEotRlzA").
// 			Delete("http://localhost:8080/api/v1/comment/bd26b0de-8330-4419-b718-e7473661d0c8")
// 		assert.NoError(t, err)
// 		assert.Equal(t, 200, resp.StatusCode())

// 	})

// }

func TestPostComment(t *testing.T) {
	t.Run("can post comment", func(*testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetHeader("Authorization", "bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.gKF3xmqVhxwnbKxDdBby1iCUXsKEsMc9UUYel-tk3Do6gOpthLdggBvTuaCTTvu__9d9S3uESxtl3QSEotRlzA").
			SetBody(`{"slug": "/", "author": "QQ", "body": "testing this"}`).
			Post("http://localhost:8080/api/v1/comment")
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode())
	})

	t.Run("can not post comment withoutjwt", func(*testing.T) {
		client := resty.New()
		resp, err := client.R().
			SetBody(`{"slug": "/", "author": "QQ","body": "testing this"}`).
			Post("http://localhost:8080/api/v1/comment")
		assert.NoError(t, err)
		assert.Equal(t, 401, resp.StatusCode())
	})
}
