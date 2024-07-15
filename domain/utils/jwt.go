package utils

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	issuer = "test"
	renewClaims = "ddc20ad0"
)

var jwtCfg JWT

type JWT struct {
	atSecretKey []byte        //Access Token Secret Key
	atd         time.Duration //Access Token Duration
	rtSecretKey []byte        //Refresh Token Secret Key
	rtd         time.Duration //Refresh Token Duration
}

type Claims struct {
	jwt.StandardClaims
	Session string `json:"session"`
	Renew   string `json:"renew,omitempty"`
}

func GenerateJWT(session string) (string, string, error) {
	//Create Access Token
	accessToken, err := generateAccessToken(session)
	if err != nil {
		return "", "", err
	}

	//Create Refresh Token
	refreshToken, err := generateRefreshToken(session)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func generateAccessToken(session string) (string, error) {
	accessClaims := Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: time.Now().UTC().Add(jwtCfg.atd).Unix(),
		},
		Session: session,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessSignedToken, err := accessToken.SignedString(jwtCfg.atSecretKey)
	if err != nil {
		return "", err
	}

	return accessSignedToken, nil
}

func generateRefreshToken(session string) (string, error) {
	refreshClaims := Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: time.Now().UTC().Add(jwtCfg.rtd).Unix(),
		},
		Session: session,
		Renew:   renewClaims,
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS384, refreshClaims)
	refreshSignedToken, err := refreshToken.SignedString(jwtCfg.rtSecretKey)
	if err != nil {
		return "", err
	}

	return refreshSignedToken, nil
}

func CheckAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return jwtCfg.atSecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid Token")
	}

	isr := fmt.Sprintf("%v", claims["iss"])
	if isr != issuer {
		return nil, fmt.Errorf("Invalid Issuer")
	}

	return claims, nil
}
