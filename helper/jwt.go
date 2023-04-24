package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// iss (Issuer) เป็น unique id ที่เอาไว้ระบุตัว client
// iat (Issued-at time) create time token
// exp (Expiration time) เป็นเวลาหมดอายุของ token
// 	Registered Claim Names  . . . . . . . . . . . . . . . . .   9
//    4.1.1.  "iss" (Issuer) Claim  . . . . . . . . . . . . . . . .   9
//    4.1.2.  "sub" (Subject) Claim . . . . . . . . . . . . . . . .   9
//    4.1.3.  "aud" (Audience) Claim  . . . . . . . . . . . . . . .   9
//    4.1.4.  "exp" (Expiration Time) Claim . . . . . . . . . . . .   9
//    4.1.5.  "nbf" (Not Before) Claim  . . . . . . . . . . . . . .  10
//    4.1.6.  "iat" (Issued At) Claim . . . . . . . . . . . . . . .  10
//    4.1.7.  "jti" (JWT ID) Claim  . . . . . . . . . . . . . . . .  10

// claims["username"] = "john.doe"

type TokenSigningJWT struct {
	AccessToken  string
	RefreshToken string
}

func GenerateJWT(iss string) (TokenSigningJWT, error) {
	jwtSecret := os.Getenv(EnvJWTSecret)
	timeNow := time.Now()

	aud := "go match ltd"

	// Create a new Token.
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = iss
	claims["iat"] = timeNow.Unix()
	claims["aud"] = aud
	claims["exp"] = timeNow.Add(time.Hour * 12).Unix()

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return TokenSigningJWT{
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}

	// Create a new Refresh Token
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["iss"] = iss
	refreshClaims["iat"] = timeNow.Unix()
	refreshClaims["aud"] = aud
	refreshClaims["exp"] = timeNow.Add(time.Hour * 13).Unix()

	// Sign and get the complete encoded refresh token as a string
	refreshTokenString, err := refreshToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return TokenSigningJWT{
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}

	// Return token, refresh token
	return TokenSigningJWT{
		AccessToken:  tokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

// func GenerateRefreshToken() (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)

// 	claims["iss"] = "6433079093b17af7e4bb8ad8"
// 	claims["iat"] = time.Now().Unix()
// 	claims["aud"] = "go match ltd"
// 	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

// 	// Sign and get the complete encoded token as a string
// 	jwtSecret := os.Getenv(JWT_SECRET)
// 	tokenString, err := token.SignedString([]byte(jwtSecret))
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil
// }

// Error types
// signature is invalid.
// Token is expired.
// invalid character '\x00' looking for beginning of value.
func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
	jwtSecret := os.Getenv(EnvJWTSecret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
