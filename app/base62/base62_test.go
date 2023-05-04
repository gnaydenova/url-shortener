package base62_test

import (
	"testing"

	"github.com/gnaydenova/url-shortener/app/base62"
)

func TestEncodeToString(t *testing.T) {
	t.Run("can encode number to a base62 string", func(t *testing.T) {
		number := uint(1999999999)
		expected := "VMnLB2"

		encoded := base62.EncodeToString(number)
		if encoded != expected {
			t.Fatalf("expected %s, got %s", expected, encoded)
		}
	})
}
