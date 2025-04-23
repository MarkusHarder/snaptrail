package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	svc service
}

func New() Handler { return Handler{svc: newService()} }

func (h *Handler) Hello(c *gin.Context) {
	hello, err := h.svc.getHello()

	if err != nil {
		log.Error().Err(err).Msg("failed to fetch all articles")
		c.String(http.StatusInternalServerError, "failed to fetch all articles")
		return
	}

	c.JSON(http.StatusOK, hello)

}
