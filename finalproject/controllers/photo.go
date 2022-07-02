package controllers

import (
	"finalproject/database"
	"finalproject/helpers"
	"finalproject/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photo := models.Photo{}

	if err := ctx.ShouldBindJSON(&photo); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userDataId := userData["id"].(float64)

	photo.Created_at = time.Now().String()
	photo.UserID = uint(userDataId)

	err := db.Create(&photo).Error
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
			"msg":     "Failed to Create Photo",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data": gin.H{
			"id":         photo.ID,
			"title":      photo.Title,
			"caption":    photo.Caption,
			"photo_url":  photo.Photo_url,
			"user_id":    photo.UserID,
			"created_at": photo.Created_at,
		},
	})
}

func GetAllPhotos(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userDataId := userData["id"].(float64)
	photos := []models.Photo{}
	err := db.Preload("User", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("ID", "email", "username")
	}).Find(&photos).Where("Userid : ?", userDataId).Error
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
			"msg":     "Failed to Get Photo List",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data":   photos,
	})
}

func UpdatePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photo := models.Photo{}
	contentType := helpers.GetContentType(ctx)
	if contentType == "application/json" {
		ctx.ShouldBindJSON(&photo)
	} else {
		ctx.ShouldBind(&photo)
	}

	temp, _ := strconv.Atoi(ctx.Param("photoId"))
	photo.ID = uint(temp)
	photo.Updated_at = time.Now().String()

	fmt.Printf("Value Update: %+v\n", photo)
	err := db.Model(&photo).Where("id = ?", photo.ID).Updates(&photo).Error
	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedPhoto := models.Photo{}
	_ = db.First(&updatedPhoto).Error

	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"id":         updatedPhoto.ID,
			"title":      updatedPhoto.Title,
			"caption":    updatedPhoto.Caption,
			"photo_url":  updatedPhoto.Photo_url,
			"user_id":    updatedPhoto.UserID,
			"updated_at": updatedPhoto.Updated_at,
		},
	})
}

func DeletePhoto(ctx *gin.Context) {
	db := database.GetDB()
	photo := models.Photo{}
	temp, _ := strconv.Atoi(ctx.Param("photoId"))

	photo.ID = uint(temp)

	err := db.Where("id= ?", photo.ID).Delete(&photo).Error

	if err != nil {
		log.Println(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
			"msg":     "Failed to Delete Photo",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"message": "Your Photo Has been Successfully Deleted",
		},
	})
}
