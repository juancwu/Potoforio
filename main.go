package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func main() {
	// Load connection string from .env file
	if os.Getenv("GO_ENV") != "production" {
		fmt.Println("Loading env...")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("failed to load env", err)
		}
	}

	templates, err := template.New("").ParseGlob("views/*.html")
	if err != nil {
		log.Fatalf("Error initializing templates: %v", err)
		os.Exit(1)
	}

	e := echo.New()
	e.Renderer = &TemplateRenderer{
		templates: templates,
	}

	e.Use(middleware.Logger())
	e.Static("/static", "static")

	e.GET("/", renderPage)

	e.GET("/service/health-check", func(c echo.Context) error {
		c.Response().Writer.WriteHeader(http.StatusOK)
		c.Response().Write([]byte("APP VERSION: 1.0.0"))
		return nil
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "5173"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
