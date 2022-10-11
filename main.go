package main

import (
	"errors"
	"fmt"
	"gorm/database"
	"gorm/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	// createUser("kresna@gmail.com")
	// getUserById(1)
	// updateUserByID(1, "vespri@gmail.com")
	// createProduct(1, "Kijang", "Innova")
	getUsersWithProducts()
}

func createUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data :", err)
		return
	}

	fmt.Println("New User Data :", User)
}

func getUserById(id uint) {
	db := database.GetDB()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user :", err)
	}

	fmt.Printf("User Data : %+v \n", user)
}

func updateUserByID(id uint, email string) {
	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error updating user data :", err)
		return
	}
	fmt.Printf("Update user's email : %+v \n", user.Email)
}

func createProduct(userID uint, brand string, name string) {
	db := database.GetDB()

	Product := models.Product{
		UserID: userID,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating product data :", err.Error())
		return
	}

	fmt.Println("New Product Data :", Product)
}

func getUsersWithProducts() {
	db := database.GetDB()

	users := models.User{}

	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user datas with products :", err.Error())
		return
	}

	fmt.Println("User Datas With Products")
	fmt.Printf("%+v", users)
}
