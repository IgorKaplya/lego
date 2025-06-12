package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Ping(url string) chan struct{} {
	var result = make(chan struct{})
	go func() {
		http.Get(url)
		close(result)
	}()
	return result
}

func Racer(urlA, urlB string) (string, error) {
	return ConfigurableRacer(urlA, urlB, 10*time.Second)
}

func ConfigurableRacer(urlA, urlB string, timeout time.Duration) (string, error) {
	select {
	case <-Ping(urlA):
		return urlA, nil
	case <-Ping(urlB):
		return urlB, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %q and %q", urlA, urlB)
	}
}
