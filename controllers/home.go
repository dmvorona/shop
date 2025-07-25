package controllers

import (
	"github.com/dmvorona/shop/db"
	"github.com/dmvorona/shop/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	var products []models.Product
	db.DB.Find(&products)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":    "Welcome to GoShop",
		"products": products,
	})
}
