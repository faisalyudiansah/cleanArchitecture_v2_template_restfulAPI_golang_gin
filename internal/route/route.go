package route

import (
	controllers "server/internal/controller"

	"github.com/gin-gonic/gin"
)

func ExampleControllerRoute(c *controllers.ExampleController, r *gin.Engine) {
	g := r.Group("")
	{
		g.GET("/example", c.Get)
	}
}
