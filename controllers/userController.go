package controllers

import "github.com/gin-gonic/gin"

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func HashPassword(password string) string {

}
func VerifyPassword(password string, providePassword string) (bool, string) {

}