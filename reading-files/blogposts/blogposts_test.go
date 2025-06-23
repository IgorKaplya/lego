package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/IgorKaplya/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		post1 = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		post2 = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)
	var fs = fstest.MapFS{
		"hello world.md":  {Data: []byte(post1)},
		"hello-world2.md": {Data: []byte(post2)},
	}
	var want = []blogposts.Post{
		{Title: "Post 1", Description: "Description 1", Tags: []string{"tdd", "go"}, Body: "Hello\nWorld"},
		{Title: "Post 2", Description: "Description 2", Tags: []string{"rust", "borrow-checker"}, Body: "B\nL\nM"},
	}
	var got, err = blogposts.NewPostsFromFS(fs)

	assertError(t, err)
	assertPosts(t, got, want)
}

func assertPosts(t *testing.T, got []blogposts.Post, want []blogposts.Post) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func assertError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
