package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const summarizerURL = "http://localhost:8000/summarize"

// Struct to define the request payload
type SummarizationRequest struct {
	Text string `json:"text"`
}

// Struct to define the response from the summarizer
type SummarizationResponse struct {
	Summary string `json:"summary"`
}

// Function to call the FastAPI backend for summarization
func summarizeText(text string) (string, error) {
	requestData := SummarizationRequest{Text: text}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	// Send POST request to the Python API
	resp, err := http.Post(summarizerURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read and process the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response SummarizationResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	return response.Summary, nil
}

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))

	// Route to handle summarization requests
	app.Post("/summarize", func(c *fiber.Ctx) error {

		// Parse the input from the request body
		request := new(SummarizationRequest)

		if err := c.BodyParser(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		// Call the Python API to get the summary
		summary, err := summarizeText(request.Text)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to summarize text"})
		}

		// Return the summarized text as a response
		return c.JSON(fiber.Map{"summary": summary})
	})

	// Start the Fiber web server
	log.Println("Server running on port 8080")
	log.Fatal(app.Listen(":8080"))
}
