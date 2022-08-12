package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// CREATE TABLE `todo_items` (
// 	`id` int NOT NULL AUTO_INCREMENT,
// 	`title` varchar(150) CHARACTER SET utf8 NOT NULL,
// 	`status` enum('Doing','Finished') DEFAULT 'Doing',
// 	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 	`updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
// 	PRIMARY KEY (`id`)
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

//B1 tạo struct là 1 model của TodoItems
type TodoItem struct {
	//'json:"id" gorm:"column:id;` => gọi là tag: là những tuỳ chọn sử dụng để khai báo trường
	//json:
	Id     int    `json:"id" gorm:"column:id;"`
	Title  string `json:"title" gorm:"column:title;"`
	Status string `json:"status" gorm:"column:status;"`
}

//method TableName trả về 1 string là tên của bảng
func (TodoItem) TableName() string {
	return "todo_items"
}

func main() {
	mysqlConnStr, ok := os.LookupEnv("MYSQL_CONNECTION")

	if !ok {
		log.Fatalln("Missing MySQL connection string")
	}

	dsn := mysqlConnStr
	println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	newTodoItem := TodoItem{Title: "English", Status: "Doing"}
	db.Create(&newTodoItem)
	fmt.Println(newTodoItem)
	//insert new todo item
}
