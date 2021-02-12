package follow

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"hvxahv/pkg/db"
	"hvxahv/pkg/utils"
	"log"
)
type Follow struct {
	ID    string `json:"_id"`
	Actor string `json:"actor"`
	Date  string `json:"date"`
	Name  string `json:"name"`
}

// GetFollowerHandler 获取关注者
func GetFollowerHandler(c *gin.Context) {
	name := utils.GetUserName(c)
	r := GetFollow(name, "follower")
	c.JSON(200, gin.H{
		"res": r,
		"count": len(r),
	})
}

func GetFollow(name, collection string) []string {
	db := db.GetMongo()
	f := bson.M{"name": name}
	log.Println(name)
	co := db.Collection(collection)
	var i []string
	findA, err := co.Find(context.TODO(), f, nil)
	if err != nil {
		log.Println(err)
	}
	for findA.Next(context.TODO()) {
		var el Follow
		if err := findA.Decode(&el); err != nil {
			log.Println(err)
		}
		i = append(i, el.Actor)
	}
	if err := findA.Err(); err != nil {
		log.Println(err)
	}
	_ = findA.Close(context.TODO())


	return i
}