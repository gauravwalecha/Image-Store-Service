package main

import (
	"domain"
	"fmt"
	"methods"
	"net/http"
	"path"

	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
)

func main() {
	wsContainer := restful.NewContainer()
	RegisterCreateAlbum(wsContainer)
	RegisterDeleteAlbum(wsContainer)
	RegisterCreateImage(wsContainer)
	RegisterDeleteImage(wsContainer)
	RegisterGetAllImages(wsContainer)
	RegisterOpenAPI(wsContainer)
	RegisterSwaggerUI(wsContainer)
	fmt.Println("Starting the server on port 8080.....")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	fmt.Println(server.ListenAndServe())
}

// RegisterSwaggerUI - Register route for swagger
func RegisterSwaggerUI(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/swagger").To(swaggerUIPart1))
	container.Add(ws)
}

// swaggerUiPart1 - function for accessing static files
func swaggerUIPart1(req *restful.Request, resp *restful.Response) {
	http.ServeFile(
		resp.ResponseWriter,
		req.Request,
		path.Join("./src/resources/swagger-ui/", ""))
}

// RegisterOpenAPI - Registering routes for OpenApi
func RegisterOpenAPI(container *restful.Container) {
	config := restfulspec.Config{
		WebServices:                   container.RegisteredWebServices(),
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	container.Add(restfulspec.NewOpenAPIService(config))
}

//function for swagger details
func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "Image Store Service",
			Description: "Storing Images to Album",
		},
	}
	swo.Tags = []spec.Tag{spec.Tag{TagProps: spec.TagProps{
		Name:        "Image Store Service",
		Description: "backend for Image Store Service Application"}}}
}

// RegisterCreateAlbum -
func RegisterCreateAlbum(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/createAlbum")
	ws.Route(ws.POST("").
		Doc("Creates the Photo Album.").
		Returns(200, "ok", "success").
		Returns(501, "Internal Server Error", nil).
		Returns(405, "Method Not Allowed", nil).
		Produces(restful.MIME_JSON).
		To(methods.CreateAlbum))
	container.Add(ws)
}

// RegisterDeleteAlbum -
func RegisterDeleteAlbum(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/deleteAlbum")
	ws.Route(ws.DELETE("").
		Doc("Deletes the Photo Album.").
		Returns(200, "ok", "success").
		Returns(501, "Internal Server Error", nil).
		Returns(405, "Method Not Allowed", nil).
		Produces(restful.MIME_JSON).
		To(methods.DeleteAlbum))
	container.Add(ws)
}

// RegisterCreateImage -
func RegisterCreateImage(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/createImage")
	ws.Route(ws.POST("").
		Doc("Creates the Image into Photo Album.").
		Reads(domain.Image{}).
		Returns(200, "ok", "success").
		Returns(501, "Internal Server Error", nil).
		Returns(405, "Method Not Allowed", nil).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		To(methods.CreateImage))
	container.Add(ws)
}

// RegisterDeleteImage -
func RegisterDeleteImage(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/deleteImage")
	ws.Route(ws.DELETE("").
		Doc("Creates the Image in Photo Album.").
		Reads(domain.ImageName{}).
		Returns(200, "ok", "success").
		Returns(501, "Internal Server Error", nil).
		Returns(405, "Method Not Allowed", nil).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		To(methods.DeleteImage))
	container.Add(ws)
}

// RegisterGetAllImages -
func RegisterGetAllImages(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/getAllImages")
	ws.Route(ws.GET("").
		Doc("Creates the Image in Photo Album.").
		Returns(200, "ok", "success").
		Returns(501, "Internal Server Error", nil).
		Returns(405, "Method Not Allowed", nil).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON).
		To(methods.GetAllImages))
	container.Add(ws)
}
