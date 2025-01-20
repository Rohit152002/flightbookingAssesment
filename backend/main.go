package main

import (
	"bytes"
	"flight/configuration"
	"flight/routes"
	"html/template"
	"log"

	"github.com/andrewcharlton/wkhtmltopdf-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func generatePDF(c *gin.Context) {
	// HTML content
	htmlContent := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Sample PDF</title>
        <style>
            body { font-family: Arial, sans-serif; margin: 20px; }
            h1 { color: #4CAF50; }
            p { font-size: 14px; }
        </style>
    </head>
    <body>
        <h1>Hello from Gin!</h1>
        <p>This PDF was generated dynamically using go-htmltopdf.</p>
    </body>
    </html>
    `

	// Generate the PDF
	tmpl := template.Must(template.New("page").Parse(htmlContent))
	buf := &bytes.Buffer{}
	tmpl.Execute(buf, c.Request.URL.String())
	doc := wkhtmltopdf.NewDocument()
	pg, err := wkhtmltopdf.NewPageReader(buf)
	if err != nil {
		log.Fatal("Error reading page buffer")
	}
	doc.AddPages(pg)

	// Set headers and serve the PDF
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", `attachment; filename="test.pdf"`)
	err = doc.Write(c.Writer)
	if err != nil {
		log.Fatal("Error serving pdf")
	}
}
func main() {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load env file")
	}
	httpServer := gin.Default()

	db := configuration.ConnectDB()

	httpServer.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200,
			gin.H{
				"working": "everything well",
			},
		)
	})

	httpServer.GET("/pdf", generatePDF)

	httpServer.Use(cors.Default())

	routes.FlightRoutes(httpServer, db)
	routes.PaymentRoutes(httpServer)
	routes.BookRoutes(httpServer, db)
	httpServer.Run(":8080")
}
