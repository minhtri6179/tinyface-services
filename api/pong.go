package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})

}
