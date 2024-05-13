package controller 

import(
    "github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc{
	return func(c *gin.Contex){

	}
}

func GetUser() gin.HandlerFunc{
	return func(c *gin.Contex){
    }
}

func SignUp() gin.HandlerFunc{
	return func(c *gin.Contex){

    }
}

func Login() gin.HandlerFunc{
	return func(c *gin.Contex){
    }
}

func HashPassword(password string ) string{

}

func VerifyPassword(userPassword string, providePassword string)(bool, string) {
	
}