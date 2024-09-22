package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string `bson:"email"`
	Name     string `bson:"name"`
	Password string `bson:"password"`
	Age      int8   `bson:"age"`
}
