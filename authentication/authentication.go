package authentication

import (
	"strings"
	"vinid_project/model"

	"github.com/gin-gonic/gin"

	jwt_lib "github.com/dgrijalva/jwt-go"
)

type AuthorizationHeader struct {
	Token string `header:"Authorization"`
}

func MakeJWT(user model.User) (string, error) {
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))

	token.Claims = jwt_lib.MapClaims{
		"userId": user.ID,
	}

	tokenString, err := token.SignedString([]byte(model.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseJWT(header AuthorizationHeader) (*jwt_lib.Token, error) {
	tokenFromHeader := strings.TrimSpace(strings.Replace(header.Token, "Bearer ", "", -1))
	claims := jwt_lib.MapClaims{}

	token, err := jwt_lib.ParseWithClaims(tokenFromHeader, claims, func(token *jwt_lib.Token) (interface{}, error) {
		return []byte(model.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil

}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var headerInfo AuthorizationHeader
		if err := c.ShouldBindHeader(&headerInfo); err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Header must include Authentication token!"})
		}

		token, err := ParseJWT(headerInfo)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token invalid : " + err.Error()})
		}

		if !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token invalid!"})
		}

		var userID int
		for k, v := range token.Claims.(jwt_lib.MapClaims) {
			if k == "userId" {
				userID = int(v.(float64))
			}
		}
		c.Set("userID", userID)
		c.Next()
	}
}
