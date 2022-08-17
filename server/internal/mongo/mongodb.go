package mongo

import (
	"backViewer/internal/config"
	"backViewer/pkg/entity"
	"backViewer/pkg/handlers"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	emailsDB         = "emails"
	emailsCollection = "emails"
)

var (
	ErrObjectIDTypeConversion = errors.New("object id type conversion")
)

// Verify Interface Compliance
var _ handlers.EmailRepo = (*EmailRepo)(nil)

type EmailRepo struct {
	Client *mongo.Client
}

func NewConn(cfg config.MongoDB) (*EmailRepo, error) {
	client, err := mongo.NewClient(
		options.Client().ApplyURI(cfg.URI).
			SetAuth(options.Credential{
				Username: cfg.User,
				Password: cfg.Password,
			}))
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	// if healthy
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return &EmailRepo{
		Client: client,
	}, nil
}

func (er *EmailRepo) Create(ctx context.Context, email *entity.Email) (*entity.Email, error) {

	collection := er.Client.Database(emailsDB).Collection(emailsCollection)

	result, err := collection.InsertOne(ctx, email, &options.InsertOneOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "InsertOne")
	}

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.Wrap(ErrObjectIDTypeConversion, "email.InsertedID")
	}

	email.ID = objectID

	return email, nil
}

func (er *EmailRepo) GetEmailByID(ctx context.Context, id string) (*entity.Email, error) {
	collection := er.Client.Database(emailsDB).Collection(emailsCollection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(ErrObjectIDTypeConversion, "email.GetID")
	}
	filter := bson.M{"_id": objectID}

	var email entity.Email

	if err := collection.FindOne(ctx, filter).Decode(&email); err != nil {
		return nil, err
	}

	return &email, err
}

func (er *EmailRepo) GetEmailByAddress(ctx context.Context, address string) (*entity.Email, error) {
	collection := er.Client.Database(emailsDB).Collection(emailsCollection)
	filter := bson.M{"address": address}
	var email entity.Email
	err := collection.FindOne(ctx, filter).Decode(&email)
	return &email, err
}

func (er *EmailRepo) GetAll(ctx context.Context) ([]entity.Email, error) {
	collection := er.Client.Database(emailsDB).Collection(emailsCollection)

	emails := make([]entity.Email, 0)

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)
	for cursor.Next(ctx) {
		var email entity.Email
		if err = cursor.Decode(&email); err != nil {
			return nil, err
		}
		emails = append(emails, email)
	}

	return emails, nil
}

func (er *EmailRepo) AddOne(ctx context.Context, email *entity.Email) (*entity.Email, error) {
	collection := er.Client.Database(emailsDB).Collection(emailsCollection)
	result, err := collection.InsertOne(ctx, email)
	if err != nil {
		return nil, err
	}

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.Wrap(ErrObjectIDTypeConversion, "report.InsertedID")
	}
	email.ID = objectID

	return email, err
}

func (er *EmailRepo) RemoveOne(ctx context.Context, id string) error {

	collection := er.Client.Database(emailsDB).Collection(emailsCollection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(ErrObjectIDTypeConversion, "email.RemoveByID")
	}

	filter := bson.M{"_id": objectID}

	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return err
}
