package http

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
)

func TestFetchArticleSingleflight(t *testing.T) {
	var hits atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		io.WriteString(w, "<html>ok</html>")
	}))
	t.Cleanup(server.Close)

	client := NewBleepingComputerClient(server.Client())
	url := server.URL + "/article"

	var wg sync.WaitGroup
	errCh := make(chan error, 2)
	bodies := make(chan []byte, 2)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			body, err := client.FetchArticle(context.Background(), url)
			if err != nil {
				errCh <- err
				return
			}
			bodies <- body
		}()
	}

	wg.Wait()
	close(errCh)
	close(bodies)

	for err := range errCh {
		t.Fatalf("FetchArticle err = %v, want nil", err)
	}

	if got := hits.Load(); got != 1 {
		t.Errorf("origin GETs = %d, want 1", got)
	}

	for body := range bodies {
		if string(body) != "<html>ok</html>" {
			t.Errorf("body = %q, want %q", body, "<html>ok</html>")
		}
	}
}
