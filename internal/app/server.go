package app

import (
	"fmt"
	"log"

	"github.com/certified-juniors/AtomHackEarthBackend/docs"
	"github.com/certified-juniors/AtomHackEarthBackend/internal/app/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Run запускает приложение
func (app *Application) Run() {
	r := gin.Default()

	docs.SwaggerInfo.Title = "AtomHackEarthBackend RestAPI"
	docs.SwaggerInfo.Description = "API server for Earth application"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8081"
	docs.SwaggerInfo.BasePath = "/api/v1"

	r.Use(middleware.CorsMiddleware())

	ApiGroup := r.Group("/api/v1")
	{
		DocumentGroup := ApiGroup.Group("/document")
		{	
			DocumentGroup.GET("/formed", app.handler.GetFormedDocuments)
			DocumentGroup.GET("/:docID", app.handler.GetDocumentByID)
			DocumentGroup.POST("/", app.handler.AcceptDocument)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	addr := fmt.Sprintf("%s:%d", app.cfg.API.ServiceHost, app.cfg.API.ServicePort)
	r.Run(addr)

	log.Println("Server down")
}
