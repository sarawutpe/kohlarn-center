package service

import (
	"main/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWT(c *gin.Context) {
	jwt, err := helper.GenerateJWT("6433079093b17af7e4bb8ad8")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "error": err.Error()})
	}

	data := map[string]string{"accessToken": jwt.AccessToken, "refreshToken": jwt.RefreshToken}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}

func RefreshToken(c *gin.Context) {
	refreshToken := c.Param("refresh-token")

	jwtData, err := helper.VerifyJWT(refreshToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "error": err.Error()})
		return
	}

	// Get Issuer and generate a new token
	issuer, ok := jwtData[helper.JwtIssuer].(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Issuer is not a string"})
		return
	}

	newJwt, err := helper.GenerateJWT(issuer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}

	// log.Println(jwtData["jwt_data"])

	data := map[string]string{"token": newJwt.AccessToken, "refreshToken": newJwt.RefreshToken}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}
