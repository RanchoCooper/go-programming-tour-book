package handle

import (
    "time"

    "github.com/dgrijalva/jwt-go"

    "go-programming-tour-book/blog-service/config"
    "go-programming-tour-book/blog-service/util"
)

/**
 * @author Rancho
 * @date 2021/12/7
 */

type Claims struct {
    jwt.StandardClaims
    AppKey    string `json:"app_key"`
    AppSecret string `json:"app_secret"`
}

func GetJWTSecret() []byte {
    return []byte(config.Config.JWT.Secret)
}

func GenerateToken(appKey, appSecret string) (string, error) {
    now := time.Now()
    expireTime := now.Add(time.Duration(config.Config.JWT.Expire) * time.Second)
    claims := Claims{
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expireTime.Unix(),
            Issuer:    config.Config.JWT.Issuer,
        },
        AppKey:    util.EncodeMD5(appKey),
        AppSecret: util.EncodeMD5(appSecret),
    }

    tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    token, err := tokenClaims.SignedString(GetJWTSecret())

    return token, err
}

func ParseToken(token string) (*Claims, error) {
    tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return GetJWTSecret(), nil
    })
    if err != nil {
        return nil, err
    }
    if tokenClaims != nil {
        claims, ok := tokenClaims.Claims.(*Claims)
        if ok && tokenClaims.Valid {
            return claims, nil
        }
    }

    return nil, err
}
