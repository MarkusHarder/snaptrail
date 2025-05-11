package appuser

import (
	"fmt"
	"regexp"
	"snaptrail/internal/config"
	"snaptrail/internal/structs"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = config.Get().JwtSecret

type service interface {
	login(user structs.User) (jwt string, err error)
	changePassword(pwChange structs.PasswordChange) (err error)
}

func newService() service {
	return svc{
		repo: newRepo(),
	}
}

type svc struct {
	repo repository
}

func (s svc) login(user structs.User) (jwt string, err error) {
	foundUser, err := s.repo.getUserByName(user.Username)
	if err != nil {
		log.Err(err).Msgf("could not find username in database")
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		log.Err(err).Msgf("invalid password")
		return "", fmt.Errorf("invalid username or password")
	}
	log.Info().Msgf("found user: %v", foundUser)

	return generateJWT(foundUser)
}

func (s svc) changePassword(pwChange structs.PasswordChange) (err error) {
	if !ValidatePassword(pwChange.NewPassowrd) {
		log.Err(err).Msgf("password does not meet requirements")
		return fmt.Errorf("password does not meet requirements")
	}
	foundUser, err := s.repo.getUserByName(pwChange.Username)
	if err != nil {
		log.Err(err).Msgf("could not find username in database")
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(pwChange.OldPassword))
	if err != nil {
		log.Err(err).Msgf("invalid password")
		return fmt.Errorf("invalid username or password")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(pwChange.NewPassowrd), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal().Msg("failed generating password hash")
	}

	newPwHash := string(hash[:])
	newVersion := foundUser.Version + 1

	return s.repo.changeUserPassword(foundUser.ID, newPwHash, newVersion)
}

func ValidatePassword(password string) bool {
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*]`).MatchString(password)

	return hasNumber && hasUpper && hasSpecial
}

func generateJWT(user structs.User) (string, error) {
	claims := structs.CustomClaims{
		Role:    user.Role,
		Version: user.Version,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatUint(user.ID, 10),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	log.Info().Msgf("got user for claims: %v", user)
	log.Info().Msgf("generated claims: %v", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
