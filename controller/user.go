package controller

import (
	"crypto/sha256"
	"fmt"

	"github.com/cheenusrivel/nft-go-api/config"
	"github.com/cheenusrivel/nft-go-api/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find((&users))
	c.JSON(200, &users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.FirstOrCreate(&user)
	c.JSON(200, ReceiptHash(&user))
}

// Used hasing due to no decryption required at client side
func ReceiptHash(user *models.User) string {
	h := sha256.New()
	h.Write([]byte(user.NRIC + user.WalletAddress))
	bs := h.Sum(nil)
	r := fmt.Sprintf("%x\n", bs)
	return r
}

func DeleteUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Where("NRIC = ?", c.Param(("nric"))).Delete((&user))
	c.String(200, "Row deleted")
}

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("NRIC = ?", c.Param(("nric"))).First((&user))
	c.BindJSON(&user)
	config.DB.Save((&user))
	c.JSON(200, &user)
}
