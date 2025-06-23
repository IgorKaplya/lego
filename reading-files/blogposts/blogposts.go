package blogposts

import (
	"bufio"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFS(fileSystem fs.FS) (result []Post, error error) {
	var dir, readError = fs.ReadDir(fileSystem, ".")
	if readError != nil {
		return nil, readError
	}

	for _, dirEntry := range dir {
		var post, getPostError = getPost(fileSystem, dirEntry.Name())
		if getPostError != nil {
			return nil, getPostError
		}

		result = append(result, post)
	}

	return
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	var file, openError = fileSystem.Open(fileName)
	if openError != nil {
		return Post{}, openError
	}
	defer file.Close()

	return newPost(file)
}

func newPost(reader io.Reader) (Post, error) {
	var scanner = bufio.NewScanner(reader)

	scanner.Scan()
	var titleLine = scanner.Text()

	scanner.Scan()
	var descriptionLine = scanner.Text()

	scanner.Scan()
	var tagsLine = scanner.Text()

	scanner.Scan()

	var bodyLines []string
	for scanner.Scan() {
		bodyLines = append(bodyLines, scanner.Text())
	}

	return Post{
			Title:       titleLine[7:],
			Description: descriptionLine[13:],
			Tags:        strings.Split(tagsLine[6:], ", "),
			Body:        strings.Join(bodyLines, "\n")},
		nil
}
