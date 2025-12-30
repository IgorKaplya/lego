package errtyps

import (
	"fmt"
	"io"
	"net/http"
)

type BadStatusError struct {
	URL    string
	Status int
}

func (e BadStatusError) Error() string {
	return fmt.Sprintf("did not get %d from %s, got %d", http.StatusOK, e.URL, e.Status)
}

func DumbGetter(url string) (string, error) {
	res, errGet := http.Get(url)
	if errGet != nil {
		return "", fmt.Errorf("problem fetching %q, %v", url, errGet)
	}

	if res.StatusCode != http.StatusOK {
		return "", BadStatusError{URL: url, Status: res.StatusCode}
	}

	defer res.Body.Close()
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		return "", fmt.Errorf("problem reading response, %v", errBody)
	}

	return string(body), nil
}
