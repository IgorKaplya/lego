package concurrency

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return !strings.Contains(url, "waka.waka")
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http:\\google.com",
		"http:\\blog.gypsydave5.com",
		"waka:\\waka.waka",
	}

	want := map[string]bool{
		"http:\\google.com":          true,
		"http:\\blog.gypsydave5.com": true,
		"waka:\\waka.waka":           false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

func slowWebsiteCheckerStub(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 1000)

	for b.Loop() {
		CheckWebsites(slowWebsiteCheckerStub, urls)
	}
}
