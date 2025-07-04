package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"fmt"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var NetflixCollection *mongo.Collection

type Movie struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Movie   string             `json:"movie" bson:"movie"`
	Watched bool               `json:"watched" bson:"watched"`
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movies []Movie
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := NetflixCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Error fetching movies:", err)
		http.Error(w, "Failed to fetch movies: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var movie Movie
		if err := cursor.Decode(&movie); err != nil {
			log.Println("Error decoding movie:", err)
			http.Error(w, "Failed to decode movie: "+err.Error(), http.StatusInternalServerError)
			return
		}
		movies = append(movies, movie)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Movies retrieved successfully",
		"data":    movies,
	})
}

func GetOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "Invalid movie ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	var movie Movie
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = NetflixCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&movie)
	if err != nil {
		http.Error(w, "Movie not found: "+err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Movie retrieved successfully",
		"data":    movie,
	})
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if movie.Movie == "" {
		http.Error(w, "Missing required field: movie", http.StatusBadRequest)
		return
	}
	movie.ID = primitive.NewObjectID()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := NetflixCollection.InsertOne(ctx, movie)
	if err != nil {
		log.Println("Error creating movie:", err)
		http.Error(w, "Failed to create movie: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Movie created successfully",
		"data":    movie,
	})
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "Invalid movie ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if movie.Movie == "" {
		http.Error(w, "Missing required field: movie", http.StatusBadRequest)
		return
	}
	movie.ID = id
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := NetflixCollection.ReplaceOne(ctx, bson.M{"_id": id}, movie)
	if err != nil {
		log.Println("Error updating movie:", err)
		http.Error(w, "Failed to update movie: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if result.MatchedCount == 0 {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Movie updated successfully",
		"data":    movie,
	})
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "Invalid movie ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := NetflixCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Println("Error deleting movie:", err)
		http.Error(w, "Failed to delete movie: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if result.DeletedCount == 0 {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Movie deleted successfully",
	})
}

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movies []Movie
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := NetflixCollection.Find(ctx, bson.M{"watched": true})
	if err != nil {
		log.Println("Error fetching watched movies:", err)
		http.Error(w, "Failed to fetch watched movies: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var movie Movie
		if err := cursor.Decode(&movie); err != nil {
			log.Println("Error decoding movie:", err)
			http.Error(w, "Failed to decode movie: "+err.Error(), http.StatusInternalServerError)
			return
		}
		movies = append(movies, movie)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Watched movies retrieved successfully",
		"data":    movies,
	})
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "Invalid movie ID: "+err.Error(), http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := NetflixCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"watched": true}})
	if err != nil {
		log.Println("Error marking movie as watched:", err)
		http.Error(w, "Failed to mark movie as watched: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if result.MatchedCount == 0 {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	var movie Movie
	err = NetflixCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&movie)
	if err != nil {
		log.Println("Error fetching updated movie:", err)
		http.Error(w, "Movie not found: "+err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Movie marked as watched",
		"data":    movie,
	})
}

func DeleteAllMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := NetflixCollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		log.Println("Error deleting all movies:", err)
		http.Error(w, "Failed to delete all movies: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": fmt.Sprintf("Deleted %d movies successfully", result.DeletedCount),
	})
}