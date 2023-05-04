package handlers_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gnaydenova/url-shortener/app/handlers"
	"github.com/gnaydenova/url-shortener/app/shortener"
	"github.com/gnaydenova/url-shortener/app/storage"
)

func TestURLHandler(t *testing.T) {
	handler := handlers.NewURLHandler(shortener.NewShortener(storage.NewInMemoryMapRepository()))

	t.Run("returns bad request if request has no body", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", http.NoBody)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("wrong status code: got %v, expected %v", rr.Code, http.StatusBadRequest)
		}
	})

	t.Run("returns newly generated short url", func(t *testing.T) {
		body := []byte(`https://someverylongdomainnamehere.com/some/very/very/long/path/here?foo=bar`)
		req, err := http.NewRequest(http.MethodGet, "/", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("wrong status code: got %v, expected %v", rr.Code, http.StatusOK)
		}

		resp := rr.Body.String()

		if resp == "" {
			t.Fatal("got empty body")
		}

		if len(resp) > 5 {
			t.Fatalf("got short url with len > 5: %s", resp)
		}
	})

	t.Run("return not found when trying to resolve an invalid short url", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/invalid", http.NoBody)

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Fatalf("wrong status code: got %v, expected %v", rr.Code, http.StatusNotFound)
		}
	})

	t.Run("redirects when called with valid short url", func(t *testing.T) {
		body := []byte(`https://someotherdomain.com/some/very/very/long/path/here?foo=bar`)
		req, err := http.NewRequest(http.MethodGet, "/", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}

		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		shortened := rec.Body.String()

		resolveReq, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/%s", shortened), http.NoBody)
		if err != nil {
			t.Fatal(err)
		}

		rRec := httptest.NewRecorder()
		handler.ServeHTTP(rRec, resolveReq)

		if rRec.Code != http.StatusMovedPermanently {
			t.Fatalf("wrong status code: got %v, expected %v", rRec.Code, http.StatusMovedPermanently)
		}

		redirectUrl, err := rRec.Result().Location()
		if err != nil {
			t.Fatal(err)
		}
		if redirectUrl.String() != string(body) {
			t.Fatalf("wrong url: got %s, expected %s", redirectUrl.RequestURI(), body)
		}
	})
}
