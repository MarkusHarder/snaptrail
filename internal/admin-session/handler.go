package adminsession

import (
	"context"
	"net/http"
	"snaptrail/internal/structs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type sessionForm struct {
	SessionName string `form:"sessionName" binding:"required"`
	Subtitle    string `form:"subtitle" binding:"required"`
	Description string `form:"description" binding:"required"`
	Published   bool   `form:"published"` // binding is not required because gin cannot differentiate between zero value and false
	Date        string `form:"date" binding:"required"`
}

type Handler struct {
	svc Service
}

func New() Handler { return Handler{svc: NewService(context.Background())} }

func (h *Handler) Session(c *gin.Context) {
	session, err := h.svc.GetSessions(false)
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch all sessions")
		c.String(http.StatusInternalServerError, "failed to fetch all sessions")
		return
	}

	c.JSON(http.StatusOK, session)
}

func (h *Handler) CreateOrUpdateSession(c *gin.Context) {
	id := c.Param("id")
	update := true
	log.Info().Msgf("creating or updating session, got id: %s, will update: %t", id, update)
	if id == "" {
		update = false
		log.Warn().Msg("failed to parse session id update, creating session instead")
	}

	var form sessionForm
	if err := c.ShouldBind(&form); err != nil {
		log.Err(err).Msg("failed to parse session from form")
		c.String(http.StatusBadRequest, "failed parse session from form")
		return
	}

	uploadedThumbnail, err := c.FormFile("uploadedThumbnail")
	if err != nil {
		log.Err(err).Msg("failed to parse thumbnail from form")
		c.String(http.StatusBadRequest, "failed to parse thumbnail from form")
		return
	}

	sessionDate, err := time.Parse(time.RFC3339, form.Date)
	if err != nil {
		log.Err(err).Msg("failed to parse session date from form")
		c.String(http.StatusBadRequest, "failed to parse session date from form")
		return
	}

	session := structs.Session{
		ID:          id,
		Name:        form.SessionName,
		Subtitle:    form.Subtitle,
		Date:        sessionDate,
		Description: form.Description,
		Published:   form.Published,
	}
	session.ID = id
	if update {
		if err := h.svc.updateSession(&session, uploadedThumbnail); err != nil {
			log.Err(err).Msg("failed to create or update session")
			c.String(http.StatusInternalServerError, "failed to create or update session")
			return
		}
	} else {
		if err := h.svc.createSession(&session, uploadedThumbnail); err != nil {
			log.Err(err).Msg("failed to create or update session")
			c.String(http.StatusInternalServerError, "failed to create or update session")
			return
		}
	}
	c.JSON(http.StatusOK, session)
}

func (h *Handler) DeleteSession(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		log.Error().Msg("failed to parse session id to delete")
		c.JSON(http.StatusBadRequest, "failed to parse session id to delete")
		return
	}

	if err := h.svc.deleteSession(id); err != nil {
		log.Err(err).Msg("failed to delete session")
		c.String(http.StatusInternalServerError, "failed to delete session")
		return
	}

	c.JSON(http.StatusOK, id)
}
