package jwt

import (
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/spf13/cast"

    "blog-service/config"
    "blog-service/util"
)

/**
 * @author Rancho
 * @date 2022/1/7
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
    expire := now.Add(cast.ToDuration(config.Config.JWT.Expire))
    claims := Claims{
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expire.Unix(),
            IssuedAt:  now.Unix(),
            Issuer:    config.Config.JWT.Issuer,
            Id:        "",
            Subject:   "",
            Audience:  "",
            NotBefore: 0,
        },
        AppKey:    util.EncodeMD5(appKey),
        AppSecret: util.EncodeMD5(appSecret),
    }

    tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return tokenClaims.SignedString(GetJWTSecret())
}

func ParseToken(token string) (*Claims, error) {
    tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return GetJWTSecret(), nil
    })

    if tokenClaims != nil {
        claims, ok := tokenClaims.Claims.(*Claims)
        if ok && tokenClaims.Valid {
            return claims, nil
        }
    }

    return nil, err
}
