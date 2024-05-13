package controller 

import(
	"gituhb.com/gin-gonic/gin"
)

func GetOrderItems() gin.HandlerFunc{
	return func(c *gin.Context){
		
	}
}

func GetOrderItemByOrder() gin.HandlerFunc{
	return func(c *gin.Context){
		
	}
}

func ItemsByOrder(id string) (orderItems []primitive.M, err error){

}

func GetOrderItem() gin.HandlerFunc{
	return func(c *gin.Context){

	}
}

func UpdateOrderItem() gin.HandlerFunc{
	return func(c* gin.Context)
}

func CreateOrderItem() gin.HandlerFunc{
	return func(c *gin.Context){
		
	}
}
