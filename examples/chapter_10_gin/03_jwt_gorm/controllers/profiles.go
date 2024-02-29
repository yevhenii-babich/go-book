package controllers

import (
	"github.com/gin-gonic/gin"
	"jwt_gorm/initializers"
	"jwt_gorm/models"
	"net/http"
	"strconv"
)

type profile struct {
	FullName string `json:"full_name" binding:"required"`
	Age      int    `json:"age"`
}

func CreateProfile(c *gin.Context) {
	// Get the user id from the token
	tmp, ok := c.Get("user")
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	user := tmp.(models.User)
	// Get the full name and age off req body
	var body profile
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Create the profile
	profile := models.Profile{UserID: user.ID, FullName: body.FullName, Age: body.Age}
	result := initializers.DB.Create(&profile)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Failed to create profile.",
			"database": result.Error,
		})
		return
	}
	// Respond
	c.JSON(http.StatusOK, profile)
}

func GetProfile(c *gin.Context) {
	// Get the user id from the token
	tmp, ok := c.Get("user")
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	user := tmp.(models.User)
	// Find the profile
	var profile models.Profile
	result := initializers.DB.Where("user_id = ?", user.ID).First(&profile)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Profile not found.",
		})
		return
	}
	// Respond
	c.JSON(http.StatusOK, profile)
}

func UpdateProfile(c *gin.Context) {
	// Get the user id from the token
	tmp, ok := c.Get("user")
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	user := tmp.(models.User)
	// Get the full name and age off req body
	var body profile
	if c.ShouldBindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	// Find the profile
	var profile models.Profile
	result := initializers.DB.Where("user_id = ?", user.ID).First(&profile)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Profile not found.",
		})
		return
	}
	// Update the profile
	profile.FullName = body.FullName
	profile.Age = body.Age
	result = initializers.DB.Save(&profile)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update profile.",
		})
		return
	}
	// Respond
	c.JSON(http.StatusOK, profile)
}

func GetProfileByID(c *gin.Context) {
	// Get the profile id from the req params
	id := c.Param("id")
	if id != "" {
		singleProfile(c, id)
	}
	c.AbortWithStatus(http.StatusBadRequest)
}

func AllProfiles(c *gin.Context) {
	var profiles []models.Profile
	limit := c.Query("start")
	var l, o int = 10, 0
	if limit != "" {
		var err error
		o, err = strconv.Atoi(limit)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}
	offset := c.Query("length")
	if offset != "" {
		var err error
		l, err = strconv.Atoi(offset)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}
	result := initializers.DB.Debug().Limit(l).Offset(o).Find(&profiles)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Profile not found.",
		})
		return
	}
	// Respond
	c.JSON(http.StatusOK, profiles)

}

func singleProfile(c *gin.Context, id string) {
	var profile models.ProfileWithUserEmail
	result := initializers.DB.Debug().Table("profiles").
		Select("profiles.*, users.email").
		Joins("left join users on profiles.user_id = users.id").
		Where("profiles.id = ?", id).
		First(&profile)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Profile not found.",
		})
		return
	}
	// Respond
	c.JSON(http.StatusOK, profile)
}

func GetDataByID(c *gin.Context) {
	// Get the profile id from the req params
	id := c.Param("profile_id")
	// Find the profile
	var profile models.ProfileWithUser
	result := initializers.DB.Debug().Table("profiles").
		Joins("User").
		Where("profiles.id = ?", id).
		First(&profile)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Profile not found.",
		})
		return
	}
	// Respond
	c.JSON(http.StatusOK, profile)
}
