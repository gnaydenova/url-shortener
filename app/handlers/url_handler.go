package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gnaydenova/url-shortener/app/shortener"
)

type URLHandler struct {
	shortener *shortener.Shortener
}

func NewURLHandler(shortener *shortener.Shortener) *URLHandler {
	return &URLHandler{shortener}
}

func (h *URLHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		h.shorten(w, req)
		return
	}

	h.resolve(w, req)
}

func (h *URLHandler) shorten(w http.ResponseWriter, req *http.Request) {
	body, _ := io.ReadAll(req.Body)
	uri := string(body)

	if _, err := url.ParseRequestURI(uri); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := io.WriteString(w, "Invalid URL"); err != nil {
			fmt.Println(err.Error())
		}
		return
	}

	shortened, err := h.shortener.Generate(uri)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := io.WriteString(w, shortened); err != nil {
		fmt.Println(err.Error())
	}
}

func (h *URLHandler) resolve(w http.ResponseWriter, req *http.Request) {
	path := strings.TrimLeft(req.URL.Path, "/")
	url, err := h.shortener.Resolve(path)

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, req, url, http.StatusMovedPermanently)
}
