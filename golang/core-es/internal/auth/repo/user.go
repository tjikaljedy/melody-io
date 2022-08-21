package authrepo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAuth struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	userID    string             `bson:"userID" json:"userID"`
	username  string             `bson:"username" json:"username"`
	password  string             `bson:"password" json:"displayName"`
	email     string             `bson:"email" json:"email"`
	active    bool               `bson:"active" json:"active"`
	createdAt string             `bson:"createdAt" json:"createdAt"`
	updatedAt string             `bson:"updatedAt" json:"updatedAt"`
}

type UserProfile struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	profileID string             `bson:"profileID" json:"profileID"`
	email     string             `bson:"email" json:"email"`
	fullName  string             `bson:"fullName" json:"fullName"`
	dob       string             `bson:"dob" json:"dob"`
	active    bool               `bson:"active" json:"active"`
	createdAt string             `bson:"createdAt" json:"createdAt"`
	updatedAt string             `bson:"updatedAt" json:"updatedAt"`
}

type UserAuthRepo interface {
	InsertOne(ctx context.Context, userAuth *UserAuth) (interface{}, error)
	FindByUsername(ctx context.Context, username string) (*UserAuth, error)
}
