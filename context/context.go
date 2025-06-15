package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(context context.Context) string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data = store.Fetch(r.Context())
		fmt.Fprint(w, data)
	}
}
