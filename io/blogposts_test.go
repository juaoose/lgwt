package blogposts

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tag1, tag2
---
Cat
Dog`
		secondBody = `Title: Post 2
Description: Description 2
Tags: tag2, tag3
---
one
two
three`
	)

	fs := fstest.MapFS{
		"example.md":  {Data: []byte(firstBody)},
		"example1.md": {Data: []byte(secondBody)},
	}

	posts, err := NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	assertPort(t, posts[0], Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tag1", "tag2"},
		Body: `Cat
Dog`,
	})
}

func assertPort(t *testing.T, got Post, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
