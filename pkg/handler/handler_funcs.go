package handler

import "github.com/gin-gonic/gin"

func (h *Handler) function(c *gin.Context) {
	c.Writer.Write([]byte("The router is working!"))
}
