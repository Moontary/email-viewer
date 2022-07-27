package mongo

import (
	"backViewer/internal/model"
	"context"
	"time"

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

type NewEmailRepo struct {
	client   *mongo.Client
	database string
}

func NewHandler(email string) *NewEmailRepo {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cl, _ := mongo.Connect(ctx, options.Client().ApplyURI(email))
	mr := &NewEmailRepo{
		client:   cl,
		database: emailsDB,
	}
	return mr
}

func (mr *NewEmailRepo) Create(ctx context.Context, report *model.Email) (*model.Email, error) {

	collection := mr.client.Database(mr.database).Collection(emailsCollection)

	result, err := collection.InsertOne(ctx, report, &options.InsertOneOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "InsertOne")
	}

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.Wrap(ErrObjectIDTypeConversion, "email.InsertedID")
	}

	report.ID = objectID

	return report, nil
}

func (mr *NewEmailRepo) GetOne(e *model.Email, filter interface{}) error {
	collection := mr.client.Database(mr.database).Collection("email")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, filter).Decode(e)
	return err
}

func (mr *NewEmailRepo) GetAll(ctx context.Context) ([]model.Email, error) {
	collection := mr.client.Database(mr.database).Collection(emailsCollection)

	emails := make([]model.Email, 0)

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)
	for cursor.Next(ctx) {
		var email model.Email
		if err = cursor.Decode(&email); err != nil {
			return nil, err
		}
		emails = append(emails, email)
	}

	return emails, nil
}

func (mr *NewEmailRepo) AddOne(e *model.Email) (*mongo.InsertOneResult, error) {
	collection := mr.client.Database(mr.database).Collection("email")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, e)
	return result, err
}

func (mr *NewEmailRepo) RemoveOne(filter interface{}) (*mongo.DeleteResult, error) {
	collection := mr.client.Database(mr.database).Collection("email")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	result, err := collection.DeleteOne(ctx, filter)
	return result, err
}
