package main

import (
	"log" // Import the log package to use logging
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Event represents a time block with a start and end time
type Event struct {
	StartTime int
	EndTime   int
}

// Scheduler holds a list of events
type Scheduler struct {
	events []Event
}

// AddEvent adds a new event if it doesn't overlap with existing ones
func (s *Scheduler) AddEvent(startTime, endTime int) bool {
	for _, event := range s.events {
		if startTime < event.EndTime && endTime > event.StartTime {
			log.Printf("Event overlap detected: existing event %d:00 to %d:00 conflicts with new event %d:00 to %d:00\n", event.StartTime, event.EndTime, startTime, endTime)
			return false
		}
	}
	log.Printf("Adding new event from %d:00 to %d:00\n", startTime, endTime)
	s.events = append(s.events, Event{StartTime: startTime, EndTime: endTime})
	return true
}

// DeleteEvent deletes an event by index
func (s *Scheduler) DeleteEvent(index int) {
	if index >= 0 && index < len(s.events) {
		log.Printf("Deleting event: %d:00 to %d:00\n", s.events[index].StartTime, s.events[index].EndTime)
		s.events = append(s.events[:index], s.events[index+1:]...)
	} else {
		log.Printf("Invalid index for deletion: %d\n", index)
	}
}

// GetEvents returns the list of all scheduled events
func (s *Scheduler) GetEvents() []Event {
	log.Printf("Fetching all scheduled events: %d event(s) found\n", len(s.events))
	return s.events
}

// Global scheduler
var scheduler = Scheduler{}

func main() {
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Show the main page
	router.GET("/", func(c *gin.Context) {
		log.Println("Rendering main page with scheduled events")
		events := scheduler.GetEvents()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":  "Event Scheduler",
			"events": events,
		})
	})

	// Handle form submission to add events
	router.POST("/addevent", func(c *gin.Context) {
		startTimeStr := c.PostForm("start_time")
		endTimeStr := c.PostForm("end_time")
		log.Printf("Received request to add event: Start Time: %s, End Time: %s\n", startTimeStr, endTimeStr)

		// Convert start_time and end_time to integers
		startTime, err1 := strconv.Atoi(startTimeStr)
		endTime, err2 := strconv.Atoi(endTimeStr)

		if err1 != nil || err2 != nil || startTime >= endTime || startTime < 0 || endTime > 23 {
			log.Printf("Invalid input: Start Time: %s, End Time: %s\n", startTimeStr, endTimeStr)
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"title":     "Event Scheduler",
				"error":     "Invalid input! Make sure times are numbers between 0-23, and start time is before end time.",
				"events":    scheduler.GetEvents(),
				"animation": "failure", // Pass a failure animation flag
			})
			return
		}

		// Try adding the event to the scheduler
		if scheduler.AddEvent(startTime, endTime) {
			log.Println("Event added successfully")
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title":     "Event Scheduler",
				"events":    scheduler.GetEvents(),
				"animation": "success", // Pass a success animation flag
			})
		} else {
			log.Println("Failed to add event due to overlap")
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"title":     "Event Scheduler",
				"error":     "Event overlaps with an existing event.",
				"events":    scheduler.GetEvents(),
				"animation": "failure", // Pass a failure animation flag
			})
		}
	})

	// Handle event deletion
	router.POST("/deleteevent/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		log.Printf("Received request to delete event with ID: %s\n", idStr)
		id, err := strconv.Atoi(idStr)
		if err == nil && id >= 0 && id < len(scheduler.GetEvents()) {
			scheduler.DeleteEvent(id)
			log.Println("Event deleted successfully")
		} else {
			log.Printf("Invalid ID for deletion: %s\n", idStr)
		}

		// Redirect to the homepage after deletion
		c.Redirect(http.StatusFound, "/")
	})

	// Run the web server
	log.Println("Starting server on port 8080")
	router.Run(":8080") // Runs on http://localhost:8080
}
