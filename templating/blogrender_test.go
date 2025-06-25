package blogrender_test

import (
	"bytes"
	"io"
	"testing"

	blogrender "github.com/IgorKaplya/lego/templating"
	approvals "github.com/approvals/go-approval-tests"
)

func TestRenderPost(t *testing.T) {
	testCases := []struct {
		description string
		post        blogrender.Post
	}{
		{
			description: "post",
			post: blogrender.Post{
				Title:       "hello world",
				Description: "This is a description",
				Tags:        []string{"go", "tdd"},
				Body:        "# This is a post",
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			var buf = bytes.Buffer{}
			var err = blogrender.RenderPost(&buf, testCase.post)

			assertError(t, err)
			approvals.VerifyString(t, buf.String())
		})
	}
}

func BenchmarkRenderPost(b *testing.B) {
	var post = blogrender.Post{
		Title:       "hello world",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
		Body:        "This is a post",
	}
	for b.Loop() {
		blogrender.RenderPost(io.Discard, post)
	}
}

func TestIndex(t *testing.T) {
	testCases := []struct {
		desc  string
		given []blogrender.Post
		want  string
	}{
		{
			desc: "posts",
			given: []blogrender.Post{
				{Title: "Title 1"},
				{Title: "Title 2"},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			var buf = bytes.Buffer{}

			var err = blogrender.RenderIndex(&buf, testCase.given)

			assertError(t, err)
			approvals.VerifyString(t, buf.String())
		})
	}
}

func assertError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
