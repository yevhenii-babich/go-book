package main

import "github.com/gin-gonic/gin"

func serveDifferentEncoding(c *gin.Context) {
	data := gin.H{
		"message": "Hello, World!",
	}
	switch c.GetHeader("Accept") {
	case "application/xml":
		c.XML(200, data)
	case "application/x-yaml":
		c.YAML(200, data)
	case "application/x-toml":
		c.TOML(200, data)
	case "text/plain":
		c.String(200, data["message"].(string))
	case "application/json":
		fallthrough
	default:
		c.JSON(200, data)
	}
}

func main() {
	router := gin.Default()
	router.GET("/hello", serveDifferentEncoding)
	router.Run("127.0.0.1:8080")
}
