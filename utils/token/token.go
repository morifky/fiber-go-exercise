package token

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	Secret string
}

func New(s string) *Token {
	return &Token{
		Secret: s,
	}
}

func ExtractTokenFromHeader(c *fiber.Ctx) string {
	bearer := c.Get("Authorization")

	token := strings.Split(bearer, " ")

	if len(token) == 2 {
		return token[1]
	}

	return ""
}

func (t *Token) CreateToken(uid uint32) (string, error) {
	//create jwt claim maps
	claims := jwt.MapClaims{}
	claims["user_id"] = uid
	claims["exp"] = time.Now().Add(time.Hour + 1) //exp after 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(t.Secret))
}

func (t *Token) parseJWT(str string) (*jwt.Token, error) {
	token, err := jwt.Parse(str, func(j *jwt.Token) (interface{}, error) {
		if _, ok := j.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", j.Header["alg"])
		}
		return []byte(t.Secret), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func (t *Token) ValidateToken(str string) (bool, error) {
	_, err := jwt.Parse(str, func(jt *jwt.Token) (interface{}, error) {
		if _, ok := jt.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jt.Header["alg"])
		}
		return []byte(t.Secret), nil
	})

	if err != nil {
		return false, err
	}
	return true, nil
}

func (t *Token) ExtractTokenID(str string) (uint32, error) {
	token, err := t.parseJWT(str)
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}
	return 0, nil
}
