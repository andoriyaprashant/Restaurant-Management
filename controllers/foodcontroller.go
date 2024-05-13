package controller 

import(
	"context"
	"golang-restaurant-management/database"
	"golang-restaurant-management/models"
    "net/http"
	"time"
	"gituhb.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/bluesuncorp/validator.v5"
	"gopkg.in/mgo.v2/bson"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func GetFoods() gin.HandlerFunc{
	return func(c *gin.Context){

		var ctx, cancel = context.withTimeout(context.Background(), 100*time.Second)
		recordPerPage,  err := strvconv.Atoi(c.Query("recordPerPage"))
        if err!= nil || recordPerPage < 1 {
			recordPerPage = 10 
		}

		page, err := strcov.Atoi(c.Query("page"))
		if err != nil || page < 1{
		    page = 1 	
		}
        startIndex := (page-1 ) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{{}}}}
		groupStage := bson.D{{"$group", bson.D{{"id", bson.D{{"_id", "null"}}},{"total_count", bson.D{{"$sum,1"}}},{"total_count", bson.D{{"$sum", 1}}},{"data", bson.D{{"$push", "$$ROOT"}}} }}}
		projectStage := bson.D{
			{
				"$project", bson.D{
					{"_id", 0},
					{"total-count", 1},
					{"food_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}}
				}
			}
		}
		result, err := foodCollection.Aggregate(cyx, mongo.Pipeline{
			matchStage, groupStage, projectStage

		})
		defer cancel()
		if err != nil{
			c.JSON{http.StatusInternalServerError, gin.h{"error":"error occured while listing food items"}}
		}
		var allfoods []bson.M
		if err = result.All(cyx, &allFoods); err != nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allFoods[0])
	}
}

func GetFood() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.withTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("food_id")
		var food modles.Food
		err := foodCollection.Findone(ctx, bson.M{"food_id"}).Decode(&food)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H("error": "error occured while fetching the food "))

		}
		c.JSON(http.StatusOK, food)
	}
}

func CreateFood() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.withTimeout(context.Background(), 100*time.Second)
		var menu models.Menu 
		var food models.Food

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errr.Error()})
			return
		}

		validationErr := validate.Struct(food)
		if validationErr != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error": validationErr.Error()})
			return
		}
		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food_Menu_id}).Decode(&menu)
		defer cancel()
		if err != nill {
			msg := fmt.Sprintf("menu was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return

		}
        food.Created_at, _= time.Parse(time.RFC3339, time.Now(). Format(time.RFC3339))
		food.Updated_at, _= time.Parse(time.RFC3339, time.Now(). Format(time.RFC3339))		
		food_ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = toFixed(*food.Price, 2 )
		food.Price = &num

		result, insetErr := foodCollection.InsertOne(ctx, food)
		if insetErr != nil{
			msg := fmt.Sprintf("Food item was not created")
			c.JSON(http.StatusInternalServerError,gin.H{"error": msg})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func round(num float64) int {
    return int(num + math.Copysign(0.5,num))
} 

func toFixed(num float64, precision int) float64{
	output := mat.Pow(10, float64(precision))
	return float64(round(num*output)) / output

}

func UpdateFood() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.withTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		var food models.Food

		foodId := c.Param("food_id")

		if err := c.BindJSON(&food); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return 
		} 

		var updateObj primitive.D

		if food.Name != nil{
             updateObj = append(updateObj, bson.E{"name", food.Name})
		}

		if food.Price != nil{
			updateObj = append(updateObj, bson.E{"price", food.Price})

		}

		if food.Food_image != nil{
			updateObj = append(updateObj, bson.E{"food_image", food.Food_image})

		}
        
		if food.Menu_id != nil{
			err := menuCollection.FindOne(ctx, bson.M{"menu_id": food_Menu_id}).Decode(&menu)
			defer cancel()
			if err != nil{
				msg := fmt.Sprintf("messages:Menu was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
				return 
			}
			updateObj = append(updateObj, bson.E{"menu", food.Price})
		}

        food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj= append(updateObj, bson.E{"updated_at", food.Updated_at})

		upsert := true
		filter := bson.M{"food_id": foodID}

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := foodCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj}

			},
			&opt,

		)

		if err!=nil {
			msg:= fmt.Sprint("food item update failed  ")
			c.JSON(http.StatusInternalServerError, gin.H{"error":msg})
			return
		}
		c.JSON(http.StatusOK, result)

	}
}


