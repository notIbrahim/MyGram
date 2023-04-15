package router

import (
	"MyGram/app/middleware"
	"MyGram/app/queries"

	"github.com/gin-gonic/gin"
	SwaggersFile "github.com/swaggo/files"
	Swaggers "github.com/swaggo/gin-swagger"
)

func ServerReady() *gin.Engine {
	route := gin.Default()

	// Users Route
	RouteUser := route.Group("/users")
	{
		RouteUser.POST("/register", queries.Registered)
		RouteUser.POST("/login", queries.UserLogged)
	}

	route.GET("/docs/*any", Swaggers.WrapHandler(SwaggersFile.Handler))

	//Route Photo
	RoutePhoto := route.Group("/photos")
	{
		RoutePhoto.GET("/", queries.GetAll)
		RoutePhoto.Use(middleware.Auth())
		RoutePhoto.GET("/:ID", queries.GetOne)
		RoutePhoto.POST("/Add", queries.CreatePhoto)
		RoutePhoto.PUT("/Edit/:ID", middleware.Authorize("photos"), queries.UpdatePhoto)
		RoutePhoto.DELETE("/Delete/:ID", middleware.Authorize("photos"), queries.DeletePhoto)
	}

	// Comment Route
	RouteComment := route.Group("/comments")
	{
		RouteComment.GET("/", queries.GetAllComments)
		RouteComment.Use(middleware.Auth())
		RouteComment.GET("/:ID", queries.GetOneComment)
		RouteComment.POST("/Add", queries.CreateComment)
		RouteComment.PUT("/Edit/:ID", middleware.Authorize("comments"), queries.UpdateComment)
		RouteComment.DELETE("/Delete/:ID", middleware.Authorize("comments"), queries.DeleteComment)
		// TODO
	}

	// // Social Route
	RouteSocial := route.Group("/social")
	{
		RouteSocial.GET("/", queries.GetStatusAll)
		RouteSocial.GET("/:ID", queries.GetStatus)
		RouteSocial.Use(middleware.Auth())
		RouteSocial.POST("/Add", queries.CreateSocialMedia)
		RouteSocial.PUT("/Edit/:ID", middleware.Authorize("socials"), queries.UpdateSocialMedia)
		RouteSocial.DELETE("/Delete/:ID", middleware.Authorize("socials"), queries.DeleteSocialMedia)
		// TODO
	}

	return route
}
