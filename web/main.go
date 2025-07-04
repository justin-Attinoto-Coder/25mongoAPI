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
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

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
	router.HandleFunc("/chapter/{id}", chapterHandler).Methods("GET")

	log.Println("Starting server at http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := chapterCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch chapters", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var chapters []Chapter
	if err := cursor.All(ctx, &chapters); err != nil {
		http.Error(w, "Failed to decode chapters", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, chapters)
}

func chapterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	tmpl, err := template.ParseFiles("templates/chapter.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
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
		ChapterID string
		Output    string
	}{ChapterID: id, Output: output}
	tmpl.Execute(w, data)
}

func chapter01() string {
	output := "Welcome to Go Programming!\n"
	name := "Justin"
	age := 41
	output += fmt.Sprintf("Hello, %s! You are %d years old.\n", name, age)
	yearsLater := age + 5
	output += fmt.Sprintf("In 5 years, %s will be %d years old.\n", name, yearsLater)
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
	output += fmt.Sprintf("%s\n", username)
	output += fmt.Sprintf("variable is of type: %T \n", username)
	output += fmt.Sprintf("Age: %d (Type: %T)\n", age, age)
	output += fmt.Sprintf("City: %s (Type: %T)\n", city, city)
	output += fmt.Sprintf("Height: %.1f (Type: %T)\n", height, height)
	output += fmt.Sprintf("Is Student: %t (Type: %T)\n", isStudent, isStudent)
	output += fmt.Sprintf("Birth Year: %d (Type: %T)\n", birthYear, birthYear)
	output += fmt.Sprintf("Login Token: %s (Type: %T)\n", loginToken, loginToken)
	output += fmt.Sprintf("%v\n", numberOfUser)
	output += fmt.Sprintf("Number of Users: %.1f (Type: %T)\n", numberOfUser, numberOfUser)
	yearsSinceBirth := 2025 - birthYear
	output += fmt.Sprintf("%s, you are %d years old in 2025.\n", username, yearsSinceBirth)
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
			output += fmt.Sprintf("Input %d is a string: %s\n", i+1, value)
		} else {
			output += fmt.Sprintf("Input %d is not a string\n", i+1)
		}
	}
	output += "Enter a short bio: [Input not available in web mode]\n"
	bio := "I love coding in Go!"
	output += fmt.Sprintf("Your bio: %s\n", bio)
	runeCount := len([]rune(bio))
	output += fmt.Sprintf("Your bio has %d runes\n", runeCount)
	output += "Enter a favorite word: [Input not available in web mode]\n"
	word := "Go"
	output += fmt.Sprintf("Your favorite word: %s\n", word)
	return output
}

func chapter04() string {
	output := "Welcome to our pizza app\n"
	output += "Please rate our pizza between 1 and 5\n"
	output += "Enter rating: [Input not available in web mode]\n"
	input := "4.5"
	output += fmt.Sprintf("Thanks for rating, %s\n", input)
	numRating, err := strconv.ParseFloat(input, 64)
	if err != nil {
		output += fmt.Sprintf("Error converting rating to float: %v\n", err)
		return output
	}
	output += fmt.Sprintf("Added 1 to your rating: %.1f\n", numRating+1)
	ratingStr := strconv.FormatFloat(numRating, 'f', 1, 64)
	output += fmt.Sprintf("Your rating as string: %s\n", ratingStr)
	output += "Enter number of pizzas ordered: [Input not available in web mode]\n"
	pizzaInput := "3"
	numPizzas, err := strconv.Atoi(pizzaInput)
	if err != nil {
		output += fmt.Sprintf("Error converting to integer: %v\n", err)
		return output
	}
	output += fmt.Sprintf("You ordered %d pizzas\n", numPizzas)
	pizzaCountStr := strconv.Itoa(numPizzas)
	output += fmt.Sprintf("Number of pizzas as string: %s\n", pizzaCountStr)
	output += "Are you a premium member? (true/false): [Input not available in web mode]\n"
	premiumInput := "true"
	isPremium, err := strconv.ParseBool(premiumInput)
	if err != nil {
		output += fmt.Sprintf("Error converting to boolean: %v\n", err)
		return output
	}
	output += fmt.Sprintf("Premium member status: %t\n", isPremium)
	premiumStr := strconv.FormatBool(isPremium)
	output += fmt.Sprintf("Premium status as string: %s\n", premiumStr)
	var data interface{} = numRating
	if val, ok := data.(float64); ok {
		output += fmt.Sprintf("Rating is a float64: %.1f\n", val)
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
	output += fmt.Sprintf("Adding %v and %v gives us %v!\n", number1, number2, sum)
	difference := number1 - float64(number2)
	output += fmt.Sprintf("Taking %v away from %v leaves %v!\n", number2, number1, difference)
	product := number1 * float64(number2)
	output += fmt.Sprintf("%v groups of %v makes %v!\n", number2, number1, product)
	quotient := number1 / float64(number2)
	output += fmt.Sprintf("Sharing %v among %v gives %v each!\n", number1, number2, quotient)
	rounded := math.Round(number1)
	output += fmt.Sprintf("Rounding %v gives us %v!\n", number1, rounded)
	ceiling := math.Ceil(number1)
	output += fmt.Sprintf("Ceiling of %v is %v!\n", number1, ceiling)
	floor := math.Floor(number1)
	output += fmt.Sprintf("Floor of %v is %v!\n", number1, floor)
	maxNumber := math.Max(number1, float64(number2))
	output += fmt.Sprintf("The bigger number between %v and %v is %v!\n", number1, number2, maxNumber)
	power := math.Pow(2, 3)
	output += fmt.Sprintf("2 grown 3 times is %v!\n", power)
	sqrt := math.Sqrt(16)
	output += fmt.Sprintf("The square root of 16 is %v!\n", sqrt)
	mathRand.Seed(time.Now().UnixNano())
	output += fmt.Sprintf("Random number: %v\n", mathRand.Intn(5))
	output += fmt.Sprintf("Crypto random number: %v\n", mathRand.Intn(5))
	return output
}

func chapter06() string {
	output := "Welcome to time study of golang\n"
	presentTime := time.Now()
	output += fmt.Sprintf("Right now, the time is: %v\n", presentTime)
	formattedTime := presentTime.Format("01-02-2006 15:04:05 Monday")
	output += fmt.Sprintf("Pretty time is: %s\n", formattedTime)
	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	output += fmt.Sprintf("Special date is: %v\n", createdDate)
	formattedDate := createdDate.Format("01-02-2006 Monday")
	output += fmt.Sprintf("Pretty special date is: %s\n", formattedDate)
	futureTime := presentTime.Add(2 * time.Hour)
	output += fmt.Sprintf("Two hours from now: %v\n", futureTime)
	timeDifference := futureTime.Sub(presentTime)
	output += fmt.Sprintf("Time difference is: %v hours\n", timeDifference.Hours())
	isBefore := createdDate.Before(presentTime)
	output += fmt.Sprintf("Is special date before now? %t\n", isBefore)
	shortFormat := presentTime.Format("2006-01-02")
	output += fmt.Sprintf("Short date format: %s\n", shortFormat)
	timeOnly := presentTime.Format("15:04")
	output += fmt.Sprintf("Time only format: %s\n", timeOnly)
	return output
}

func chapter07() string {
	output := "Welcome to building Go programs for different computers!\n"
	hostOS := runtime.GOOS
	output += fmt.Sprintf("Our computer's operating system (GOHOSTOS): %s\n", hostOS)
	output += fmt.Sprintf("Default target operating system (GOOS): %s\n", runtime.GOOS)
	user := "Justin"
	output += fmt.Sprintf("Hello from %s on %s !\n", user, hostOS)
	switch hostOS {
	case "windows":
		output += "Yay, we're on Windows! This program loves .exe files!\n"
	case "linux":
		output += "Cool, we're on Linux! This program runs without extensions!\n"
	case "darwin":
		output += "Awesome, we're on a Mac! This program loves macOS!\n"
	}
	goArch := os.Getenv("GOARCH")
	output += fmt.Sprintf("Our computer's brain type (GOARCH): %s\n", goArch)
	goPath := os.Getenv("GOPATH")
	output += fmt.Sprintf("Where Go keeps its toys (GOPATH): %s\n", goPath)
	return output
}

func chapter08() string {
	output := "Welcome to our fruit basket adventure with arrays!\n"
	var fruitBasket [4]string
	fruitBasket[0] = "Apple"
	fruitBasket[1] = "Banana"
	fruitBasket[2] = "Orange"
	fruitBasket[3] = "Mango"
	output += fmt.Sprintf("My fruit basket: %v\n", fruitBasket)
	output += fmt.Sprintf("First fruit in slot 0: %s\n", fruitBasket[0])
	fruitBasket[2] = "Grape"
	output += fmt.Sprintf("New fruit basket: %v\n", fruitBasket)
	basketSize := len(fruitBasket)
	output += fmt.Sprintf("My basket has %d slots\n", basketSize)
	veggieBasket := [3]string{"Carrot", "Potato", "Cucumber"}
	output += fmt.Sprintf("My veggie basket: %v\n", veggieBasket)
	output += "Checking each fruit in the fruit basket:\n"
	for i := 0; i < len(fruitBasket); i++ {
		output += fmt.Sprintf("Slot %d has %s\n", i, fruitBasket[i])
	}
	output += "Using magic hand to check veggies:\n"
	for index, veggie := range veggieBasket {
		output += fmt.Sprintf("Slot %d has %s\n", index, veggie)
	}
	fruitCount := [5]int{10, 20, 15, 30, 25}
	output += fmt.Sprintf("Fruit counts: %v\n", fruitCount)
	totalFruits := 0
	for _, count := range fruitCount {
		totalFruits += count
	}
	output += fmt.Sprintf("Total fruits in the big basket: %d\n", totalFruits)
	doubleBasket := [2][3]string{
		{"Apple", "Banana", "Grape"},
		{"Mango", "Orange", "Pineapple"},
	}
	output += fmt.Sprintf("My double basket: %v\n", doubleBasket)
	output += fmt.Sprintf("Fruit in row 1, column 2: %s\n", doubleBasket[0][1])
	output += "Checking all fruits in the double basket:\n"
	for row := 0; row < len(doubleBasket); row++ {
		for col := 0; col < len(doubleBasket[row]); col++ {
			output += fmt.Sprintf("Row %d Column %d has %s\n", row, col, doubleBasket[row][col])
		}
	}
	return output
}

func chapter09() string {
	mathRand.Seed(time.Now().UnixNano())
	output := "Welcome to our slice adventure with fruits, scores, and courses!\n"
	fruitBasket := []string{"Apple", "Tomato", "Peach"}
	output += fmt.Sprintf("My stretchy fruit basket: %v\n", fruitBasket)
	fruitBasket = append(fruitBasket, "Mango", "Banana")
	output += fmt.Sprintf("After adding Mango and Banana: %v\n", fruitBasket)
	fruitBasket = append(fruitBasket[:1], fruitBasket[2:]...)
	output += fmt.Sprintf("Fruit basket after removing Tomato: %v\n", fruitBasket)
	highScores := []int{
		mathRand.Intn(1000) + 1,
		mathRand.Intn(1000) + 1,
		mathRand.Intn(1000) + 1,
		mathRand.Intn(1000) + 1,
	}
	output += fmt.Sprintf("My game high scores: %v\n", highScores)
	newScore := mathRand.Intn(1000) + 1
	highScores = append(highScores, newScore)
	output += fmt.Sprintf("After adding new score %d: %v\n", newScore, highScores)
	highScores = append(highScores[:1], highScores[2:]...)
	output += fmt.Sprintf("High scores after removing second score: %v\n", highScores)
	courses := []string{"ReactJS", "JavaScript", "Swift", "Python", "Ruby"}
	output += fmt.Sprintf("My courses box: %v\n", courses)
	newCourseIndex := len(courses)
	newCourse := fmt.Sprintf("Course%d", newCourseIndex+1)
	courses = append(courses, newCourse)
	output += fmt.Sprintf("After adding %s: %v\n", newCourse, courses)
	courses = append(courses[:2], courses[3:]...)
	output += fmt.Sprintf("Courses after removing Swift: %v\n", courses)
	coursesSize := len(courses)
	coursesCapacity := cap(courses)
	output += fmt.Sprintf("Courses box has %d courses and room for %d\n", coursesSize, coursesCapacity)
	output += fmt.Sprintf("Course in slot 0: %s\n", courses[0])
	courses[1] = "NodeJS"
	output += fmt.Sprintf("New courses box: %v\n", courses)
	someCourses := courses[1:4]
	output += fmt.Sprintf("Some courses: %v\n", someCourses)
	output += "Checking each course:\n"
	for i := 0; i < len(courses); i++ {
		output += fmt.Sprintf("Slot %d has course %s\n", i, courses[i])
	}
	output += "Using magic hand to check courses:\n"
	for index, course := range courses {
		output += fmt.Sprintf("Slot %d has course %s\n", index, course)
	}
	someCourses[0] = "Angular"
	output += fmt.Sprintf("Changed some courses: %v\n", someCourses)
	output += fmt.Sprintf("Original courses after change: %v\n", courses)
	coursesCopy := make([]string, len(courses))
	copy(coursesCopy, courses)
	output += fmt.Sprintf("Copied courses: %v\n", coursesCopy)
	coursesCopy[0] = "VueJS"
	output += fmt.Sprintf("Changed copy: %v\n", coursesCopy)
	output += fmt.Sprintf("Original courses unchanged: %v\n", courses)
	return output
}

func chapter10() string {
	output := "Welcome to our magic toy box adventure with maps!\n"
	fruitInventory := make(map[string]int)
	fruitInventory["Apple"] = 10
	fruitInventory["Banana"] = 15
	fruitInventory["Mango"] = 8
	output += fmt.Sprintf("My fruit inventory: %v\n", fruitInventory)
	output += "Checking all fruits with a magic hand:\n"
	for fruit, count := range fruitInventory {
		output += fmt.Sprintf("Tag %s has %d fruits\n", fruit, count)
	}
	bananaCount := fruitInventory["Banana"]
	output += fmt.Sprintf("Number of bananas: %d\n", bananaCount)
	count, exists := fruitInventory["Mango"]
	if exists {
		output += fmt.Sprintf("Found Mango with %d items!\n", count)
	} else {
		output += "No Mango in the box!\n"
	}
	delete(fruitInventory, "Mango")
	output += fmt.Sprintf("After removing Mango: %v\n", fruitInventory)
	studentScores := map[string]float64{
		"Justin": 85.5,
		"Alice":  92.0,
		"Bob":    78.5,
	}
	output += fmt.Sprintf("Student scores: %v\n", studentScores)
	output += "Checking all student scores with a magic hand:\n"
	for student, score := range studentScores {
		output += fmt.Sprintf("Student %s has score %v\n", student, score)
	}
	studentScores["Emma"] = 88.0
	output += fmt.Sprintf("After adding Emma: %v\n", studentScores)
	courseSchedule := make(map[string]string)
	courseSchedule["ReactJS"] = "Monday 9AM"
	courseSchedule["Python"] = "Tuesday 2PM"
	courseSchedule["Swift"] = "Wednesday 11AM"
	output += fmt.Sprintf("Course schedule: %v\n", courseSchedule)
	output += "Checking all courses with a magic hand:\n"
	for course, time := range courseSchedule {
		output += fmt.Sprintf("Course %s is at %s\n", course, time)
	}
	courseSchedule["Swift"] = "Friday 10AM"
	output += fmt.Sprintf("Updated course schedule: %v\n", courseSchedule)
	languages := map[string]string{
		"Go":         "Compiled",
		"Python":     "Interpreted",
		"JavaScript": "Interpreted",
		"Ruby":       "Interpreted",
		"Swift":      "Compiled",
	}
	output += fmt.Sprintf("Programming languages: %v\n", languages)
	output += "Checking all languages with a magic hand:\n"
	for lang, langType := range languages {
		output += fmt.Sprintf("Language %s is %s\n", lang, langType)
	}
	languages["Java"] = "Compiled"
	output += fmt.Sprintf("After adding Java: %v\n", languages)
	langType, exists := languages["Python"]
	if exists {
		output += fmt.Sprintf("Found Python, it's %s !\n", langType)
	} else {
		output += "No Python in the box!\n"
	}
	delete(languages, "Ruby")
	output += fmt.Sprintf("After removing Ruby: %v\n", languages)
	langCount := len(languages)
	output += fmt.Sprintf("Number of languages in the box: %d\n", langCount)
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
	output += fmt.Sprintf("Justin's toy box: %+v\n", justin)
	output += captureOutput(func() { justin.showDetails() })
	justin.verifyEmail()
	output += fmt.Sprintf("Justin's status after verification: %t\n", justin.status)
	justin.updatePhone("555-9999")
	output += fmt.Sprintf("Justin's new phone: %s\n", justin.phone)
	yearsToRetire := justin.yearsUntilRetirement(65)
	output += fmt.Sprintf("Years until Justin retires: %d\n", yearsToRetire)
	alice := createUser("Alice", "alice@example.com", 25, "555-5678", "456 Banana Avenue")
	output += fmt.Sprintf("Alice's toy box: %+v\n", alice)
	output += captureOutput(func() { alice.showDetails() })
	alice.verifyEmail()
	output += fmt.Sprintf("Alice's status after verification: %t\n", alice.status)
	users := []user{justin, alice}
	output += fmt.Sprintf("All user toy boxes: %v\n", users)
	output += "Checking all users with helpers:\n"
	for i, u := range users {
		output += fmt.Sprintf("User %d:\n", i)
		output += captureOutput(func() { u.showDetails() })
		output += fmt.Sprintf("Years to retire: %d\n", u.yearsUntilRetirement(65))
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
		output += fmt.Sprintf("Justin can vote! He's %d years old.\n", justin.age)
	} else {
		output += fmt.Sprintf("Justin is too young to vote. He's only %d\n", justin.age)
	}
	if justin.status {
		output += "Justin's email is verified! Welcome!\n"
	} else {
		output += "Justin, please verify your email!\n"
		output += captureOutput(func() { justin.verifyEmail() })
		output += fmt.Sprintf("After verification, status is: %t\n", justin.status)
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
		output += fmt.Sprintf("Login okay! You tried %d times.\n", loginAttempts)
	} else {
		output += fmt.Sprintf("Too many tries! %d is more than %d\n", loginAttempts, maxAttempts)
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
		output += fmt.Sprintf("User %d:\n", i)
		if u.status {
			output += fmt.Sprintf("%s is verified!\n", u.name)
		} else {
			output += fmt.Sprintf("%s needs to verify their email!\n", u.name)
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
		output += fmt.Sprintf("Value of dice is %d\n", diceNumber)
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
		output += fmt.Sprintf("You're now at %s\n", board[playerPosition])
		output += "---\n"
	}
	if playerPosition == maxPosition {
		output += "Yay! You reached the Finish! You win the fruit game!\n"
	}
	output += "\nLet's roll one more dice for fun!\n"
	anotherDice := mathRand.Intn(6) + 1
	output += fmt.Sprintf("Dice value: %d\n", anotherDice)
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
		output += fmt.Sprintf("Value of dice is %d\n", diceNumber)
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
		output += fmt.Sprintf("You're now at %s\n", board[playerPosition])
		output += "---\n"
	}
	if playerPosition == maxPosition {
		output += "Yay! You reached the Finish! You win the fruit game!\n"
	}
	output += "\nLet's roll one more dice for fun!\n"
	anotherDice := mathRand.Intn(6) + 1
	output += fmt.Sprintf("Dice value: %d\n", anotherDice)
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
	output += fmt.Sprintf("Price for 2 apples: %d coins\n", applePrice)
	haveBananas := checkStock("Banana", 5)
	output += fmt.Sprintf("Enough bananas for 5? %t\n", haveBananas)
	total, discount := calculateOrder("Mango", 10)
	output += fmt.Sprintf("Total for 10 mangos: %d coins, Discount: %d coins\n", total, discount)
	totalCost := buyFruits("Apple", "Banana", "Mango")
	output += fmt.Sprintf("Total cost for fruits: %d coins\n", totalCost)
	countFruits := func() int { return 3 }
	output += fmt.Sprintf("Number of fruit types: %d\n", countFruits())
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
		output += fmt.Sprintf("Error creating file: %v\n", err)
		return output
	}
	defer os.Remove(file.Name())
	length, err := io.WriteString(file, content)
	if err != nil {
		output += fmt.Sprintf("Error writing to file: %v\n", err)
		return output
	}
	output += fmt.Sprintf("Wrote %d letters to the file!\n", length)
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
	output += fmt.Sprintf("\nFetching from %s\n", url)
	output += captureOutput(func() { fetchWebsite(url) })
	const fruitUrl = "https://example.com"
	output += fmt.Sprintf("\nFetching from %s\n", fruitUrl)
	output += captureOutput(func() { fetchWebsite(fruitUrl) })
	output += "\nStarting fruit shop web server on http://localhost:8080 [Not running in web mode]\n"
	output += "Visit http://localhost:8080/ for welcome page\n"
	output += "Visit http://localhost:8080/fruits for fruit list\n"
	return output
}

func chapter20() string {
	output := "Welcome to handling URLs in golang!\n"
	const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"
	output += fmt.Sprintf("\nParsing URL: %s\n", myurl)
	output += captureOutput(func() { parseURL(myurl) })
	const fruitUrl = "https://fruitshop.com:8080/order?fruit=apple&quantity=5&customer=Justin"
	output += fmt.Sprintf("\nParsing fruit shop URL: %s\n", fruitUrl)
	output += captureOutput(func() { parseURL(fruitUrl) })
	output += "\nBuilding a new fruit shop URL\n"
	newUrl := buildFruitShopURL("https", "fruitshop.com", "/cart", map[string]string{
		"fruit":    "mango",
		"quantity": "3",
		"promo":    "FRUIT10",
	})
	output += fmt.Sprintf("New URL: %s\n", newUrl)
	return output
}

func chapter21() string {
	output := "Welcome to our fruit shop web server adventure!\n"
	output += "\nMaking GET and POST requests\n"
	output += captureOutput(func() { PerformGetRequest("https://example.com") })
	output += "\nStarting fruit shop web server on http://localhost:8000 [Not running in web mode]\n"
	output += "Visit http://localhost:8000/ for welcome page\n"
	output += "Visit http://localhost:8000/get for message\n"
	output += "POST to http://localhost:8000/post or /postform for JSON/form handling\n"
	return output
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
	output += "\nStarting fruit shop server at http://localhost:8000 [Not running in web mode]\n"
	output += "Visit http://localhost:8000/ for welcome page\n"
	output += "Visit http://localhost:8000/fruits for fruit list\n"
	output += "Visit http://localhost:8000/fruit/Mango to pick a fruit\n"
	return output
}

func chapter24() string {
	output := "Welcome to our API toy shop!\n"
	output += "\nStarting API server at http://localhost:8000 [Not running in web mode]\n"
	output += "Visit http://localhost:8000/ for welcome page\n"
	output += "GET http://localhost:8000/courses for all courses\n"
	output += "GET http://localhost:8000/course/1 for one course\n"
	output += "POST http://localhost:8000/course to create a course\n"
	output += "PUT http://localhost:8000/course/1 to update a course\n"
	return output
}

func chapter26() string {
	output := "GoRoutines in golang - LearnCodeOnline.in\n"
	wg := &sync.WaitGroup{}
	outputChan := make(chan string, 10)
	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		for i := 1; i <= 3; i++ {
			outputChan <- fmt.Sprintf("Counting toy %d\n", i)
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
	output += fmt.Sprintf("Final scoreboard: %v\n", score)
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
		outputChan <- fmt.Sprintf("Value: %v Is channel open? %v\n", val, isChannelOpen)
		val, isChannelOpen = <-myCh
		outputChan <- fmt.Sprintf("Value: %v Is channel open? %v\n", val, isChannelOpen)
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

func getFruitPrice(_ string, quantity int) int {
	pricePerUnit := 2
	return pricePerUnit * quantity
}

func checkStock(fruit string, needed int) bool {
	stock := map[string]int{"Apple": 10, "Banana": 8, "Mango": 5}
	return stock[fruit] >= needed
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
		fmt.Println("\nOops, bad extra JSON from web!")
	}
}

func greeter() {
	fmt.Println("Hello mod in golang!")
}

// Five key lessons for a 5-year-old:
// 1. A web server is like a toy shop where you pick chapters to play with.
// 2. MongoDB stores chapter names like a list of toy boxes.
// 3. HTML templates are like drawing boards to show chapter results.
// 4. Routes are like shop signs pointing to different chapter games.
// 5. Channels collect toy messages to show them neatly on the webpage.

// Five key lessons for Chapter 01:
// 1. Printing is like shouting a message to show everyone.
// 2. Variables are like toy boxes holding names or numbers.
// 3. Formatting makes messages pretty, like decorating a toy.
// 4. Adding numbers is like putting more toys in a pile.
// 5. Go programs start with a main function, like a game’s start button.

// Five key lessons for Chapter 02:
// 1. Variables are like toy boxes that hold different kinds of toys.
// 2. Types tell us what kind of toy (like numbers or words) goes in the box.
// 3. Constants are like toys that never change, like a favorite teddy bear.
// 4. Short declarations make quick toy boxes for fast play.
// 5. Printing types helps us check what’s inside our toy boxes.

// Five key lessons for Chapter 03:
// 1. User input is like asking a friend to share their favorite toy.
// 2. A scanner is a tool that listens to what you type, like a magic ear.
// 3. Strings are words you clean up with TrimSpace to remove extra spaces.
// 4. Type assertion checks if a toy box holds the right kind of toy.
// 5. Runes count the special letters in your words, like counting toy blocks.

// Five key lessons for Chapter 04:
// 1. Conversion changes a toy’s type, like turning a word into a number.
// 2. A reader listens to what you type, like a magic ear for words.
// 3. ParseFloat turns a word into a number with decimals, like counting candies.
// 4. Atoi turns a word into a whole number, like counting pizzas.
// 5. Type assertion checks if a toy box has the right kind of toy inside.

// Five key lessons for Chapter 05:
// 1. Math operations like adding and sharing toys help us play with numbers.
// 2. The math package is a toy box with tools like rounding and square roots.
// 3. Random numbers are like picking a surprise toy from a box.
// 4. Crypto random numbers are super safe surprises for important games.
// 5. Import aliases keep different toy boxes from getting mixed up.

// Five key lessons for Chapter 06:
// 1. Time is like a clock toy that tells us now or later.
// 2. Formatting time makes it pretty, like choosing a clock style.
// 3. Adding time is like fast-forwarding a clock for future fun.
// 4. Comparing times is like asking if lunch is before dinner.
// 5. Short formats show just the date or time, like a quick peek.

// Five key lessons for Chapter 07:
// 1. GOOS tells us what kind of computer toy box we’re using, like Windows.
// 2. GOARCH is like checking if the computer’s brain is big or small.
// 3. GOPATH is where Go keeps its toys, like a toy storage room.
// 4. Checking the OS is like knowing if we’re playing on a Mac or Linux.
// 5. Printing user info is like saying hi from your computer.

// Five key lessons for Chapter 08:
// 1. Arrays are like toy boxes with a fixed number of slots for toys.
// 2. You can swap toys in array slots, like changing an Orange to a Grape.
// 3. Loops check each slot, like looking at every toy in the box.
// 4. Double arrays are like boxes with rows and columns for more toys.
// 5. Length tells us how many toys are in the array box.

// Five key lessons for Chapter 09:
// 1. Slices are stretchy toy boxes that can grow or shrink.
// 2. Append adds new toys to a slice, like more fruits.
// 3. Removing toys from a slice is like skipping one in the middle.
// 4. Copies make a new slice box so changes don’t mess up the original.
// 5. Length and capacity tell us how many toys fit now and later.

// Five key lessons for Chapter 10:
// 1. Maps are magic toy boxes with name tags for each toy.
// 2. Adding toys to a map is like sticking a tag on a fruit.
// 3. Removing toys uses a tag, like taking out a Mango.
// 4. Checking tags tells us if a toy is in the map or not.
// 5. Loops show all toys and tags, like checking a fruit list.

// Five key lessons for Chapter 11:
// 1. Structs are custom toy boxes holding different toys together.
// 2. Methods are helpers that play with a toy box, like showing toys.
// 3. Pointers let helpers change toys in the box, like updating a phone.
// 4. Lists of structs are like a shelf of toy boxes for many users.
// 5. Helpers can count years, like figuring out retirement time.

// Five key lessons for Chapter 12:
// 1. If statements are like deciding if a toy is right, like voting age.
// 2. Else statements do something else if the toy isn’t right.
// 3. Multiple checks sort toys, like calling someone a kid or adult.
// 4. Combining checks is like asking if two toys are true together.
// 5. Helpers verify toys, like checking an email for a stamp.

// Five key lessons for Chapter 13:
// 1. Switch statements sort toys, like picking a move from a dice roll.
// 2. Each case is a path, like moving one spot for a dice roll of 1.
// 3. Continue skips to the next game turn, like rolling again on a 6.
// 4. A fruit board is like a game path with toy spots to land on.
// 5. Random dice make the game fun with surprise moves.

// Five key lessons for Chapter 14:
// 1. Loops keep playing a game until you reach the end, like a finish line.
// 2. Switch sorts dice rolls to decide how far you move on a fruit board.
// 3. Random numbers mix up dice for different moves each time.
// 4. Checking position stops you from going past the finish.
// 5. A second switch style picks prizes, like a big fruit basket.

// Five key lessons for Chapter 15:
// 1. Functions are helpers that do jobs, like counting fruit coins.
// 2. Parameters are toys a helper needs, like a fruit name.
// 3. Returning toys gives back answers, like a price or discount.
// 4. Flexible functions take many toys, like a big fruit shopping list.
// 5. Quick helpers count toys without a big name, like fruit types.

// Five key lessons for Chapter 16:
// 1. Methods are special helpers that play with toy boxes, like users.
// 2. Structs hold toys like names or scores in one box.
// 3. Pointers let helpers change toys, like updating an age.
// 4. Fruit baskets use methods to add or remove toys like mangos.
// 5. Game players use methods to roll dice and double scores.

// Five key lessons for Chapter 17:
// 1. Defer is like promising to clean up after playing with toys.
// 2. Cleaning a shop happens after selling fruits, like sweeping.
// 3. Logging out a user is saved for the end of their visit.
// 4. Closing a file is like shutting a notebook after writing.
// 5. Multiple defers stack up and run backward, like last chore first.

// Five key lessons for Chapter 18:
// 1. Files are like notebooks for writing and reading toy lists.
// 2. Writing to a file is like adding fruit orders to a notebook.
// 3. Reading a file shows all the words, like checking orders.
// 4. Appending adds more words, like extra mangos to the list.
// 5. Defer closes the notebook to keep it safe after writing.

// Five key lessons for Chapter 19:
// 1. Web requests are like asking a website for toys or sending toys.
// 2. GET requests grab toy pages, like a fruit shop menu.
// 3. A server is like a shop giving out pages to visitors.
// 4. Closing responses is like shutting a toy box after use.
// 5. Fruit lists show up as web pages, like a shop’s fruit menu.

// Five key lessons for Chapter 20:
// 1. URLs are like maps to a website, like a fruit shop address.
// 2. Parsing URLs splits the map into parts, like shop name.
// 3. Query params are extra notes, like a fruit order.
// 4. Building URLs creates new maps with toys like mangos.
// 5. Checking params is like reading notes for course or fruit.

// Five key lessons for Chapter 21:
// 1. GET requests ask for toys, like a message from a shop.
// 2. POST requests send toy boxes, like a course or fruit order.
// 3. JSON sends fancy toy boxes, while forms send simple notes.
// 4. Servers read notes to reply with toys, like a fruit order.
// 5. Closing responses keeps the shop tidy after visitors.

// Five key lessons for Chapter 22:
// 1. JSON turns toy boxes into words to send to friends.
// 2. Decoding JSON makes toy boxes from words, like courses.
// 3. Maps are magic boxes for any toys from web JSON.
// 4. Checking JSON ensures words are good before opening.
// 5. Tags hide secret toys, like passwords, from JSON.

// Five key lessons for Chapter 23:
// 1. Modules are big toy boxes with names to hold code.
// 2. Gorilla/mux is a helper for making shop stalls.
// 3. Routes are like signs pointing to fruit or home pages.
// 4. Picking a fruit by name is like choosing a toy.
// 5. Starting a server opens a shop for web visitors.

// Five key lessons for Chapter 24:
// 1. APIs are toy shops where friends get or send course toys.
// 2. GET routes fetch all toys or one by ID, like a course.
// 3. POST routes add new toys, like a new course.
// 4. PUT routes change toys, keeping the same ID.
// 5. JSON sends course toys to visitors, like a shop list.

// Five key lessons for Chapter 26:
// 1. Goroutines are like kids doing different jobs at the same time.
// 2. The "go" keyword tells a kid to start working without waiting.
// 3. WaitGroup is like a teacher waiting for all kids to finish.
// 4. Time delays show some jobs take longer but run together.
// 5. Main goroutine is the boss who waits for all kids to finish.

// Five key lessons for Chapter 27:
// 1. A race condition is when kids mess up a shared toy box by writing together.
// 2. WaitGroup is like a teacher waiting for all kids to finish their jobs.
// 3. Mutex is a lock letting one kid touch the toy box at a time.
// 4. Goroutines are kids working on tasks all at once to go faster.
// 5. A race detector (go run -race) finds problems when kids scribble without a lock.

// Five key lessons for Chapter 28:
// 1. Channels are like a toy pipe for kids to send and receive toys.
// 2. Sending and receiving toys through a channel must happen together.
// 3. Closing a channel tells kids no more toys are coming.
// 4. WaitGroup is like a teacher waiting for all kids to finish playing.
// 5. Checking if a channel is open shows if there are toys left.
