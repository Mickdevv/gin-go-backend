package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gin-gonic/gin"
)

func ClerkAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid authorization header"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwt.Verify(c.Request.Context(), &jwt.VerifyParams{Token: token, Leeway: 5 * time.Second})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		prettyClaims, err := json.MarshalIndent(claims, "", "  ")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal claims"})
			return
		}

		fmt.Println("Decoded JWT claims:")
		fmt.Println(string(prettyClaims))

		usr, err := user.Get(c.Request.Context(), claims.Subject)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to fetch user"})
		}

		prettyUser, err := json.MarshalIndent(usr, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling claims:", err)
			return
		}

		fmt.Println("Decoded clerk user:")
		fmt.Println(string(prettyUser))

		c.Set("clerkUser", usr)

		c.Next()
	}
}
