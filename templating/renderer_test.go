package templating

import (
	"bytes"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = Post{
			Title:       "hello",
			Body:        "This is me",
			Description: "Desc",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello</h1><p>Desc</p>Tags: <ul><li>go</li><li>tdd</li></ul>`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
