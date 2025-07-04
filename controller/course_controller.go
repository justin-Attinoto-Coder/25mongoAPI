package controller

import (
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
    "github.com/gorilla/mux"
    "github.com/hiteshchoudhary/mongodb/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    mathRand "math/rand"
)

var CourseCollection *mongo.Collection

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Get all courses")
    w.Header().Set("Content-Type", "application/json")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    cursor, err := CourseCollection.Find(ctx, bson.M{})
    if err != nil {
        response := map[string]string{"error": "Failed to fetch courses"}
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
    }
    defer cursor.Close(ctx)

    var courses []model.Course
    for cursor.Next(ctx) {
        var course model.Course
        if err := cursor.Decode(&course); err != nil {
            response := map[string]string{"error": "Failed to decode course"}
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(response)
            return
        }
        courses = append(courses, course)
    }

    response := map[string]interface{}{
        "message": "Courses retrieved successfully",
        "data":    courses,
    }
    json.NewEncoder(w).Encode(response)
}

func GetOneCourse(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Get one course")
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    id, ok := params["id"]
    if !ok {
        response := map[string]string{"error": "No Course ID provided"}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    var course model.Course
    err := CourseCollection.FindOne(ctx, bson.M{"courseid": id}).Decode(&course)
    if err == mongo.ErrNoDocuments {
        response := map[string]string{"error": fmt.Sprintf("No Course found with ID %s", id)}
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(response)
        return
    } else if err != nil {
        response := map[string]string{"error": "Failed to fetch course"}
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
    }

    response := map[string]interface{}{
        "message": "Course retrieved successfully",
        "data":    course,
    }
    json.NewEncoder(w).Encode(response)
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Create a course")
    w.Header().Set("Content-Type", "application/json")

    if r.Method != http.MethodPost {
        response := map[string]string{"error": "Only POST allowed!"}
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(response)
        return
    }

    var course model.Course
    if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
        response := map[string]string{"error": "No data inside JSON or bad JSON toy!"}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }
    defer r.Body.Close()

    if course.IsEmpty() {
        response := map[string]string{"error": "Need course name!"}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }

    course.CourseID = fmt.Sprintf("%v", mathRand.Intn(100))

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    result, err := CourseCollection.InsertOne(ctx, course)
    if err != nil {
        response := map[string]string{"error": "Failed to create course"}
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
    }

    course.ID = result.InsertedID.(primitive.ObjectID)
    response := map[string]interface{}{
        "message": "Course created successfully",
        "data":    course,
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Update a course")
    w.Header().Set("Content-Type", "application/json")

    if r.Method != http.MethodPut {
        response := map[string]string{"error": "Only PUT allowed!"}
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(response)
        return
    }

    params := mux.Vars(r)
    id, ok := params["id"]
    if !ok {
        response := map[string]string{"error": "No Course ID provided"}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }

    var updatedCourse model.Course
    if err := json.NewDecoder(r.Body).Decode(&updatedCourse); err != nil {
        response := map[string]string{"error": "No data inside JSON or bad JSON toy!"}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }
    defer r.Body.Close()

    if updatedCourse.IsEmpty() {
        response := map[string]string{"error": "Need course name!"}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }

    updatedCourse.CourseID = id

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    result, err := CourseCollection.ReplaceOne(ctx, bson.M{"courseid": id}, updatedCourse)
    if err != nil {
        response := map[string]string{"error": "Failed to update course"}
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
    }
    if result.MatchedCount == 0 {
        response := map[string]string{"error": fmt.Sprintf("No Course found with ID %s", id)}
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(response)
        return
    }

    response := map[string]interface{}{
        "message": "Course updated successfully",
        "data":    updatedCourse,
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Delete a course")
    w.Header().Set("Content-Type", "application/json")

    if r.Method != http.MethodDelete {
        response := map[string]string{"error": "Only DELETE allowed!"}
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(response)
        return
    }

    params := mux.Vars(r)
    id, ok := params["id"]
    if !ok {
        response := map[string]string{"error": "No Course ID provided"}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    result, err := CourseCollection.DeleteOne(ctx, bson.M{"courseid": id})
    if err != nil {
        response := map[string]string{"error": "Failed to delete course"}
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
    }
    if result.DeletedCount == 0 {
        response := map[string]string{"error": fmt.Sprintf("No Course found with ID %s", id)}
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(response)
        return
    }

    response := map[string]string{"message": fmt.Sprintf("Course with ID %s deleted successfully", id)}
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

