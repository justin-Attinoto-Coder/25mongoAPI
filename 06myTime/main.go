package main

import (
    "fmt"
    "time"
)

// The main function is like the start button for our program!
func main() {
    // Say hello to our time adventure
    fmt.Println("Welcome to time study of golang")

    // Get the current time, like looking at a clock right now
    presentTime := time.Now()
    fmt.Println("Right now, the time is:", presentTime)
    // This shows the full time, like: 2025-06-18 03:38:00.123456789 +0000 UTC

    // Format the time to make it look pretty, like choosing a clock style
    // "01-02-2006 15:04:05 Monday" means: month-day-year hour:minute:second day
    formattedTime := presentTime.Format("01-02-2006 15:04:05 Monday")
    fmt.Println("Pretty time is:", formattedTime)
    // This looks like: 06-18-2025 03:38:00 Wednesday

    // Make a special date, like marking a birthday on a calendar
    createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
    fmt.Println("Special date is:", createdDate)
    // This shows: 2020-08-10 23:23:00 +0000 UTC

    // Format the special date to show just month, day, year, and day name
    formattedDate := createdDate.Format("01-02-2006 Monday")
    fmt.Println("Pretty special date is:", formattedDate)
    // This looks like: 08-10-2020 Monday

    // Add some time, like fast-forwarding a clock
    futureTime := presentTime.Add(2 * time.Hour)
    fmt.Println("Two hours from now:", futureTime)
    // This shows the time 2 hours later, like: 2025-06-18 05:38:00.123456789 +0000 UTC

    // Find the difference between two times, like counting hours between events
    timeDifference := futureTime.Sub(presentTime)
    fmt.Println("Time difference is:", timeDifference.Hours(), "hours")
    // This shows: Time difference is: 2 hours

    // Check if one time is before another, like asking if lunch is before dinner
    isBefore := createdDate.Before(presentTime)
    fmt.Println("Is special date before now?", isBefore)
    // This shows: true, because 2020 is before 2025

    // Format time in different styles to show more options
    shortFormat := presentTime.Format("2006-01-02") // Just year-month-day
    fmt.Println("Short date format:", shortFormat)
    // This looks like: 2025-06-18

    timeOnly := presentTime.Format("15:04") // Just hour:minute
    fmt.Println("Time only format:", timeOnly)
    // This looks like: 03:38
}