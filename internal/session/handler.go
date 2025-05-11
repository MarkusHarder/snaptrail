package session

import (
	"context"
	"net/http"
	adminsession "snaptrail/internal/admin-session"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	svc adminsession.Service
}

func New() Handler { return Handler{svc: adminsession.NewService(context.Background())} }

func (h *Handler) Session(c *gin.Context) {
	session, err := h.svc.GetSessions(true)
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch all sessions")
		c.String(http.StatusInternalServerError, "failed to fetch all sessions")
		return
	}

	c.JSON(http.StatusOK, session)
}
