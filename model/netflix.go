package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Netflix struct {
    ID      primitive.ObjectID `json:"_id" bson:"_id"`
    Movie   string             `json:"movie" bson:"movie"`
    Watched bool              `json:"watched" bson:"watched"`
}

func (n Netflix) IsEmpty() bool {
    return n.Movie == ""
}

