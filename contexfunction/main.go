package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Function to handle JSON response
func handleJSONResponse(c *gin.Context) {

	// H is a shortcut for map[string]any
    data := gin.H{
        "message": "Hello, JSON!",
        "status":  "success",
    }
    c.JSON(http.StatusOK, data)
}

// Function to handle plain text response
func handleStringResponse(c *gin.Context) {
    c.String(http.StatusOK, "Hello, %s!", "World")
}

// Function to handle URL parameter
// localhost:8080/hello/saiumar

//which is use to get the paramter

func handleURLParam(c *gin.Context) {
    name := c.Param("saiumar")
    c.String(http.StatusOK, "Hello, %s!", name)
}

// Function to handle query parameters
// localhost:8080/welcome?firstName=John&lastName=Doe

//which is use to get the data from the url

func handleQueryParams(c *gin.Context) {
    firstName := c.Query("firstName")
    lastName := c.Query("lastName")
    message := "Welcome, " + firstName + " " + lastName + "!"
    c.String(http.StatusOK, message)
}

func main() {
    router := gin.Default()

    // Register routes and corresponding handler functions
    router.GET("/json", handleJSONResponse)
    router.GET("/string", handleStringResponse)
    router.GET("/hello/:name", handleURLParam)
    router.GET("/welcome", handleQueryParams)

    // Start the server
    router.Run(":8080")
}
