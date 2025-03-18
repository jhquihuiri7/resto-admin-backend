package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"resto-admin-backend/internal/auth"
	"strings"
)

// FirebaseAuthMiddleware validates the Firebase authentication token
func FirebaseAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Extract token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// The token should be in the format: "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		idToken := tokenParts[1]

		// Verify the token
		token, err := auth.VerifyToken(idToken)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Firebase token"})
			c.Abort() // Stop further processing
			return
		}
		// Store the token claims in the context for further use
		c.Set("firebaseUser", token)
		c.Next()
	}
}
