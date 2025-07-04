package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var CourseCollection *mongo.Collection

type Course struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	CourseId   string             `json:"courseid" bson:"courseid"`
	CourseName string             `json:"coursename" bson:"coursename"`
	Price      int                `json:"price" bson:"price"`
	Author     *Author            `json:"author" bson:"author"`
}

type Author struct {
	Fullname string `json:"fullname" bson:"fullname"`
	Website  string `json:"website" bson:"website"`
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var courses []Course
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := CourseCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch courses: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var course Course
		if err := cursor.Decode(&course); err != nil {
			http.Error(w, "Failed to decode course: "+err.Error(), http.StatusInternalServerError)
			return
		}
		courses = append(courses, course)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Courses retrieved successfully",
		"data":    courses,
	})
}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var course Course
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := CourseCollection.FindOne(ctx, bson.M{"courseid": id}).Decode(&course)
	if err != nil {
		http.Error(w, "Course not found: "+err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Course retrieved successfully",
		"data":    course,
	})
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if course.CourseId == "" || course.CourseName == "" || course.Author == nil || course.Author.Fullname == "" {
		http.Error(w, "Missing required fields: courseid, coursename, or author.fullname", http.StatusBadRequest)
		return
	}
	course.ID = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := CourseCollection.InsertOne(ctx, course)
	if err != nil {
		http.Error(w, "Failed to create course: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Course created successfully",
		"data":    course,
	})
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var course Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if course.CourseName == "" || course.Author == nil || course.Author.Fullname == "" {
		http.Error(w, "Missing required fields: coursename or author.fullname", http.StatusBadRequest)
		return
	}
	course.ID = primitive.NewObjectID()
	course.CourseId = id
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := CourseCollection.ReplaceOne(ctx, bson.M{"courseid": id}, course)
	if err != nil {
		http.Error(w, "Failed to update course: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Course updated successfully",
		"data":    course,
	})
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := CourseCollection.DeleteOne(ctx, bson.M{"courseid": id})
	if err != nil {
		http.Error(w, "Failed to delete course: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if result.DeletedCount == 0 {
		http.Error(w, "Course not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Course deleted successfully",
	})
}