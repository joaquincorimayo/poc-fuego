package adapter

import (
	"context"
	"github.com/joaquincorimayo/poc-fuego/domain/model"
	"github.com/joaquincorimayo/poc-fuego/domain/ports"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type mongoRepository struct {
	client  *mongo.Client
	db      string
	timeout time.Duration
}

func NewMongoRepository(mongoServerURL, mongoDb string, timeout int) (ports.Repository, error) {
	mongoClient, err := newMongoClient(mongoServerURL)
	duration := time.Duration(timeout) * time.Second
	repo := &mongoRepository{
		client:  mongoClient,
		db:      mongoDb,
		timeout: duration,
	}

	if err != nil {
		return nil, errors.Wrap(err, "mongo client error")
	}

	return repo, nil

}

func newMongoClient(mongoServerURL string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoServerURL)
	duration := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}

func (m mongoRepository) FindByCode(code string) (*model.Report, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	report := &model.Report{}
	collection := m.client.Database(m.db).Collection("reports")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(report)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("Error finding")
		}
		return nil, errors.Wrap(err, "repository research")
	}

	return report, nil
}

func (m mongoRepository) Save(report *model.Report) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	collection := m.client.Database(m.db).Collection("reports")
	_, err := collection.InsertOne(
		ctx,
		report,
	)

	if err != nil {
		return errors.Wrap(err, "Error writing to repository")
	}

	return nil
}
