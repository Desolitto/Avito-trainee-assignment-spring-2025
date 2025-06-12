package auth

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
    "errors"
    "os"
    "log"
)

var jwtSecret []byte

func init() {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        log.Fatal("JWT_SECRET environment variable is not set")
    }
    jwtSecret = []byte(secret)
}

func GetJWTSecret() []byte {
    return jwtSecret
}

// Claims структура с ролями и стандартными полями JWT
type Claims struct {
    Role string `json:"role"`
    jwt.RegisteredClaims
}

// GenerateJWT создает JWT токен с указанной ролью
func GenerateJWT(role string) (string, error) {
    if role != "employee" && role != "moderator" {
        return "", errors.New("invalid role")
    }

    claims := &Claims{
        Role: role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // срок жизни 24 часа
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            Issuer:    "backend-service",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(GetJWTSecret())
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

func ParseJWT(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return GetJWTSecret(), nil
    })
    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}
