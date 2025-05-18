package thumbnail

import (
	"net/http"
	"snaptrail/internal/structs"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	svc service
}

func New() Handler { return Handler{svc: newService()} }

func (h *Handler) ThumbnailById(c *gin.Context) {
	image, err := h.getThumbnailById(c, false)
	if err != nil {
		log.Err(err).Msg("failed to fetch thumbnail from db")
		c.String(http.StatusInternalServerError, "failed to fetch thumbnail from db")
		return
	}

	c.Data(http.StatusOK, image.MimeType, image.Data)
}

func (h *Handler) PublicThumbnailById(c *gin.Context) {
	image, err := h.getThumbnailById(c, true)
	if err != nil {
		log.Err(err).Msg("failed to fetch thumbnail from db")
		c.String(http.StatusInternalServerError, "failed to fetch thumbnail from db")
		return
	}

	c.Data(http.StatusOK, image.MimeType, image.Data)
}

func (h *Handler) getThumbnailById(c *gin.Context, public bool) (*structs.Thumbnail, error) {
	thumbnailIdParam := c.Param("thumbnailID")
	sessionIdParam := c.Param("sessionID")
	thumbnailId, err := strconv.ParseUint(thumbnailIdParam, 10, 0)
	if err != nil {
		log.Err(err).Msg("failed parsing thumbnail id")
		return nil, err
	}
	sessionId, err := strconv.ParseUint(sessionIdParam, 10, 0)
	if err != nil {
		log.Err(err).Msg("failed parsing sesison id")
		return nil, err
	}

	image, err := h.svc.getThumbnailById(sessionId, thumbnailId, public)
	if err != nil {
		return nil, err
	}

	return &image, nil
}
