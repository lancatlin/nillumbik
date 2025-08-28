package author

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lancatlin/nillumbik/internal/db"
)

type Controller struct {
	q db.Querier
}

func NewController(queries db.Querier) Controller {
	return Controller{
		q: queries,
	}
}

func (u *Controller) ListAuthors(c *gin.Context) {
	authors, err := u.q.ListAuthors(c.Request.Context())
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, authors)
}

func (u *Controller) GetAuthorByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	author, err := u.q.GetAuthor(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(400, gin.H{"message": "Author not found"})
		return
	}

	c.JSON(200, author)
}
