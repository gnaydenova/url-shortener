package storage

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type URL struct {
	ShortURL    string `bson:"short_url"`
	OriginalURL string `bson:"original_url"`
}

// URLRepository represents a repository for storing and retrieving shortened URLs.
type URLRepository interface {
	InsertOne(url URL) error
	FindOneByShortURL(url string) (URL, error)
	FindOneByOriginalURL(url string) (URL, error)
}

// MongoURLRepository is a mongodb implementation of URLRepository.
type MongoURLRepository struct {
	collection *mongo.Collection
}

func NewMongoURLRepository(collection *mongo.Collection) URLRepository {
	return &MongoURLRepository{collection}
}

func (r *MongoURLRepository) InsertOne(url URL) error {
	_, err := r.collection.InsertOne(context.Background(), url)
	return err
}

func (r *MongoURLRepository) FindOneByShortURL(url string) (URL, error) {
	return r.findOne(bson.D{{Key: "short_url", Value: url}})
}

func (r *MongoURLRepository) FindOneByOriginalURL(url string) (URL, error) {
	return r.findOne(bson.D{{Key: "original_url", Value: url}})
}

func (r *MongoURLRepository) findOne(filter interface{}) (URL, error) {
	var result URL

	err := r.collection.
		FindOne(context.Background(), filter).
		Decode(&result)

	return result, err
}

// InMemoryMapRepository is a in memory implementation of URLRepository for testing purposes.
type InMemoryMapRepository struct {
	records map[string]string
}

func NewInMemoryMapRepository() *InMemoryMapRepository {
	r := &InMemoryMapRepository{}
	r.records = make(map[string]string)
	return r
}

func (r *InMemoryMapRepository) InsertOne(url URL) error {
	r.records[url.ShortURL] = url.OriginalURL
	return nil
}

func (r *InMemoryMapRepository) FindOneByShortURL(url string) (URL, error) {
	var err error
	originalUrl, ok := r.records[url]
	if !ok {
		err = errors.New("not found")
	}

	return URL{OriginalURL: originalUrl, ShortURL: url}, err
}

func (r *InMemoryMapRepository) FindOneByOriginalURL(url string) (URL, error) {
	var err error
	var shortURL string

	for short, val := range r.records {
		if val == url {
			shortURL = short
			break
		}
	}

	if shortURL == "" {
		err = errors.New("not found")
	}

	return URL{OriginalURL: url, ShortURL: shortURL}, err
}
