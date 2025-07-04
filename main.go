package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"math"
	mathRand "math/rand"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
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

type user struct {
	name    string
	email   string
	status  bool
	age     int
	phone   string
	address string
}

type fruitBasket struct {
	fruits map[string]int
	name   string
}

type gamePlayer struct {
	name  string
	score int
	rolls int
}

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Website  string   `json:"website"`
	Password string   `json:"-,omitempty"`
	Tags     []string `json:"tags"`
}

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var chapterCollection *mongo.Collection
var courses []Course

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

	courses = append(courses,
		Course{
			CourseId:    "1",
			CourseName:  "Go API Bootcamp",
			CoursePrice: 299,
			Author:      &Author{Fullname: "Hitesh Choudhary", Website: "FruitLearnOnline.in"},
		},
		Course{
			CourseId:    "2",
			CourseName:  "ReactJS Bootcamp",
			CoursePrice: 199,
			Author:      &Author{Fullname: "Jane Doe", Website: "FruitLearnOnline.in"},
		},
	)

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
	router.HandleFunc("/api/movie/{id}", controller.GetOneCourse).Methods("GET")
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
		log.Println("Error fetching chapters:", err)
		http.Error(w, "Failed to fetch chapters: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var chapters []Chapter
	if err := cursor.All(ctx, &chapters); err != nil {
		log.Println("Error decoding chapters:", err)
		http.Error(w, "Failed to decode chapters: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		log.Println("Error loading index template:", err)
		http.Error(w, "Failed to load template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, chapters)
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/courses.html")
	if err != nil {
		log.Println("Error loading courses template:", err)
		http.Error(w, "Failed to load courses template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func moviesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/movies.html")
	if err != nil {
		log.Println("Error loading movies template:", err)
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
		log.Println("Error fetching chapter:", err)
		http.Error(w, "Failed to fetch chapter: "+err.Error(), http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/chapter.html")
	if err != nil {
		log.Println("Error loading chapter template:", err)
		http.Error(w, "Failed to load template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	output := ""
	switch id {
	case "01":
		output = chapter01()
	case "02":
		output = chapter02()
	case "03":
		output = chapter03()
	case "04":
		output = chapter04()
	case "05":
		output = chapter05()
	case "06":
		output = chapter06()
	case "07":
		output = chapter07()
	case "08":
		output = chapter08()
	case "09":
		output = chapter09()
	case "10":
		output = chapter10()
	case "11":
		output = chapter11()
	case "12":
		output = chapter12()
	case "13":
		output = chapter13()
	case "14":
		output = chapter14()
	case "15":
		output = chapter15()
	case "16":
		output = chapter16()
	case "17":
		output = chapter17()
	case "18":
		output = chapter18()
	case "19":
		output = chapter19()
	case "20":
		output = chapter20()
	case "21":
		output = chapter21()
	case "22":
		output = chapter22()
	case "23":
		output = chapter23()
	case "24":
		output = chapter24()
	case "26":
		output = chapter26()
	case "27":
		output = chapter27()
	case "28":
		output = chapter28()
	default:
		output = "Chapter not implemented yet!"
	}

	data := struct {
		ChapterID   string
		Title       string
		Description string
		Output      string
	}{ChapterID: id, Title: chapter.Title, Description: chapter.Description, Output: output}
	tmpl.Execute(w, data)
}

func chapter01() string {
	output := "Welcome to Go Programming!\n"
	name := "Justin"
	age := 41
	output += "Hello, " + name + "! You are " + strconv.Itoa(age) + " years old.\n"
	yearsLater := age + 5
	output += "In 5 years, " + name + " will be " + strconv.Itoa(yearsLater) + " years old.\n"
	return output
}

func chapter02() string {
	output := "Exploring Variables, Types, and Constants\n"
	var username string = "Justin"
	var age int = 41
	city := "New York"
	var (
		height    float32 = 5.9
		isStudent bool    = false
	)
	const birthYear int = 1984
	const loginToken string = "gibberish"
	numberOfUser := 30000.0
	output += username + "\n"
	output += "variable is of type: " + fmt.Sprintf("%T", username) + "\n"
	output += "Age: " + strconv.Itoa(age) + " (Type: " + fmt.Sprintf("%T", age) + ")\n"
	output += "City: " + city + " (Type: " + fmt.Sprintf("%T", city) + ")\n"
	output += "Height: " + fmt.Sprintf("%.1f", height) + " (Type: " + fmt.Sprintf("%T", height) + ")\n"
	output += "Is Student: " + fmt.Sprint(isStudent) + " (Type: " + fmt.Sprintf("%T", isStudent) + ")\n"
	output += "Birth Year: " + strconv.Itoa(birthYear) + " (Type: " + fmt.Sprintf("%T", birthYear) + ")\n"
	output += "Login Token: " + loginToken + " (Type: " + fmt.Sprintf("%T", loginToken) + ")\n"
	output += fmt.Sprint(numberOfUser) + "\n"
	output += "Number of Users: " + fmt.Sprintf("%.1f", numberOfUser) + " (Type: " + fmt.Sprintf("%T", numberOfUser) + ")\n"
	yearsSinceBirth := 2025 - birthYear
	output += username + ", you are " + strconv.Itoa(yearsSinceBirth) + " years old in 2025.\n"
	return output
}

func chapter03() string {
	output := "Welcome to user input\n"
	output += "Enter your name: [Input not available in web mode]\n"
	name := "Justin"
	output += "Enter your age: [Input not available in web mode]\n"
	ageInput := "41"
	inputs := []interface{}{name, ageInput}
	for i, input := range inputs {
		if value, ok := input.(string); ok {
			output += "Input " + strconv.Itoa(i+1) + " is a string: " + value + "\n"
		} else {
			output += "Input " + strconv.Itoa(i+1) + " is not a string\n"
		}
	}
	output += "Enter a short bio: [Input not available in web mode]\n"
	bio := "I love coding in Go!"
	output += "Your bio: " + bio + "\n"
	runeCount := len([]rune(bio))
	output += "Your bio has " + strconv.Itoa(runeCount) + " runes\n"
	output += "Enter a favorite word: [Input not available in web mode]\n"
	word := "Go"
	output += "Your favorite word: " + word + "\n"
	return output
}

func chapter04() string {
	output := "Welcome to our pizza app\n"
	output += "Please rate our pizza between 1 and 5\n"
	output += "Enter rating: [Input not available in web mode]\n"
	input := "4.5"
	output += "Thanks for rating, " + input + "\n"
	numRating, err := strconv.ParseFloat(input, 64)
	if err != nil {
		output += "Error converting rating to float: " + fmt.Sprint(err) + "\n"
		return output
	}
	output += "Added 1 to your rating: " + fmt.Sprintf("%.1f", numRating+1) + "\n"
	ratingStr := strconv.FormatFloat(numRating, 'f', 1, 64)
	output += "Your rating as string: " + ratingStr + "\n"
	output += "Enter number of pizzas ordered: [Input not available in web mode]\n"
	pizzaInput := "3"
	numPizzas, err := strconv.Atoi(pizzaInput)
	if err != nil {
		output += "Error converting to integer: " + fmt.Sprint(err) + "\n"
		return output
	}
	output += "You ordered " + strconv.Itoa(numPizzas) + " pizzas\n"
	pizzaCountStr := strconv.Itoa(numPizzas)
	output += "Number of pizzas as string: " + pizzaCountStr + "\n"
	output += "Are you a premium member? (true/false): [Input not available in web mode]\n"
	premiumInput := "true"
	isPremium, err := strconv.ParseBool(premiumInput)
	if err != nil {
		output += "Error converting to boolean: " + fmt.Sprint(err) + "\n"
		return output
	}
	output += "Premium member status: " + fmt.Sprint(isPremium) + "\n"
	premiumStr := strconv.FormatBool(isPremium)
	output += "Premium status as string: " + premiumStr + "\n"
	var data interface{} = numRating
	if val, ok := data.(float64); ok {
		output += "Rating is a float64: " + fmt.Sprintf("%.1f", val) + "\n"
	} else {
		output += "Rating is not a float64\n"
	}
	return output
}

func chapter05() string {
	output := "Welcome to our math playground!\n"
	number1 := 10.5
	number2 := 3
	sum := number1 + float64(number2)
	output += "Adding " + fmt.Sprint(number1) + " and " + fmt.Sprint(number2) + " gives us " + fmt.Sprint(sum) + "!\n"
	difference := number1 - float64(number2)
	output += "Taking " + fmt.Sprint(number2) + " away from " + fmt.Sprint(number1) + " leaves " + fmt.Sprint(difference) + "!\n"
	product := number1 * float64(number2)
	output += fmt.Sprint(number2) + " groups of " + fmt.Sprint(number1) + " makes " + fmt.Sprint(product) + "!\n"
	quotient := number1 / float64(number2)
	output += "Sharing " + fmt.Sprint(number1) + " among " + fmt.Sprint(number2) + " gives " + fmt.Sprint(quotient) + " each!\n"
	rounded := math.Round(number1)
	output += "Rounding " + fmt.Sprint(number1) + " gives us " + fmt.Sprint(rounded) + "!\n"
	ceiling := math.Ceil(number1)
	output += "Ceiling of " + fmt.Sprint(number1) + " is " + fmt.Sprint(ceiling) + "!\n"
	floor := math.Floor(number1)
	output += "Floor of " + fmt.Sprint(number1) + " is " + fmt.Sprint(floor) + "!\n"
	maxNumber := math.Max(number1, float64(number2))
	output += "The bigger number between " + fmt.Sprint(number1) + " and " + fmt.Sprint(number2) + " is " + fmt.Sprint(maxNumber) + "!\n"
	power := math.Pow(2, 3)
	output += "2 grown 3 times is " + fmt.Sprint(power) + "!\n"
	sqrt := math.Sqrt(16)
	output += "The square root of 16 is " + fmt.Sprint(sqrt) + "!\n"
	mathRand.Seed(time.Now().UnixNano())
	output += "Random number: " + fmt.Sprint(mathRand.Intn(5)) + "\n"
	output += "Crypto random number: " + fmt.Sprint(mathRand.Intn(5)) + "\n"
	return output
}

func chapter06() string {
	output := "Welcome to time study of golang\n"
	presentTime := time.Now()
	output += "Right now, the time is: " + presentTime.String() + "\n"
	formattedTime := presentTime.Format("01-02-2006 15:04:05 Monday")
	output += "Pretty time is: " + formattedTime + "\n"
	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	output += "Special date is: " + createdDate.String() + "\n"
	formattedDate := createdDate.Format("01-02-2006 Monday")
	output += "Pretty special date is: " + formattedDate + "\n"
	futureTime := presentTime.Add(2 * time.Hour)
	output += "Two hours from now: " + futureTime.String() + "\n"
	timeDifference := futureTime.Sub(presentTime)
	output += "Time difference is: " + fmt.Sprint(timeDifference.Hours()) + " hours\n"
	isBefore := createdDate.Before(presentTime)
	output += "Is special date before now? " + fmt.Sprint(isBefore) + "\n"
	shortFormat := presentTime.Format("2006-01-02")
	output += "Short date format: " + shortFormat + "\n"
	timeOnly := presentTime.Format("15:04")
	output += "Time only format: " + timeOnly + "\n"
	return output
}

func chapter07() string {
	output := "Welcome to building Go programs for different computers!\n"
	hostOS := runtime.GOOS
	output += "Our computer's operating system (GOHOSTOS): " + hostOS + "\n"
	output += "Default target operating system (GOOS): " + runtime.GOOS + "\n"
	user := "Justin"
	output += "Hello from " + user + " on " + hostOS + " !\n"
	switch hostOS {
	case "windows":
		output += "Yay, we're on Windows! This program loves .exe files!\n"
	case "linux":
		output += "Cool, we're on Linux! This program runs without extensions!\n"
	case "darwin":
		output += "Awesome, we're on a Mac! This program loves macOS!\n"
	}
	goArch := os.Getenv("GOARCH")
	output += "Our computer's brain type (GOARCH): " + goArch + "\n"
	goPath := os.Getenv("GOPATH")
	output += "Where Go keeps its toys (GOPATH): " + goPath + "\n"
	return output
}

func chapter08() string {
	output := "Welcome to our fruit basket adventure with arrays!\n"
	var fruitBasket [4]string
	fruitBasket[0] = "Apple"
	fruitBasket[1] = "Banana"
	fruitBasket[2] = "Orange"
	fruitBasket[3] = "Mango"
	output += "My fruit basket: " + fmt.Sprint(fruitBasket) + "\n"
	output += "First fruit in slot 0: " + fruitBasket[0] + "\n"
	fruitBasket[2] = "Grape"
	output += "New fruit basket: " + fmt.Sprint(fruitBasket) + "\n"
	basketSize := len(fruitBasket)
	output += "My basket has " + strconv.Itoa(basketSize) + " slots\n"
	veggieBasket := [3]string{"Carrot", "Potato", "Cucumber"}
	output += "My veggie basket: " + fmt.Sprint(veggieBasket) + "\n"
	output += "Checking each fruit in the fruit basket:\n"
	for i := 0; i < len(fruitBasket); i++ {
		output += "Slot " + strconv.Itoa(i) + " has " + fruitBasket[i] + "\n"
	}
	output += "Using magic hand to check veggies:\n"
	for index, veggie := range veggieBasket {
		output += "Slot " + strconv.Itoa(index) + " has " + veggie + "\n"
	}
	fruitCount := [5]int{10, 20, 15, 30, 25}
	output += "Fruit counts: " + fmt.Sprint(fruitCount) + "\n"
	totalFruits := 0
	for _, count := range fruitCount {
		totalFruits += count
	}
	output += "Total fruits in the big basket: " + strconv.Itoa(totalFruits) + "\n"
	doubleBasket := [2][3]string{
		{"Apple", "Banana", "Grape"},
		{"Mango", "Orange", "Pineapple"},
	}
	output += "My double basket: " + fmt.Sprint(doubleBasket) + "\n"
	output += "Fruit in row 1, column 2: " + doubleBasket[0][1] + "\n"
	output += "Checking all fruits in the double basket:\n"
	for row := 0; row < len(doubleBasket); row++ {
		for col := 0; col < len(doubleBasket[row]); col++ {
			output += "Row " + strconv.Itoa(row) + " Column " + strconv.Itoa(col) + " has " + doubleBasket[row][col] + "\n"
		}
	}
	return output
}

func chapter09() string {
	mathRand.Seed(time.Now().UnixNano())
	output := "Welcome to our slice adventure with fruits, scores, and courses!\n"
	fruitBasket := []string{"Apple", "Tomato", "Peach"}
	output += "My stretchy fruit basket: " + fmt.Sprint(fruitBasket) + "\n"
	fruitBasket = append(fruitBasket, "Mango", "Banana")
	output += "After adding Mango and Banana: " + fmt.Sprint(fruitBasket) + "\n"
	fruitBasket = append(fruitBasket[:1], fruitBasket[2:]...)
	output += "Fruit basket after removing Tomato: " + fmt.Sprint(fruitBasket) + "\n"
	highScores := []int{
		mathRand.Intn(1000) + 1,
		mathRand.Intn(1000) + 1,
		mathRand.Intn(1000) + 1,
		mathRand.Intn(1000) + 1,
	}
	output += "My game high scores: " + fmt.Sprint(highScores) + "\n"
	newScore := mathRand.Intn(1000) + 1
	highScores = append(highScores, newScore)
	output += "After adding new score " + strconv.Itoa(newScore) + ": " + fmt.Sprint(highScores) + "\n"
	highScores = append(highScores[:1], highScores[2:]...)
	output += "High scores after removing second score: " + fmt.Sprint(highScores) + "\n"
	courses := []string{"ReactJS", "JavaScript", "Swift", "Python", "Ruby"}
	output += "My courses box: " + fmt.Sprint(courses) + "\n"
	newCourseIndex := len(courses)
	newCourse := "Course" + strconv.Itoa(newCourseIndex+1)
	courses = append(courses, newCourse)
	output += "After adding " + newCourse + ": " + fmt.Sprint(courses) + "\n"
	courses = append(courses[:2], courses[3:]...)
	output += "Courses after removing Swift: " + fmt.Sprint(courses) + "\n"
	coursesSize := len(courses)
	coursesCapacity := cap(courses)
	output += "Courses box has " + strconv.Itoa(coursesSize) + " courses and room for " + strconv.Itoa(coursesCapacity) + "\n"
	output += "Course in slot 0: " + courses[0] + "\n"
	courses[1] = "NodeJS"
	output += "New courses box: " + fmt.Sprint(courses) + "\n"
	someCourses := courses[1:4]
	output += "Some courses: " + fmt.Sprint(someCourses) + "\n"
	output += "Checking each course:\n"
	for i := 0; i < len(courses); i++ {
		output += "Slot " + strconv.Itoa(i) + " has course " + courses[i] + "\n"
	}
	output += "Using magic hand to check courses:\n"
	for index, course := range courses {
		output += "Slot " + strconv.Itoa(index) + " has course " + course + "\n"
	}
	someCourses[0] = "Angular"
	output += "Changed some courses: " + fmt.Sprint(someCourses) + "\n"
	output += "Original courses after change: " + fmt.Sprint(courses) + "\n"
	coursesCopy := make([]string, len(courses))
	copy(coursesCopy, courses)
	output += "Copied courses: " + fmt.Sprint(coursesCopy) + "\n"
	coursesCopy[0] = "VueJS"
	output += "Changed copy: " + fmt.Sprint(coursesCopy) + "\n"
	output += "Original courses unchanged: " + fmt.Sprint(courses) + "\n"
	return output
}

func chapter10() string {
	output := "Welcome to our magic toy box adventure with maps!\n"
	fruitInventory := make(map[string]int)
	fruitInventory["Apple"] = 10
	fruitInventory["Banana"] = 15
	fruitInventory["Mango"] = 8
	output += "My fruit inventory: " + fmt.Sprint(fruitInventory) + "\n"
	output += "Checking all fruits with a magic hand:\n"
	for fruit, count := range fruitInventory {
		output += "Tag " + fruit + " has " + strconv.Itoa(count) + " fruits\n"
	}
	bananaCount := fruitInventory["Banana"]
	output += "Number of bananas: " + strconv.Itoa(bananaCount) + "\n"
	count, exists := fruitInventory["Mango"]
	if exists {
		output += "Found Mango with " + strconv.Itoa(count) + " items!\n"
	} else {
		output += "No Mango in the box!\n"
	}
	delete(fruitInventory, "Mango")
	output += "After removing Mango: " + fmt.Sprint(fruitInventory) + "\n"
	studentScores := map[string]float64{
		"Justin": 85.5,
		"Alice":  92.0,
		"Bob":    78.5,
	}
	output += "Student scores: " + fmt.Sprint(studentScores) + "\n"
	output += "Checking all student scores with a magic hand:\n"
	for student, score := range studentScores {
		output += "Student " + student + " has score " + fmt.Sprint(score) + "\n"
	}
	studentScores["Emma"] = 88.0
	output += "After adding Emma: " + fmt.Sprint(studentScores) + "\n"
	courseSchedule := make(map[string]string)
	courseSchedule["ReactJS"] = "Monday 9AM"
	courseSchedule["Python"] = "Tuesday 2PM"
	courseSchedule["Swift"] = "Wednesday 11AM"
	output += "Course schedule: " + fmt.Sprint(courseSchedule) + "\n"
	output += "Checking all courses with a magic hand:\n"
	for course, time := range courseSchedule {
		output += "Course " + course + " is at " + time + "\n"
	}
	courseSchedule["Swift"] = "Friday 10AM"
	output += "Updated course schedule: " + fmt.Sprint(courseSchedule) + "\n"
	languages := map[string]string{
		"Go":         "Compiled",
		"Python":     "Interpreted",
		"JavaScript": "Interpreted",
		"Ruby":       "Interpreted",
		"Swift":      "Compiled",
	}
	output += "Programming languages: " + fmt.Sprint(languages) + "\n"
	output += "Checking all languages with a magic hand:\n"
	for lang, langType := range languages {
		output += "Language " + lang + " is " + langType + "\n"
	}
	languages["Java"] = "Compiled"
	output += "After adding Java: " + fmt.Sprint(languages) + "\n"
	langType, exists := languages["Python"]
	if exists {
		output += "Found Python, it's " + langType + " !\n"
	} else {
		output += "No Python in the box!\n"
	}
	delete(languages, "Ruby")
	output += "After removing Ruby: " + fmt.Sprint(languages) + "\n"
	langCount := len(languages)
	output += "Number of languages in the box: " + strconv.Itoa(langCount) + "\n"
	output += "All our magic boxes are ready to play!\n"
	return output
}

func chapter11() string {
	output := "Welcome to our custom toy box adventure with structs!\n"
	justin := user{
		name:    "Justin",
		email:   "justin@example.com",
		status:  false,
		age:     41,
		phone:   "555-1234",
		address: "123 Fruit Street",
	}
	output += "Justin's toy box: " + fmt.Sprintf("%+v", justin) + "\n"
	output += captureOutput(func() { justin.showDetails() })
	justin.verifyEmail()
	output += "Justin's status after verification: " + fmt.Sprint(justin.status) + "\n"
	justin.updatePhone("555-9999")
	output += "Justin's new phone: " + justin.phone + "\n"
	yearsToRetire := justin.yearsUntilRetirement(65)
	output += "Years until Justin retires: " + strconv.Itoa(yearsToRetire) + "\n"
	alice := createUser("Alice", "alice@example.com", 25, "555-5678", "456 Banana Avenue")
	output += "Alice's toy box: " + fmt.Sprintf("%+v", alice) + "\n"
	output += captureOutput(func() { alice.showDetails() })
	alice.verifyEmail()
	output += "Alice's status after verification: " + fmt.Sprint(alice.status) + "\n"
	users := []user{justin, alice}
	output += "All user toy boxes: " + fmt.Sprint(users) + "\n"
	output += "Checking all users with helpers:\n"
	for i, u := range users {
		output += "User " + strconv.Itoa(i) + ":\n"
		output += captureOutput(func() { u.showDetails() })
		output += "Years to retire: " + strconv.Itoa(u.yearsUntilRetirement(65)) + "\n"
	}
	return output
}

func chapter12() string {
	output := "Welcome to our decision-making adventure with if and else!\n"
	justin := user{
		name:   "Justin",
		email:  "justin@example.com",
		status: false,
		age:    41,
	}
	if justin.age >= 18 {
		output += "Justin can vote! He's " + strconv.Itoa(justin.age) + " years old.\n"
	} else {
		output += "Justin is too young to vote. He's only " + strconv.Itoa(justin.age) + "\n"
	}
	if justin.status {
		output += "Justin's email is verified! Welcome!\n"
	} else {
		output += "Justin, please verify your email!\n"
		output += captureOutput(func() { justin.verifyEmail() })
		output += "After verification, status is: " + fmt.Sprint(justin.status) + "\n"
	}
	if justin.age < 13 {
		output += "Justin is a kid!\n"
	} else if justin.age < 20 {
		output += "Justin is a teen!\n"
	} else if justin.age < 30 {
		output += "Justin is a young adult!\n"
	} else {
		output += "Justin is an adult!\n"
	}
	alice := user{
		name:   "Alice",
		email:  "alice@example.com",
		status: true,
		age:    25,
	}
	if strings.Contains(alice.email, "fruit") {
		output += "Alice has a fruity email!\n"
	} else {
		output += "Alice's email isn't fruity!\n"
	}
	loginAttempts := 6
	maxAttempts := 5
	if loginAttempts <= maxAttempts {
		output += "Login okay! You tried " + strconv.Itoa(loginAttempts) + " times.\n"
	} else {
		output += "Too many tries! " + strconv.Itoa(loginAttempts) + " is more than " + strconv.Itoa(maxAttempts) + "\n"
		output += "Account locked! Call Fruit Support at 555-FRUIT!\n"
	}
	if alice.status && alice.age < 30 {
		output += "Alice is verified and young! Special fruit discount!\n"
	} else {
		output += "Alice doesn't get the young verified discount.\n"
	}
	users := []user{justin, alice}
	output += "Checking all users:\n"
	for i, u := range users {
		output += "User " + strconv.Itoa(i) + ":\n"
		if u.status {
			output += u.name + " is verified!\n"
		} else {
			output += u.name + " needs to verify their email!\n"
		}
	}
	return output
}

func chapter13() string {
	mathRand.Seed(time.Now().UnixNano())
	output := "Welcome to our fruit board dice game adventure!\n"
	board := []string{"Start", "Apple", "Banana", "Orange", "Mango", "Peach", "Grape", "Finish"}
	playerPosition := 0
	maxPosition := len(board) - 1
	output += "Let's roll the dice and move on the fruit board!\n"
	for playerPosition < maxPosition {
		diceNumber := mathRand.Intn(6) + 1
		output += "Value of dice is " + strconv.Itoa(diceNumber) + "\n"
		switch diceNumber {
		case 1:
			playerPosition += 1
			output += "You move 1 spot!\n"
		case 2:
			playerPosition += 2
			output += "You move 2 spots!\n"
		case 3:
			playerPosition += 3
			output += "You move 3 spots!\n"
		case 4:
			playerPosition += 4
			output += "You move 4 spots!\n"
		case 5:
			playerPosition += 5
			output += "You move 5 spots!\n"
		case 6:
			playerPosition += 6
			output += "Wow, a 6! You move 6 spots and roll again!\n"
			continue
		}
		if playerPosition > maxPosition {
			playerPosition = maxPosition
		}
		output += "You're now at " + board[playerPosition] + "\n"
		output += "---\n"
	}
	if playerPosition == maxPosition {
		output += "Yay! You reached the Finish! You win the fruit game!\n"
	}
	output += "\nLet's roll one more dice for fun!\n"
	anotherDice := mathRand.Intn(6) + 1
	output += "Dice value: " + strconv.Itoa(anotherDice) + "\n"
	switch {
	case anotherDice <= 3:
		output += "Low roll! You get a small fruit prize!\n"
	case anotherDice >= 4:
		output += "High roll! You get a big fruit basket!\n"
	}
	return output
}

func chapter14() string {
	mathRand.Seed(time.Now().UnixNano())
	output := "Welcome to our fruit board dice game adventure!\n"
	board := []string{"Start", "Apple", "Banana", "Orange", "Mango", "Peach", "Grape", "Finish"}
	playerPosition := 0
	maxPosition := len(board) - 1
	output += "Let's roll the dice and move on the fruit board!\n"
	for playerPosition < maxPosition {
		diceNumber := mathRand.Intn(6) + 1
		output += "Value of dice is " + strconv.Itoa(diceNumber) + "\n"
		switch diceNumber {
		case 1:
			playerPosition += 1
			output += "You move 1 spot!\n"
		case 2:
			playerPosition += 2
			output += "You move 2 spots!\n"
		case 3:
			playerPosition += 3
			output += "You move 3 spots!\n"
		case 4:
			playerPosition += 4
			output += "You move 4 spots!\n"
		case 5:
			playerPosition += 5
			output += "You move 5 spots!\n"
		case 6:
			playerPosition += 6
			output += "Wow, a 6! You move 6 spots and roll again!\n"
			continue
		}
		if playerPosition > maxPosition {
			playerPosition = maxPosition
		}
		output += "You're now at " + board[playerPosition] + "\n"
		output += "---\n"
	}
	if playerPosition == maxPosition {
		output += "Yay! You reached the Finish! You win the fruit game!\n"
	}
	output += "\nLet's roll one more dice for fun!\n"
	anotherDice := mathRand.Intn(6) + 1
	output += "Dice value: " + strconv.Itoa(anotherDice) + "\n"
	switch {
	case anotherDice <= 3:
		output += "Low roll! You get a small fruit prize!\n"
	case anotherDice >= 4:
		output += "High roll! You get a big fruit basket!\n"
	}
	return output
}

func chapter15() string {
	output := "Welcome to our fruit shop function adventure!\n"
	output += captureOutput(sayHello)
	output += captureOutput(func() { greetUser("Justin") })
	applePrice := getFruitPrice("Apple", 2)
	output += "Price for 2 apples: " + strconv.Itoa(applePrice) + " coins\n"
	haveBananas := checkStock("Banana", 5)
	output += "Enough bananas for 5? " + fmt.Sprint(haveBananas) + "\n"
	total, discount := calculateOrder("Mango", 10)
	output += "Total for 10 mangos: " + strconv.Itoa(total) + " coins, Discount: " + strconv.Itoa(discount) + " coins\n"
	totalCost := buyFruits("Apple", "Banana", "Mango")
	output += "Total cost for fruits: " + strconv.Itoa(totalCost) + " coins\n"
	countFruits := func() int { return 3 }
	output += "Number of fruit types: " + strconv.Itoa(countFruits()) + "\n"
	return output
}

func chapter16() string {
	mathRand.Seed(time.Now().UnixNano())
	output := "Welcome to our toy box helper adventure with methods!\n"
	output += "\nExample 1: User Profile\n"
	justin := user{name: "Justin", email: "justin@example.com", status: false, age: 41}
	output += captureOutput(func() { justin.showProfile() })
	output += captureOutput(func() { justin.verifyEmail() })
	output += captureOutput(func() { justin.updateAge(42) })
	output += "After updates:\n"
	output += captureOutput(func() { justin.showProfile() })
	output += "\nExample 2: Fruit Basket\n"
	basket := fruitBasket{fruits: map[string]int{"Apple": 5, "Banana": 3}, name: "Justin's Basket"}
	output += captureOutput(func() { basket.showContents() })
	output += captureOutput(func() { basket.addFruit("Mango", 4) })
	output += captureOutput(func() { basket.removeFruit("Banana") })
	output += "After updates:\n"
	output += captureOutput(func() { basket.showContents() })
	output += "\nExample 3: Dice Game Player\n"
	alice := gamePlayer{name: "Alice", score: 0, rolls: 0}
	output += captureOutput(func() { alice.showStatus() })
	output += captureOutput(func() { alice.rollDice() })
	output += captureOutput(func() { alice.doubleScore() })
	output += "After playing:\n"
	output += captureOutput(func() { alice.showStatus() })
	return output
}

func chapter17() string {
	output := "Welcome to our fruit shop cleanup adventure with defer!\n"
	output += "\nExample 1: Cleaning up the fruit shop\n"
	output += captureOutput(sellFruits)
	output += "\nExample 2: User session with logging\n"
	justin := user{name: "Justin"}
	output += captureOutput(func() { manageSession(justin) })
	output += "\nExample 3: Writing to a fruit order file\n"
	output += captureOutput(writeFruitOrder)
	return output
}

func chapter18() string {
	output := "Welcome to files in golang!\n"
	content := "Fruit order: 5 Apples, 3 Bananas - FruitShop.in"
	file, err := os.CreateTemp("", "myfruitfile*.txt")
	if err != nil {
		output += "Error creating file: " + fmt.Sprint(err) + "\n"
		return output
	}
	defer os.Remove(file.Name())
	length, err := io.WriteString(file, content)
	if err != nil {
		output += "Error writing to file: " + fmt.Sprint(err) + "\n"
		return output
	}
	output += "Wrote " + strconv.Itoa(length) + " letters to the file!\n"
	file.Close()
	output += captureOutput(func() { readFile(file.Name()) })
	output += captureOutput(func() { appendToFile(file.Name(), "\nExtra: 2 Mangos") })
	output += "Added more fruits to the file!\n"
	output += captureOutput(func() { readFile(file.Name()) })
	return output
}

func chapter19() string {
	output := "Welcome to our fruit shop web request adventure!\n"
	const url = "https://lco.dev"
	output += "\nFetching from " + url + "\n"
	output += captureOutput(func() { fetchWebsite(url) })
	const fruitUrl = "https://example.com"
	output += "\nFetching from " + fruitUrl + "\n"
	output += captureOutput(func() { fetchWebsite(fruitUrl) })
	output += "\nVisit /chapter/19/server to start the fruit shop web server on http://localhost:8080\n"
	return output
}

func chapter19Server(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Fruit Shop!")
	})
	mux.HandleFunc("/fruits", func(w http.ResponseWriter, r *http.Request) {
		fruits := []string{"Apple", "Banana", "Mango"}
		json.NewEncoder(w).Encode(fruits)
	})
	go func() {
		log.Println("Starting Chapter 19 server at http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", mux))
	}()
	fmt.Fprintln(w, "Chapter 19 server started at http://localhost:8080")
}

func chapter20() string {
	output := "Welcome to handling URLs in golang!\n"
	const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"
	output += "\nParsing URL: " + myurl + "\n"
	output += captureOutput(func() { parseURL(myurl) })
	const fruitUrl = "https://fruitshop.com:8080/order?fruit=apple&quantity=5&customer=Justin"
	output += "\nParsing fruit shop URL: " + fruitUrl + "\n"
	output += captureOutput(func() { parseURL(fruitUrl) })
	output += "\nBuilding a new fruit shop URL\n"
	newUrl := buildFruitShopURL("https", "fruitshop.com", "/cart", map[string]string{
		"fruit":     "mango",
		"quantity":  "3",
		"promo":     "FRUIT10",
	})
	output += "New URL: " + newUrl + "\n"
	return output
}

func chapter21() string {
	output := "Welcome to our fruit shop web server adventure!\n"
	output += "\nMaking GET and POST requests\n"
	output += captureOutput(func() { PerformGetRequest("https://example.com") })
	output += "\nVisit /chapter/21/server to start the fruit shop web server on http://localhost:8000\n"
	return output
}

func chapter21Server(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Fruit Shop Web Server!")
	})
	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is a GET request response!")
	})
	mux.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var data map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				http.Error(w, "Error decoding JSON", http.StatusBadRequest)
				return
			}
			json.NewEncoder(w).Encode(data)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			r.ParseForm()
			data := make(map[string]string)
			for key, values := range r.Form {
				data[key] = values[0]
			}
			json.NewEncoder(w).Encode(data)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	go func() {
		log.Println("Starting Chapter 21 server at http://localhost:8000")
		log.Fatal(http.ListenAndServe(":8000", mux))
	}()
	fmt.Fprintln(w, "Chapter 21 server started at http://localhost:8000")
}

func chapter22() string {
	output := "Welcome to our JSON toy box adventure!\n"
	output += captureOutput(EncodeJson)
	output += captureOutput(DecodeJson)
	output += captureOutput(ConsumeWebJson)
	return output
}

func chapter23() string {
	output := "Welcome to our Go module adventure!\n"
	output += captureOutput(greeter)
	output += "\nVisit /chapter/23/server to start the fruit shop server on http://localhost:8001\n"
	return output
}

func chapter23Server(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Fruit Shop!")
	})
	mux.HandleFunc("/fruits", func(w http.ResponseWriter, r *http.Request) {
		fruits := []string{"Apple", "Banana", "Mango"}
		json.NewEncoder(w).Encode(fruits)
	})
	mux.HandleFunc("/fruit/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 3 {
			http.Error(w, "Fruit not specified", http.StatusBadRequest)
			return
		}
		fruit := parts[2]
		fmt.Fprintf(w, "You picked %s!", fruit)
	})
	go func() {
		log.Println("Starting Chapter 23 server at http://localhost:8001")
		log.Fatal(http.ListenAndServe(":8001", mux))
	}()
	fmt.Fprintln(w, "Chapter 23 server started at http://localhost:8001")
}

func chapter24() string {
	output := "Welcome to our API toy shop!\n"
	output += "\nVisit /chapter/24/server to start the API server on http://localhost:8002\n"
	return output
}

func chapter24Server(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the API Toy Shop!")
	})
	mux.HandleFunc("/courses", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(courses)
	})
	mux.HandleFunc("/course/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 3 {
			http.Error(w, "Course ID not specified", http.StatusBadRequest)
			return
		}
		id := parts[2]
		for _, course := range courses {
			if course.CourseId == id {
				json.NewEncoder(w).Encode(course)
				return
			}
		}
		http.Error(w, "Course not found", http.StatusNotFound)
	})
	mux.HandleFunc("/course", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var newCourse Course
			if err := json.NewDecoder(r.Body).Decode(&newCourse); err != nil {
				http.Error(w, "Error decoding JSON", http.StatusBadRequest)
				return
			}
			courses = append(courses, newCourse)
			json.NewEncoder(w).Encode(newCourse)
		case http.MethodPut:
			parts := strings.Split(r.URL.Path, "/")
			if len(parts) < 3 {
				http.Error(w, "Course ID not specified", http.StatusBadRequest)
				return
			}
			id := parts[2]
			var updatedCourse Course
			if err := json.NewDecoder(r.Body).Decode(&updatedCourse); err != nil {
				http.Error(w, "Error decoding JSON", http.StatusBadRequest)
				return
			}
			for i, course := range courses {
				if course.CourseId == id {
					courses[i] = updatedCourse
					json.NewEncoder(w).Encode(updatedCourse)
					return
				}
			}
			http.Error(w, "Course not found", http.StatusNotFound)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	go func() {
		log.Println("Starting Chapter 24 server at http://localhost:8002")
		log.Fatal(http.ListenAndServe(":8002", mux))
	}()
	fmt.Fprintln(w, "Chapter 24 server started at http://localhost:8002")
}

func chapter26() string {
	output := "GoRoutines in golang - LearnCodeOnline.in\n"
	wg := &sync.WaitGroup{}
	outputChan := make(chan string, 10)
	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		for i := 1; i <= 3; i++ {
			outputChan <- "Counting toy " + strconv.Itoa(i) + "\n"
			time.Sleep(time.Millisecond * 500)
		}
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		outputChan <- "Delivering message: Hello, toy shop!\n"
		outputChan <- "Delivering message: Toys are ready!\n"
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		outputChan <- "Starting slow task...\n"
		time.Sleep(time.Second * 2)
		outputChan <- "Slow task done!\n"
		wg.Done()
	}(wg)
	wg.Wait()
	close(outputChan)
	for msg := range outputChan {
		output += msg
	}
	output += "All toys counted and tasks done!\n"
	return output
}

func chapter27() string {
	output := "Race condition - LearnCodeOnline.in\n"
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	score := []int{0}
	outputChan := make(chan string, 10)
	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		outputChan <- "One R\n"
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		outputChan <- "Two R\n"
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		outputChan <- "Three R\n"
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	wg.Wait()
	close(outputChan)
	for msg := range outputChan {
		output += msg
	}
	output += "Final scoreboard: " + fmt.Sprint(score) + "\n"
	return output
}

func chapter28() string {
	output := "Channels in golang - LearnCodeOnline.in\n"
	myCh := make(chan int)
	wg := &sync.WaitGroup{}
	outputChan := make(chan string, 10)
	wg.Add(3)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 5
		ch <- 10
		wg.Done()
	}(myCh, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		close(ch)
		wg.Done()
	}(myCh, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh
		outputChan <- "Value: " + fmt.Sprint(val) + " Is channel open? " + fmt.Sprint(isChannelOpen) + "\n"
		val, isChannelOpen = <-myCh
		outputChan <- "Value: " + fmt.Sprint(val) + " Is channel open? " + fmt.Sprint(isChannelOpen) + "\n"
		wg.Done()
	}(myCh, wg)
	wg.Wait()
	close(outputChan)
	for msg := range outputChan {
		output += msg
	}
	return output
}

func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)
	go func() {
		var buf strings.Builder
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	f()
	w.Close()
	os.Stdout = old
	return <-outC
}

func createUser(name, email string, age int, phone, address string) user {
	return user{
		name:    name,
		email:   email,
		status:  false,
		age:     age,
		phone:   phone,
		address: address,
	}
}

func (u user) showDetails() {
	fmt.Printf("Name: %s, Email: %s, Verified: %t, Age: %d, Phone: %s, Address: %s\n",
		u.name, u.email, u.status, u.age, u.phone, u.address)
}

func (u *user) verifyEmail() {
	if strings.Contains(u.email, "@") && strings.HasSuffix(u.email, ".com") {
		u.status = true
		fmt.Println(u.name, "email verified!")
	}
}

func (u *user) updatePhone(newPhone string) {
	u.phone = newPhone
	fmt.Println(u.name, "phone updated to", newPhone)
}

func (u user) yearsUntilRetirement(retireAge int) int {
	years := retireAge - u.age
	if years < 0 {
		return 0
	}
	return years
}

func sayHello() {
	fmt.Println("Hello from the fruit shop!")
}

func greetUser(name string) {
	fmt.Println("Welcome,", name, ", to the fruit shop!")
}

func getFruitPrice(fruit string, quantity int) int {
	pricePerUnit := 2
	return pricePerUnit * quantity
}

func checkStock(fruit string, needed int) bool {
	stock := map[string]int{"Apple": 10, "Banana": 8, "Mango": 5}
	count, exists := stock[fruit]
	return exists && count >= needed
}

func calculateOrder(fruit string, quantity int) (int, int) {
	price := getFruitPrice(fruit, quantity)
	discount := 0
	if quantity > 5 {
		discount = price / 10
	}
	return price, discount
}

func buyFruits(fruits ...string) int {
	total := 0
	for _, fruit := range fruits {
		total += getFruitPrice(fruit, 2)
	}
	return total
}

func (u user) showProfile() {
	fmt.Printf("Profile: Name=%s, Email=%s, Verified=%t, Age=%d\n",
		u.name, u.email, u.status, u.age)
}

func (u *user) updateAge(newAge int) {
	u.age = newAge
	fmt.Println(u.name, "age updated to", newAge)
}

func (b fruitBasket) showContents() {
	fmt.Println(b.name, "contains:")
	for fruit, count := range b.fruits {
		fmt.Println(fruit, ":", count)
	}
}

func (b *fruitBasket) addFruit(fruit string, count int) {
	b.fruits[fruit] += count
	fmt.Println("Added", count, fruit, "to", b.name)
}

func (b *fruitBasket) removeFruit(fruit string) {
	delete(b.fruits, fruit)
	fmt.Println("Removed", fruit, "from", b.name)
}

func (p gamePlayer) showStatus() {
	fmt.Printf("Player %s: Score=%d, Rolls=%d\n", p.name, p.score, p.rolls)
}

func (p *gamePlayer) rollDice() {
	roll := mathRand.Intn(6) + 1
	p.score += roll
	p.rolls++
	fmt.Println(p.name, "rolled a", roll, "!")
}

func (p *gamePlayer) doubleScore() {
	p.score *= 2
	fmt.Println(p.name, "score doubled to", p.score)
}

func sellFruits() {
	defer cleanupShop()
	fmt.Println("Selling apples and bananas...")
	fmt.Println("Checking fruit stock...")
}

func cleanupShop() {
	fmt.Println("Sweeping the shop floor!")
	fmt.Println("Putting away fruit baskets!")
}

func manageSession(u user) {
	defer fmt.Println("Logging out", u.name)
	defer saveSessionLog(u.name)
	fmt.Println("Starting session for", u.name)
	fmt.Println("User browsing fruit catalog...")
}

func saveSessionLog(name string) {
	fmt.Println("Saving session log for", name)
}

func writeFruitOrder() {
	defer closeFile()
	fmt.Println("Opening fruit order file...")
	fmt.Println("Writing order: 5 apples, 3 mangos")
}

func closeFile() {
	fmt.Println("Closing fruit order file!")
}

func readFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Text in the file is:\n", string(data))
}

func appendToFile(filename, content string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	_, err = io.WriteString(file, content)
	if err != nil {
		fmt.Println("Error appending to file:", err)
	}
}

func fetchWebsite(urlStr string) {
	response, err := http.Get(urlStr)
	if err != nil {
		fmt.Println("Error fetching:", err)
		return
	}
	defer response.Body.Close()
	fmt.Printf("Response is of type: %T\n", response)
	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	content := string(dataBytes)
	if len(content) > 100 {
		fmt.Println("First 100 letters of content:\n", content[:100], "...")
	} else {
		fmt.Println("Content:\n", content)
	}
}

func parseURL(urlStr string) {
	result, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println("Scheme (like https):", result.Scheme)
	fmt.Println("Host (like lco.dev):", result.Host)
	fmt.Println("Path (like /learn):", result.Path)
	fmt.Println("Port (like 3000):", result.Port())
	fmt.Println("Raw Query (like coursename=reactjs):", result.RawQuery)
	qparams := result.Query()
	fmt.Printf("Query params are of type: %T\n", qparams)
	fmt.Println("Course/Fruit param:", qparams.Get("coursename"), qparams.Get("fruit"))
	fmt.Println("All query params:")
	for key, values := range qparams {
		for _, val := range values {
			fmt.Printf("Param %s: %s\n", key, val)
		}
	}
}

func buildFruitShopURL(scheme, host, path string, queryParams map[string]string) string {
	partsOfUrl := &url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}
	q := url.Values{}
	for key, value := range queryParams {
		q.Add(key, value)
	}
	partsOfUrl.RawQuery = q.Encode()
	return partsOfUrl.String()
}

func PerformGetRequest(myurl string) {
	fmt.Println("Fetching GET from", myurl)
	response, err := http.Get(myurl)
	if err != nil {
		fmt.Println("Oops, couldn't reach", myurl, ":", err)
		return
	}
	defer response.Body.Close()
	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength, "letters")
	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	content := string(dataBytes)
	if len(content) > 100 {
		fmt.Println("First 100 letters:\n", content[:100], "...")
	} else {
		fmt.Println("Content:\n", content)
	}
}

func EncodeJson() {
	lcoCourses := []course{
		{
			Name:     "ReactJS Bootcamp",
			Price:    299,
			Website:  "FruitLearnOnline.in",
			Password: "abc123",
			Tags:     []string{"web", "javascript"},
		},
		{
			Name:     "MERN Bootcamp",
			Price:    199,
			Website:  "FruitLearnOnline.in",
			Password: "bcd123",
			Tags:     []string{"mern", "fullstack"},
		},
		{
			Name:     "Angular Bootcamp",
			Price:    299,
			Website:  "FruitLearnOnline.in",
			Password: "hit123",
			Tags:     nil,
		},
	}
	jsonData, err := json.MarshalIndent(lcoCourses, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println("\nEncoded JSON:")
	fmt.Println(string(jsonData))
}

func DecodeJson() {
	jsonData := []byte(`
	[
		{
			"coursename": "Go Bootcamp",
			"price": 99,
			"website": "FruitLearnOnline.in",
			"tags": ["go", "backend"]
		},
		{
			"coursename": "Python Bootcamp",
			"price": 149,
			"website": "FruitLearnOnline.in",
			"tags": null
		}
	]
	`)
	if !json.Valid(jsonData) {
		fmt.Println("\nOops, bad JSON words from file!")
		return
	}
	var courses []course
	err := json.Unmarshal(jsonData, &courses)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("\nDecoded courses from file:")
	for i, c := range courses {
		fmt.Printf("Course %d: Name=%s, Price=%d, Website=%s, Tags=%v\n",
			i+1, c.Name, c.Price, c.Website, c.Tags)
	}
}

func ConsumeWebJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "NodeJS Bootcamp",
		"price": 299,
		"website": "FruitLearnOnline.in",
		"extra": "Special Offer"
	}
	`)
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("\nJSON from web was valid!")
		var lcoCourse course
		err := json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		if err != nil {
			fmt.Println("Error decoding web JSON:", err)
			return
		}
		fmt.Printf("Web course (struct): %#v\n", lcoCourse)
		var myonlineData map[string]interface{}
		err = json.Unmarshal(jsonDataFromWeb, &myonlineData)
		if err != nil {
			fmt.Println("Error decoding web JSON to map:", err)
			return
		}
		fmt.Println("Web course (map):")
		for key, value := range myonlineData {
			fmt.Printf("  %s: %v\n", key, value)
		}
	} else {
		fmt.Println("\nOops, bad JSON from web!")
	}
	jsonDataExtra := []byte(`
	{
		"coursename": "Java Bootcamp",
		"price": 249,
		"website": "FruitLearnOnline.in",
		"tags": ["java", "backend"],
		"discount": 10
	}
	`)
	if json.Valid(jsonDataExtra) {
		var extraData map[string]interface{}
		err := json.Unmarshal(jsonDataExtra, &extraData)
		if err != nil {
			fmt.Println("Error decoding extra JSON:", err)
			return
		}
		fmt.Println("\nExtra web course (map):")
		for key, value := range extraData {
			fmt.Printf("  %s: %v\n", key, value)
		}
	} else {
		fmt.Println("\nOops, bad JSON from web!")
	}
}

func greeter() {
	fmt.Println("Hello mod in golang!")
}