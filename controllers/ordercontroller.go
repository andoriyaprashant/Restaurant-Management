package controller 

import(
	"gituhb.com/gin-gonic/gin"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.withTimeout(context.Background(), 100*time.Second)
		
	    result, err := orderCollectionFind(context.TODO(), bson.M{})
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing order items"})

		}
		var allOrders []bson.M
		if err = result.All(cyx, &allOrders); err !=  nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allOrders[0])

	}
}

func GetOrder() gin.HandlerFunc{
	return func(c *gin.Context){
		
		var ctx, cancel = context.withTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("order_id")
		var food modles.Order
		err := foodCollection.Findone(ctx, bson.M{"order_id":order_Id}).Decode(&order)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H("error": "error occured while fetching the order"))

		}
		c.JSON(http.StatusOK, order)
	}
}

func CreateOrder() gin.HandlerFunc{
	return func(c *gin.Context){
		var table models.Table 
		var order model.Order

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errr.Error()})
			return
		}
		validationErr := validate.Struct(order)

		if validationErr!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return 
		}
		if order.Table_id! = nil {
			err := tableCollection.FindOne(cyx, bson.M{"table_id": order.Table_id}).Decode(&table)
			defer cancel()
			if err!=nil{
				msg:= fmt.Sprintf("message: Table was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return 
		}
	}
	order.Created_at, _ =  time.Parse(time.RFC3339, time.Now(). Format(time.RFC3339)) 
	order.Updated_at,_ =  time.Parse(time.RFC3339, time.Now(). Format(time.RFC3339))

	order.ID = primitive.NewObjectID()
	order.Order_id = order.ID.Hex()

	result, insertErr := orderCollection.InsertOne(cyx, order)

	if insertErr != nil {
		msg := fmt.Sprintf("order item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error"; msg})
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result)
 }
}


func UpdateOrder() gin.HandlerFunc{
	return func(c *gin.Context){
		// var ctx, cancel = context.withTimeout(context.Background(), 100*time.Second)
		var menu models.Table 
		var food models.Order

		var updateObj primitive.D 

		orderId := c.Param("order_id")
		if err := c.BindJSON(&order); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if order.Table_id!=nil{
			err := menuCollection.FindOne(ctx, bson.M{"tabled_id": food_Table_id}).Decode(&table)
			defer cancel()
			if err != nil{
				msg := fmt.Sprintf("messages:Menu was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
				return 
			}
			updateObj = append(updateObj, bson.E{"menu", order.Table_id})
		}
		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj= append(updateObj, bson.E{"updated_at", food.Updated_at})

		upsert : true
		filter := bson.M{"order_id":orderId}
		opt := options.UpdateOptions{
			Upsert:&upsert,
		}
		result, err := orderCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			$opt,
		)

		if err != nil {
			msg:= fmt.Sprintf("order item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return 
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func OrderItemOrderCreator(order models.Order) string{
	order.Created_at, _ =  time.Parse(time.RFC3339, time.Now(). Format(time.RFC3339)) 
	order.Updated_at,_ =  time.Parse(time.RFC3339, time.Now(). Format(time.RFC3339))

	orderCollection.InsertOne(ctx,order)
	defer cancel()

	return order.Order_id

}