package main

import "testing"

func TestKlipySearchResponseParserFuncReturnsResultFromSearchBatch(t *testing.T) {
	body := `{"data":{"data":[{"file":{"hd":{"gif":{"url":"https://example.com/one.gif"}}}},{"file":{"hd":{"gif":{"url":"https://example.com/two.gif"}}}}]}}`
	allowed := map[string]bool{
		"https://example.com/one.gif": true,
		"https://example.com/two.gif": true,
	}

	for range 20 {
		if got := klipySearchResponseParserFunc(body); !allowed[got] {
			t.Fatalf("parser returned %q, want a URL from the search batch", got)
		}
	}
}

func TestKlipySearchResponseParserFuncReturnsEmptyForNoResults(t *testing.T) {
	if got := klipySearchResponseParserFunc(`{"data":{"data":[]}}`); got != "" {
		t.Fatalf("parser returned %q, want empty result", got)
	}
}

func TestKlipySearchResponseParserFuncDoesNotRepeatLastGIF(t *testing.T) {
	body := `{"data":{"data":[{"file":{"hd":{"gif":{"url":"https://example.com/one.gif"}}}},{"file":{"hd":{"gif":{"url":"https://example.com/two.gif"}}}}]}}`

	lastGifURLMu.Lock()
	lastGifURL = ""
	lastGifURLMu.Unlock()
	defer func() {
		lastGifURLMu.Lock()
		lastGifURL = ""
		lastGifURLMu.Unlock()
	}()

	previous := ""
	for range 20 {
		if got := klipySearchResponseParserFunc(body); got == previous {
			t.Fatalf("parser repeated %q", got)
		} else {
			previous = got
		}
	}
}
