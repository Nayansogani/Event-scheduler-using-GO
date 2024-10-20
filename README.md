# Event Scheduler

This is a web-based event scheduler built using the [Gin web framework](https://gin-gonic.com/). The scheduler allows users to add and delete time-blocked events while ensuring no event overlaps with another. 

## Features

- **Add Events:** Users can schedule events by selecting a start and end time. The scheduler ensures that no events overlap.
- **Delete Events:** Users can delete any event from the schedule.
- **HTML UI:** The application displays a simple web interface for managing events.
- **Real-Time Feedback:** The UI shows success or failure animations based on the result of event additions or deletions.

## Getting Started

### Prerequisites

To run the application, you need:

- [Go](https://golang.org/doc/install) (version 1.16 or later)
- [Gin](https://github.com/gin-gonic/gin) web framework
- Basic knowledge of Go and HTTP server operations

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/event-scheduler.git
   cd event-scheduler
   
**2.Install the required dependencies:**
go mod init github.com/yourusername/event-scheduler
go get -u github.com/gin-gonic/gin

**3.Set up templates:**
Make sure your HTML files (such as index.html) are stored in a templates directory. 
You can create a simple HTML form that allows users to enter event times and view existing events.

**4.Running the Application**
go run main.go

