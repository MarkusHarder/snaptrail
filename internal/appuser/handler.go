package appuser

import (
	"net/http"
	"snaptrail/internal/structs"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type jwtToken struct {
	Token string `json:"token"`
}

type Handler struct {
	svc service
}

func New() Handler { return Handler{svc: newService()} }

func (h *Handler) Login(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	jwt, err := h.svc.login(structs.User{Username: username, Password: password})
	if err != nil {
		log.Err(err).Msg("failed to login user")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, jwtToken{Token: jwt})
}

func (h *Handler) PasswordChange(c *gin.Context) {
	var pwChange structs.PasswordChange

	err := c.Bind(&pwChange)
	if err != nil {
		log.Err(err).Msg("failed to parse PasswordChange")
		c.String(http.StatusBadRequest, "failed to parse PasswordChange")
		return
	}
	log.Info().Msgf("Got PW Change: %v", pwChange)

	err = h.svc.changePassword(pwChange)
	if err != nil {
		log.Err(err).Msg("failed to update password")
		c.String(http.StatusBadRequest, "failed to update password")
		return
	}

	c.Status(http.StatusOK)
}
