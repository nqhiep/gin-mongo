package service

import (
	"context"

	"go-mongo/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	All(ctx context.Context) (*[]model.User, error)
	Load(ctx context.Context, id string) (*model.User, error)
	Insert(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User, id string) error
	Delete(ctx context.Context, id string) error
}

type userService struct {
	Collection *mongo.Collection
}

func NewUserService(db *mongo.Database) UserService {
	collectionName := "users"
	return &userService{Collection: db.Collection(collectionName)}
}

func (s *userService) All(ctx context.Context) (*[]model.User, error) {
	cursor, err := s.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	var users []model.User

	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (s *userService) Load(ctx context.Context, id string) (*model.User, error) {
	user := model.User{}
	err := s.Collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) Insert(ctx context.Context, user *model.User) (*model.User, error) {
	_, err := s.Collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, nil
	}
	return user, nil
}

func (s *userService) Update(ctx context.Context, user *model.User, id string) error {
	updateUser := bson.M{
		"$set": bson.M{
			"username":    user.Username,
			"email":       user.Email,
			"phone":       user.Phone,
			"dateOfBirth": user.DateOfBirth,
		},
	}

	_, err := s.Collection.UpdateOne(context.TODO(), bson.M{"id": id}, updateUser)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) Delete(ctx context.Context, id string) error {
	_, err := s.Collection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}
