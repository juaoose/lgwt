package context

import (
	"context"
	"fmt"
	"net/http"
)

// https://faiface.github.io/post/context-should-go-away-go2/
// https://go.dev/blog/context

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			// Error should at least be logged
			return
		}

		fmt.Fprint(w, data)
	}
}
