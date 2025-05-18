package middleware

import (
	"fmt"
	"net/http"
	"snaptrail/internal/config"
	"snaptrail/internal/db"
	"snaptrail/internal/structs"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var (
	jwtSecret = config.Get().JwtSecret
	repo      *gorm.DB
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &structs.CustomClaims{}, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrInvalidKeyType
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			log.Err(err).Msgf("received an invalid jwt")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if claims, ok := token.Claims.(*structs.CustomClaims); ok && token.Valid {
			err = checkVersion(claims.Subject, claims.Version)
			if err != nil {
				log.Err(err).Msgf("received an invalid jwt")
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token version"})
				return
			}
		} else {
			log.Err(err).Msgf("received an invalid jwt")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Next()
	}
}

func checkVersion(userID string, version int64) error {
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return err
	}

	if repo == nil {
		repo = db.GetDb()
	}

	var user structs.User
	err = repo.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}

	if user.Version != version {
		return fmt.Errorf("jwt version does not match")
	}
	return nil
}
