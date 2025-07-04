package main

import (
    "encoding/json"
    "fmt"
    "io"
    "math/rand"
    "net/http"
    "strconv"
    "time"

    "github.com/gorilla/mux"
)

// A Course toy box holds course info
type Course struct {
    CourseId    string  `json:"courseid"`   // Course ID, like "42"
    CourseName  string  `json:"coursename"` // Course name, like "Go API"
    CoursePrice int     `json:"price"`      // Price, like 299
    Author      *Author `json:"author"`     // Author info
}

// An Author toy box holds author info
type Author struct {
    Fullname string `json:"fullname"` // Author name, like "Hitesh Choudhary"
    Website  string `json:"website"`  // Website, like "FruitLearnOnline.in"
}

// Fake database: a shelf of course toy boxes
var courses []Course

// The main function is like the start button for our program!
func main() {
    // Say hello to our API adventure
    fmt.Println("Welcome to our API toy shop!")

    // Add some starter courses to our shelf
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

    // Set up shop stalls (routes) with gorilla/mux
    router := mux.NewRouter()
    router.HandleFunc("/", serveHome).Methods("GET")              // Home page
    router.HandleFunc("/courses", getAllCourses).Methods("GET")    // Get all courses
    router.HandleFunc("/course/{id}", getOneCourse).Methods("GET") // Get one course by ID
    router.HandleFunc("/course", createCourse).Methods("POST")     // Create a course
    router.HandleFunc("/course/{id}", updateCourse).Methods("PUT") // Update a course

    // Start the shop at port 8000
    fmt.Println("Starting API server at http://localhost:8000")
    err := http.ListenAndServe(":8000", router)
    if err != nil {
        panic(err)
    }
}

// Helper for the home page, like a welcome sign
func serveHome(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
}

// Helper to get all courses, like showing all toys
func getAllCourses(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Get all courses")
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(courses)
}

// Helper to get one course by ID, like picking a toy
func getOneCourse(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Get one course")
    w.Header().Set("Content-Type", "application/json")

    // Grab ID from the visitor's note
    params := mux.Vars(r)
    courseId := params["id"]

    // Look for the course toy on the shelf
    for _, course := range courses {
        if course.CourseId == courseId {
            json.NewEncoder(w).Encode(course)
            return
        }
    }

    // If no toy found, send a sad note
    response := map[string]string{"error": fmt.Sprintf("No Course found with ID %s", courseId)}
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(response)
}

// Helper to create a course, like adding a new toy
func createCourse(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Create a course")
    w.Header().Set("Content-Type", "application/json")

    // Check if it's a POST request
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST allowed!", http.StatusMethodNotAllowed)
        return
    }

    // Read the toy box (visitor's note)
    body, err := io.ReadAll(r.Body)
    if err != nil {
        response := map[string]string{"error": "Can't read toy box!"}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }
    defer r.Body.Close()

    // Turn JSON words into a course toy
    var course Course
    if err := json.Unmarshal(body, &course); err != nil {
        response := map[string]string{"error": "No data inside JSON or bad JSON toy!"}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }

    // Check if we got the right toys
    if course.CourseName == "" || course.Author == nil || course.Author.Fullname == "" {
        response := map[string]string{"error": "Need course name and author fullname!"}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(response)
        return
    }

    // Give the course a random ID, like picking a lucky number
    rand.Seed(time.Now().UnixNano())
    course.CourseId = strconv.Itoa(rand.Intn(100))

    // Add the course to our shelf
    courses = append(courses, course)

    // Send back the new course
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(course)
}

// Helper to update a course, like changing a toy
func updateCourse(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Update a course")
    w.Header().Set("Content-Type", "application/json")

    // Check if it's a PUT request
    if r.Method != http.MethodPut {
        http.Error(w, "Only PUT allowed!", http.StatusMethodNotAllowed)
        return
    }

    // Grab ID from the visitor's note
    params := mux.Vars(r)
    courseId := params["id"]

    // Look for the course toy on the shelf
    for index, course := range courses {
        if course.CourseId == courseId {
            // Remove the old course
            courses = append(courses[:index], courses[index+1:]...)

            // Read the new toy box (visitor's note)
            var updatedCourse Course
            if err := json.NewDecoder(r.Body).Decode(&updatedCourse); err != nil {
                response := map[string]string{"error": "No data inside JSON or bad JSON toy!"}
                w.WriteHeader(http.StatusBadRequest)
                json.NewEncoder(w).Encode(response)
                return
            }
            defer r.Body.Close()

            // Check if we got the right toys
            if updatedCourse.CourseName == "" || updatedCourse.Author == nil || updatedCourse.Author.Fullname == "" {
                response := map[string]string{"error": "Need course name and author fullname!"}
                w.WriteHeader(http.StatusBadRequest)
                json.NewEncoder(w).Encode(response)
                return
            }

            // Keep the same ID
            updatedCourse.CourseId = courseId

            // Add the updated course back
            courses = append(courses, updatedCourse)

            // Send back the updated course
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode(updatedCourse)
            return
        }
    }

    // If no toy found, send a sad note
    response := map[string]string{"error": fmt.Sprintf("No Course found with ID %s", courseId)}
    w.WriteHeader(http.StatusNotFound)
    json.NewEncoder(w).Encode(response)
}

// Explanation comments for a 5-year-old:
// - What's an API? It's like a toy shop where friends can ask for toys (courses) or send new or changed toys using special letters (requests)!
// - What's a module? It's like a big toy box with a name (github.com/hiteshchoudhary/buildapi) that holds your toys and friends (gorilla/mux).
// - What's 'go mod init github.com/hiteshchoudhary/buildapi'? It's like making a new toy box with a name and a go.mod file.
// - What's 'go.sum'? It's like a list of toy receipts to make sure friend toys (like gorilla/mux) are safe.
// - How is go.sum created? Hitesh ran:
//   1. 'go mod init github.com/hiteshchoudhary/buildapi' to make go.mod.
//   2. 'go get github.com/gorilla/mux@v1.8.0' to add the router toy, creating go.sum.
//   3. 'go mod tidy' to clean up and update go.sum.
// - What's 'github.com/gorilla/mux@v1.8.0'? It's like a super helper toy for making shop stalls, with version 1.8.0.
// - What's 'type Course struct'? It's like a toy box for course toys, with places for ID, name, price, and an author toy.
// - What's 'type Author struct'? It's like a smaller toy box for author toys, with name and website.
// - What's '*Author'? It's like a map pointing to an author toy box, so we can share it.
// - What's 'var courses []Course'? It's like a shelf to hold course toy boxes, like a fake shop storage.
// - What's 'router := mux.NewRouter()'? It's like building a shop with stalls for visitors.
// - What's 'router.HandleFunc("/course/{id}", getOneCourse)'? It's like a stall where a friend asks for one course toy by its ID.
// - What's 'mux.Vars(r)'? It's like checking a visitor's note for the course ID, like "42".
// - What's 'getOneCourse(w, r)'? It's like finding a course toy by ID and giving it to the visitor.
// - What's 'createCourse(w, r)'? It's like taking a new course toy from a visitor and putting it on the shelf.
// - What's 'updateCourse(w, r)'? It's like finding a course toy, throwing it out, and putting a new one with the same ID.
// - What's 'courses = append(courses[:index], courses[index+1:]...)'? It's like taking a toy off the shelf by cutting it out.
// - What's 'json.NewDecoder(r.Body).Decode(&updatedCourse)'? It's like turning a visitor's JSON words into a new course toy.
// - What's 'rand.Seed(time.Now().UnixNano())'? It's like mixing a bag of number toys to pick a random one each time.
// - What's 'strconv.Itoa(rand.Intn(100))'? It's like picking a random number (0-99) and turning it into a word (like "42") for the course ID.
// - What's 'json.NewEncoder(w).Encode(course)'? It's like sending a course toy as JSON words to the visitor.
// - What's 'response := map[string]string{...}'? It's like sending a sad note if we can't find or use the toy.
// - What's 'w.WriteHeader(http.StatusCreated)'? It's like saying "New toy added!" (201) for create.
// - What's 'w.WriteHeader(http.StatusNotFound)'? It's like saying "Toy not here!" (404) for missing toys.
// - How to set up? 
//   1. Make folder '24buildapi'.
//   2. Run 'go mod init github.com/hiteshchoudhary/buildapi' to create go.mod.
//   3. Run 'go get github.com/gorilla/mux@v1.8.0' to add the router toy and create go.sum.
//   4. Save this main.go file.
//   5. Run 'go mod tidy' to clean go.mod and update go.sum.
//   6. Run 'go run main.go' to start the server.
// - How to test with ThunderClient? In VS Code, get ThunderClient:
//   1. GET Home: New Request > http://localhost:8000 > Send. See "<h1>Welcome to API by LearnCodeOnline</h1>".
//   2. GET All Courses: New Request > http://localhost:8000/courses > Send. See JSON like [{"courseid":"1","coursename":"Go API Bootcamp",...}].
//   3. GET One Course: New Request > http://localhost:8000/course/1 > Send. See JSON like {"courseid":"1","coursename":"Go API Bootcamp",...}. Try http://localhost:8000/course/99 for {"error":"No Course found with ID 99"}.
//   4. POST Course: New Request > http://localhost:8000/course > POST > Body > JSON > {"coursename":"Python Bootcamp","price":149,"author":{"fullname":"John Doe","website":"FruitLearnOnline.in"}} > Send. See JSON with new course and random ID.
//   5. PUT Update Course: New Request > http://localhost:8000/course/1 > PUT > Body > JSON > {"coursename":"Advanced Go","price":399,"author":{"fullname":"John Doe","website":"FruitLearnOnline.in"}} > Send. See JSON like {"courseid":"1","coursename":"Advanced Go",...}. Try http://localhost:8000/course/99 for {"error":"No Course found with ID 99"}.
// - Why fruit theme? It keeps the fun theme (like FruitLearnOnline.in) from earlier, like a fruit school shop.
// - How does this help beginners? It shows how to make a toy shop (API) where friends can get, add, or change course toys with random IDs, using a cool router (gorilla/mux) and JSON, like a real web shop in Go.