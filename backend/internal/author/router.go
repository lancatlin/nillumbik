package author

import "github.com/gin-gonic/gin"

func Register(r gin.IRouter, ctl *Controller) {
	g := r.Group("/authors")
	g.GET("", ctl.ListAuthors)
	g.GET("/:id", ctl.GetAuthorByID)
}
