package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
    ID        primitive.ObjectID `json:"_id" bson:"_id"`
    CourseID  string             `json:"courseid" bson:"courseid"`
    CourseName string            `json:"coursename" bson:"coursename"`
    Cost      int               `json:"cost" bson:"cost"`
    Author    Author            `json:"author" bson:"author"`
}

type Author struct {
    Fullname string `json:"fullname" bson:"fullname"`
    Website  string `json:"website" bson:"website"`
}

func (c Course) IsEmpty() bool {
    return c.CourseName == ""
}

