package shortener

import (
	"math/rand"

	"github.com/gnaydenova/url-shortener/app/base62"
	"github.com/gnaydenova/url-shortener/app/storage"
)

const (
	// max is the number of possible permutations with a maximum length of 5 chars (62 to the power of 5)
	max int = 916132832
)

type Shortener struct {
	repo storage.URLRepository
}

func NewShortener(repo storage.URLRepository) *Shortener {
	return &Shortener{repo}
}

func (s *Shortener) Generate(url string) (string, error) {
	if result, err := s.repo.FindOneByOriginalURL(url); err == nil {
		return result.ShortURL, nil
	}

	shortened := s.getShortURL(url)
	err := s.repo.InsertOne(storage.URL{ShortURL: shortened, OriginalURL: url})

	return shortened, err
}

func (s *Shortener) Resolve(short string) (string, error) {
	result, err := s.repo.FindOneByShortURL(short)

	return result.OriginalURL, err
}

func (s *Shortener) getShortURL(url string) string {
	// ensure we don't try to encode 0
	shortened := base62.EncodeToString(uint(rand.Intn(max-1) + 1))

	// check if shortened already exists, if it exists try to generate a new one
	_, err := s.repo.FindOneByShortURL(shortened)
	if err == nil {
		return s.getShortURL(url)
	}

	return shortened
}
