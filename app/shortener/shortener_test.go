package shortener_test

import (
	"testing"

	"github.com/gnaydenova/url-shortener/app/shortener"
	"github.com/gnaydenova/url-shortener/app/storage"
)

func TestShortener(t *testing.T) {
	shortener := shortener.NewShortener(storage.NewInMemoryMapRepository())

	t.Run("can generate a short url with less than 5 characters", func(t *testing.T) {
		url := "https://someverylongdomainnamehere.com/some/very/very/long/path/here?foo=bar"
		shortened, err := shortener.Generate(url)

		if err != nil {
			t.Fatalf("got error %s", err.Error())
		}

		if len(shortened) > 5 {
			t.Fatalf("len of %s is > 5", shortened)
		}
	})

	t.Run("returns the same short url if it was already generated", func(t *testing.T) {
		url := "https://someotherlongdomainnamehere.com/some/path"
		shortened, err := shortener.Generate(url)
		if err != nil {
			t.Fatalf("got error %s", err.Error())
		}

		newShortened, err := shortener.Generate(url)
		if err != nil {
			t.Fatalf("got error %s", err.Error())
		}

		if newShortened != shortened {
			t.Fatalf("expected: %s, got: %s", newShortened, shortened)
		}
	})

	t.Run("can resolve previously generated short url", func(t *testing.T) {
		url := "https://anotherlongdomainnamehere.com/some/long/path"
		shortened, err := shortener.Generate(url)
		if err != nil {
			t.Fatalf("got error %s", err.Error())
		}

		resolvedUrl, err := shortener.Resolve(shortened)
		if err != nil {
			t.Fatalf("got error %s", err.Error())
		}

		if resolvedUrl != url {
			t.Fatalf("expected: %s, got: %s", url, resolvedUrl)
		}
	})
}
