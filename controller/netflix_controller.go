package controller

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
    "github.com/gorilla/mux"
    "github.com/hiteshchoudhary/mongodb/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

var NetflixCollection *mongo.Collection

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Get all movies")
    w.Header().Set("Content-Type", "application/json")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    cursor, err := NetflixCollection.Find(ctx, bson.M{})
    if err != nil {
        response := map[string]string{"error": "Failed to fetch movies"}
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
    }
    defer cursor.Close(ctx)

    var movies []model.Netflix
    for cursor.Next(ctx) {
        var movie model.Netflix
        if err := cursor.Decode(&movie); err != nil {
            response := map[string]string{"error": "Failed to decode movie"}
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(response)
            return
        }
        movies = append(movies, movie)
    }

    response := map[string]interface{}{
        "message": "Movies retrieved successfully",
        "data":    movies,
    }
    json.NewEncoder(w).Encode(response)
}

func GetOneMovie(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Get one movie")
    w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)
    id, err := primitive.ObjectIDFromHex(params["id"])
    if err != nil {
        response := map[string]string{"error": "Invalid movie ID"}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    var movie model.Netflix
    err = NetflixCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&movie)
    if err == mongo.ErrNoDocuments {
        response := map[string]string{"error": fmt.Sprintf("No Movie found with ID %s", params["id"])}
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(response)
        return
    } else if err != nil {
        response := map[string]string{"error": "Failed to fetch movie"}
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
    }

    response := map[string]interface{}{
        "message": "Movie retrieved successfully",
        "data":    movie,
    }
    json.NewEncoder(w).Encode(response)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Create a movie")
    w.Header().Set("Content-Type", "application/json")

    if r.Method != http.MethodPost {
        response := map[string]string{"error": "Only POST allowed!"}
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(response)
        return
    }

    var movie model.Netflix
    if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
        response := map[string]string{"error": "No data inside JSON or bad JSON toy!"}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }
    defer r.Body.Close()

    if movie.IsEmpty() {
        response := map[string]string{"error": "Need movie name!"}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    result, err := NetflixCollection.InsertOne(ctx, movie)
    if err != nil {
        response := map[string]string{"error": "Failed to create movie"}
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
    }

    movie.ID = result.InsertedID.(primitive.ObjectID)
    response := map[string]interface{}{
        "message": "Movie created successfully",
        "data":    movie,
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Update a movie")
    w.Header().Set("Content-Type", "application/json")

    if r.Method != http.MethodPut {
        response := map[string]string{"error": "Only PUT allowed!"}
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(response)
        return
    }

    params := mux.Vars(r)
    id, err := primitive.ObjectIDFromHex(params["id"])
    if err != nil {
        response := map[string]string{"error": "Invalid movie ID"}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }

    var updatedMovie model.Netflix
    if err := json.NewDecoder(r.Body).Decode(&updatedMovie); err != nil {
        response := map[string]string{"error": "No data inside JSON or bad JSON toy!"}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }
    defer r.Body.Close()

    if updatedMovie.IsEmpty() {
        response := map[string]string{"error": "Need movie name!"}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }

    updatedMovie.ID = id

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    result, err := NetflixCollection.ReplaceOne(ctx, bson.M{"_id": id}, updatedMovie)
    if err != nil {
        response := map[string]string{"error": "Failed to update movie"}
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
    }
    if result.MatchedCount == 0 {
        response := map[string]string{"error": fmt.Sprintf("No Movie found with ID %s", params["id"])}
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(response)
        return
    }

    response := map[string]interface{}{
        "message": "Movie updated successfully",
        "data":    updatedMovie,
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Delete a movie")
    w.Header().Set("Content-Type", "application/json")

    if r.Method != http.MethodDelete {
        response := map[string]string{"error": "Only DELETE allowed!"}
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(response)
        return
    }

    params := mux.Vars(r)
    id, err := primitive.ObjectIDFromHex(params["id"])
    if err != nil {
        response := map[string]string{"error": "Invalid movie ID"}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    result, err := NetflixCollection.DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        response := map[string]string{"error": "Failed to delete movie"}
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
    }
    if result.DeletedCount == 0 {
        response := map[string]string{"error": fmt.Sprintf("No Movie found with ID %s", params["id"])}
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(response)
        return
    }

    response := map[string]string{"message": fmt.Sprintf("Movie with ID %s deleted successfully", params["id"])}
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

func DeleteOneMovie(movieId string) {
    id, err := primitive.ObjectIDFromHex(movieId)
    if err != nil {
        log.Printf("Invalid movie ID: %v", err)
        return
    }
    filter := bson.M{"_id": id}
    result, err := NetflixCollection.DeleteOne(context.Background(), filter)
    if err != nil {
        log.Printf("Failed to delete movie: %v", err)
        return
    }
    fmt.Println("Movie deleted with count: ", result.DeletedCount)
}

func DeleteAllMovies() int64 {
    filter := bson.D{{}}
    result, err := NetflixCollection.DeleteMany(context.Background(), filter)
    if err != nil {
        log.Printf("Failed to delete all movies: %v", err)
        return 0
    }
    fmt.Println("Number of movies deleted: ", result.DeletedCount)
    return result.DeletedCount
}

func UpdateOneMovie(movieId string, update bson.M) {
    id, err := primitive.ObjectIDFromHex(movieId)
    if err != nil {
        log.Printf("Invalid movie ID: %v", err)
        return
    }
    filter := bson.M{"_id": id}
    result, err := NetflixCollection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        log.Printf("Failed to update movie: %v", err)
        return
    }
    fmt.Println("Modified count: ", result.ModifiedCount)
}

func getAllMovies() []primitive.M {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    cursor, err := NetflixCollection.Find(ctx, bson.M{})
    if err != nil {
        log.Printf("Failed to fetch movies: %v", err)
        return nil
    }
    defer cursor.Close(ctx)

    var movies []primitive.M
    if err := cursor.All(ctx, &movies); err != nil {
        log.Printf("Failed to decode movies: %v", err)
        return nil
    }
    return movies
}

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Get my all movies")
    w.Header().Set("Content-Type", "application/json")
    allMovies := getAllMovies()
    if allMovies == nil {
        response := map[string]string{"error": "Failed to fetch movies"}
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
    }
    response := map[string]interface{}{
        "message": "Movies retrieved successfully",
        "data":    allMovies,
    }
    json.NewEncoder(w).Encode(response)
}

func insertOneMovie(movie model.Netflix) interface{} {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    result, err := NetflixCollection.InsertOne(ctx, movie)
    if err != nil {
        log.Printf("Failed to insert movie: %v", err)
        return nil
    }
    return result.InsertedID
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Mark movie as watched")
    w.Header().Set("Content-Type", "application/json")

    if r.Method != http.MethodPut {
        response := map[string]string{"error": "Only PUT allowed!"}
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(response)
        return
    }

    params := mux.Vars(r)
    // Remove the unused id variable
    // id, err := primitive.ObjectIDFromHex(params["id"])
    // if err != nil {
    //     response := map[string]string{"error": "Invalid movie ID"}
    //     w.WriteHeader(http.StatusBadRequest)
    //     json.NewEncoder(w).Encode(response)
    //     return
    // }

    update := bson.M{"$set": bson.M{"watched": true}}
    UpdateOneMovie(params["id"], update)

    response := map[string]interface{}{
        "message": fmt.Sprintf("Movie with ID %s marked as watched", params["id"]),
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Delete a movie")
    w.Header().Set("Content-Type", "application/json")

    if r.Method != http.MethodDelete {
        response := map[string]string{"error": "Only DELETE allowed!"}
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(response)
        return
    }

    params := mux.Vars(r)
    id, err := primitive.ObjectIDFromHex(params["id"])
    if err != nil {
        response := map[string]string{"error": "Invalid movie ID"}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    result, err := NetflixCollection.DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        response := map[string]string{"error": "Failed to delete movie"}
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(response)
        return
    }
    if result.DeletedCount == 0 {
        response := map[string]string{"error": fmt.Sprintf("No Movie found with ID %s", params["id"])}
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(response)
        return
    }

    response := map[string]string{"message": fmt.Sprintf("Movie with ID %s deleted successfully", params["id"])}
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

func DeleteAllMoviesHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Delete all movies")
    w.Header().Set("Content-Type", "application/json")

    if r.Method != http.MethodDelete {
        response := map[string]string{"error": "Only DELETE allowed!"}
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(response)
        return
    }

    deletedCount := DeleteAllMovies()
    response := map[string]interface{}{
        "message": fmt.Sprintf("Deleted %d movies successfully", deletedCount),
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}