package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

// User 모델 정의
type User struct {
	ID      uint   `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Email   string `json:"email" gorm:"type:varchar(100);unique_index"`
	RegDate string `json:"reg_date"`
}

var db *gorm.DB
var err error

// 데이터베이스 연결 설정
func init() {
	dsn := "host=localhost user=postgres dbname=postgres password='qhdks12#' port=5432 sslmode=disable TimeZone=Asia/Seoul"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 데이터베이스 마이그레이션
	db.AutoMigrate(&User{})
}

func main() {
	router := gin.Default()

	router.GET("/users", GetUsers)
	router.GET("/users/:id", GetUser)
	router.POST("/users", CreateUser)
	router.PUT("/users/:id", UpdateUser)
	router.DELETE("/users/:id", DeleteUser)

	router.Run(":8080")
}

// GetUsers: 모든 사용자 조회
func GetUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

// GetUser: 단일 사용자 조회
func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// CreateUser: 사용자 생성
func CreateUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	db.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// UpdateUser: 사용자 정보 업데이트
func UpdateUser(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DeleteUser: 사용자 삭제
func DeleteUser(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")
	d := db.Where("id = ?", id).Delete(&user)
	fmt.Println(d)
	c.JSON(http.StatusOK, gin.H{"id #" + id: "deleted"})
}
