package auth

import (
	"log"
	"time"

	"github.com/Ucell/client_manager/configuration"
	"github.com/dgrijalva/jwt-go"
)

func GenerateJWTToken(id, role string) (string, error) {
	conf, err := configuration.Load()
	if err != nil {
		log.Fatalf("Konfiguratsiya yuklashda xatolik: %v", err)
	}
	token := *jwt.New(jwt.SigningMethodHS256)
	//payload
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	claims["role"] = role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(30 * time.Minute).Unix()

	newToken, err := token.SignedString([]byte(conf.Token.TOKEN))
	if err != nil {
		return "", err
	}

	return newToken, nil
}

func ValidateToken(tokenStr string) (bool, error) {
	_, err := ExtractClaim(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaim(tokenStr string) (*jwt.MapClaims, error) {
	conf, err := configuration.Load()
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(conf.Token.TOKEN), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, err
	}

	return &claims, nil
}

func GetUserIdFromToken(req string) (Id string, Role string, err error) {
	conf, err := configuration.Load()
	Token, err := jwt.Parse(req, func(token *jwt.Token) (interface{}, error) { return []byte(conf.Token.TOKEN), nil })
	if err != nil || !Token.Valid {
		return "", "", err
	}
	claims, ok := Token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", err
	}
	Id = claims["user_id"].(string)
	Role = claims["role"].(string)

	return Id, Role, nil
}
