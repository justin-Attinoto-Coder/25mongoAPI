package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hiteshchoudhary/mongodb/controller"
	"github.com/hiteshchoudhary/mongodb/router"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Chapter struct {
	ID          string `bson:"_id" json:"id"`
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
}

var chapterCollection *mongo.Collection

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb+srv://mongotut:Mongo1234@cluster0.tknm0jh.mongodb.net/go_app?retryWrites=true&w=majority&appName=Cluster0")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}
	log.Println("Successfully connected to MongoDB")

	chapterCollection = client.Database("go_app").Collection("chapters")
	controller.CourseCollection = client.Database("go_app").Collection("courses")
	controller.NetflixCollection = client.Database("go_app").Collection("movies")

	router := router.Router()
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/courses", coursesHandler).Methods("GET")
	router.HandleFunc("/movies", moviesHandler).Methods("GET")
	router.HandleFunc("/chapter/{id}", chapterHandler).Methods("GET")
	router.HandleFunc("/chapter/19/server", chapter19Server).Methods("GET")
	router.HandleFunc("/chapter/21/server", chapter21Server).Methods("GET")
	router.HandleFunc("/chapter/23/server", chapter23Server).Methods("GET")
	router.HandleFunc("/chapter/24/server", chapter24Server).Methods("GET")
	router.HandleFunc("/api/courses", controller.GetAllCourses).Methods("GET")
	router.HandleFunc("/api/course/{id}", controller.GetOneCourse).Methods("GET")
	router.HandleFunc("/api/course", controller.CreateCourse).Methods("POST")
	router.HandleFunc("/api/course/{id}", controller.UpdateCourse).Methods("PUT")
	router.HandleFunc("/api/course/{id}", controller.DeleteCourse).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie/{id}", controller.GetOneMovie).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.UpdateMovie).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/my-movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie/{id}/watched", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMoviesHandler).Methods("DELETE")

	log.Println("Starting main server at http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := chapterCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch chapters: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var chapters []Chapter
	if err := cursor.All(ctx, &chapters); err != nil {
		http.Error(w, "Failed to decode chapters: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		http.Error(w, "Failed to load template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, chapters)
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/courses.html")
	if err != nil {
		http.Error(w, "Failed to load courses template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func moviesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/movies.html")
	if err != nil {
		http.Error(w, "Failed to load movies template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func chapterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var chapter Chapter
	err := chapterCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&chapter)
	if err != nil {
		http.Error(w, "Failed to fetch chapter: "+err.Error(), http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/chapter.html")
	if err != nil {
		http.Error(w, "Failed to load template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		ChapterID   string
		Title       string
		Description string
		Output      string
	}{ChapterID: id, Title: chapter.Title, Description: chapter.Description, Output: "Chapter " + id + " output"}
	tmpl.Execute(w, data)
}

func chapter19Server(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Fruit Shop!")
	})
	go func() {
		log.Println("Starting Chapter 19 server at http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", mux))
	}()
	fmt.Fprintln(w, "Chapter 19 server started at http://localhost:8080")
}

func chapter21Server(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Fruit Shop Web Server!")
	})
	go func() {
		log.Println("Starting Chapter 21 server at http://localhost:8000")
		log.Fatal(http.ListenAndServe(":8000", mux))
	}()
	fmt.Fprintln(w, "Chapter 21 server started at http://localhost:8000")
}

func chapter23Server(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Fruit Shop!")
	})
	go func() {
		log.Println("Starting Chapter 23 server at http://localhost:8001")
		log.Fatal(http.ListenAndServe(":8001", mux))
	}()
	fmt.Fprintln(w, "Chapter 23 server started at http://localhost:8001")
}

func chapter24Server(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the API Toy Shop!")
	})
	go func() {
		log.Println("Starting Chapter 24 server at http://localhost:8002")
		log.Fatal(http.ListenAndServe(":8002", mux))
	}()
	fmt.Fprintln(w, "Chapter 24 server started at http://localhost:8002")
}