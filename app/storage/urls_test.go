package storage

import (
	"testing"
)

func TestInMemoryMapRepositoryInsertOne(t *testing.T) {
	t.Run("can save a URL", func(t *testing.T) {
		r := NewInMemoryMapRepository()
		short, original := "abcde", "https://somedomain.com/some/path/here?foo=bar"
		err := r.InsertOne(URL{ShortURL: short, OriginalURL: original})

		if err != nil {
			t.Fatalf("got err %s", err.Error())
		}

		if l := len(r.records); l != 1 {
			t.Fatalf("expected 1 record, got %d", l)
		}

		if r.records[short] != original {
			t.Fatalf("record not found in %v", r.records)
		}
	})
}

func TestInMemoryMapRepositoryFindOneByShortURL(t *testing.T) {
	t.Run("returns not found when trying to retrieve nonexistent records", func(t *testing.T) {
		r := NewInMemoryMapRepository()

		_, err := r.FindOneByShortURL("aaaaaa")

		if err.Error() != "not found" {
			t.Fatalf("expected not found error, got %s", err.Error())
		}
	})

	t.Run("can retrieve a record by short URL", func(t *testing.T) {
		r := NewInMemoryMapRepository()
		short, original := "aabbcc", "https://someotherdomain.com/some/path/here?foo=bar"
		r.records[short] = original

		url, err := r.FindOneByShortURL(short)

		if err != nil {
			t.Fatalf("got err %s", err.Error())
		}

		if url.ShortURL != short {
			t.Fatalf("wrong short url: expected %s, got %s", short, url.ShortURL)
		}

		if url.OriginalURL != original {
			t.Fatalf("wrong original url: expected %s, got %s", original, url.OriginalURL)
		}
	})
}

func TestInMemoryMapRepositoryFindOneByOriginalURL(t *testing.T) {
	t.Run("returns not found when trying to retrieve nonexistent records", func(t *testing.T) {
		r := NewInMemoryMapRepository()

		_, err := r.FindOneByOriginalURL("https://domain.com/some/path/here?foo=bar")

		if err.Error() != "not found" {
			t.Fatalf("expected not found error, got %s", err.Error())
		}
	})

	t.Run("can retrieve a record by original URL", func(t *testing.T) {
		r := NewInMemoryMapRepository()
		short, original := "aaffffff", "https://anotherdomain.com/some/path/here?foo=bar"
		r.records[short] = original

		url, err := r.FindOneByOriginalURL(original)

		if err != nil {
			t.Fatalf("got err %s", err.Error())
		}

		if url.ShortURL != short {
			t.Fatalf("wrong short url: expected %s, got %s", short, url.ShortURL)
		}

		if url.OriginalURL != original {
			t.Fatalf("wrong original url: expected %s, got %s", original, url.OriginalURL)
		}
	})
}
