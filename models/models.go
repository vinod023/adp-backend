package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	User struct {
		ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		UserID   string             `bson:"userid,omitempty" json:"userId"`
		Password string             `bson:"password,omitempty" json:"password"`
	}
)
