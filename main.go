package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// CREATE TABLE `restaurants` (
// 	`id` int(11) NOT NULL AUTO_INCREMENT,
// 	`owner_id` int(11) NULL,
// 	`name` varchar(50) NOT NULL,
// 	`addr` varchar(255) NOT NULL,
// 	`city_id` int(11) DEFAULT NULL,
// 	`lat` double DEFAULT NULL,
// 	`lng` double DEFAULT NULL,
// 	`cover` json NULL,
// 	`logo` json NULL,
// 	`shipping_fee_per_km` double DEFAULT '0',
// 	`status` int(11) NOT NULL DEFAULT '1',
// 	`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
// 	`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// 	PRIMARY KEY (`id`),
// 	KEY `owner_id` (`owner_id`) USING BTREE,
// 	KEY `city_id` (`city_id`) USING BTREE,
// 	KEY `status` (`status`) USING BTREE
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

//B2 tạo struct là 1 model của Restaurants
type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

//B3: method TableName trả về 1 string là tên của bảng
func (Restaurant) TableName() string {
	return "restaurants"
}

func main() {
	//B1 connect db
	dsn := os.Getenv("MYSQL_CONNECTION")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}

}

func runService(db *gorm.DB) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	restautants := r.Group("/restautans")
	{
		restautants.POST("", func(ctx *gin.Context) {
			var data Restaurant
			//shouldbind lấy dữ liệu từ request body có gì sẽ nạp vào &data
			if err := ctx.ShouldBind(&data); err != nil {
				//gin.H = map[string]interface{}
				ctx.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}

			if err := db.Create(&data).Error; err != nil {
				ctx.JSON(401, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, data)
		})
		restautants.GET("/:id", func(ctx *gin.Context) {
			id, err := strconv.Atoi(ctx.Param("id"))
			if err != nil {
				ctx.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}
			var data Restaurant
			if err := db.Where("id = ?", id).First(&data).Error; err != nil {
				ctx.JSON(401, gin.H{
					"error": err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusOK, data)
		})
	}
	return r.Run()
}
