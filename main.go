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

//B2 tạo struct là 1 model của TodoItems
type TodoItem struct {
	//'json:"id" gorm:"column:id;` => gọi là tag: là những tuỳ chọn sử dụng để khai báo trường
	//json:
	Id     int    `json:"id" gorm:"column:id;"`
	Title  string `json:"title" gorm:"column:title;"`
	Status string `json:"status" gorm:"column:status;"`
}

type TodoItemUpdate struct {
	//khi update ta để con trỏ string. lúc update bên gorm sẽ kiểm tra xem con trỏ đó có giá trị hay không nếu có giá trị thì lúc đó mới update
	Title *string `json:"title" gorm:"column:title;"`
}

//B3: method TableName trả về 1 string là tên của bảng
func (TodoItem) TableName() string {
	return "todo_items"
}

func main() {
	//B1 connect db
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

	//B4.1 insert new todo item
	// newTodoItem := TodoItem{Title: "English", Status: "Doing"}
	// if err := db.Create(&newTodoItem); err != nil {
	// 	fmt.Println(err)
	// }

	//B4.2 select all with where
	var todos []TodoItem
	db.Where("status = ?", "Doing").Find(&todos)
	fmt.Println(todos)

	//B4.3 select one
	var todo TodoItem
	if err := db.Where("id = 1").First(&todo); err != nil {
		//Khác nhau giữa log và fmt là log mình có thể log thêm thời gian
		log.Println(err)
	}
	fmt.Println(todo)

	//B5 delete
	//db.Table(TodoItem{}.TableName()).Where("id=3").Delete(nil)

	//B6.1 Update c1
	// db.Table(TodoItem{}.TableName()).Where("id=3").Updates(map[string]interface{}{
	// 	"title": "lean go",
	// })

	//B6.1 Update c2
	newTitle := "lean go"
	db.Table(TodoItem{}.TableName()).Where("id=3").Updates(&TodoItemUpdate{Title: &newTitle})
}
