package samples

import (
	"container/list"
	"fiff_golang_draft/models"
	"fmt"

	"github.com/gin-gonic/gin"

	"net/http"
	"time"
)

func one() {

	main2()

	var current_time = time.Now()

	fmt.Println("ANSIC: ", current_time.Format(time.ANSIC))
	fmt.Println("UnixDate: ", current_time.Format(time.UnixDate))
	fmt.Println("RFC1123: ", current_time.Format(time.RFC1123))
	fmt.Println("RFC3339Nano: ", current_time.Format(time.RFC3339Nano))
	fmt.Println("RubyDate: ", current_time.Format(time.RubyDate))

	fmt.Println("Hello, World!")
	var init1 = "initial"
	fmt.Println(init1)

	l := list.New() // Initialize an empty list

	{
		var employee models.Employee // dot notation
		employee.FirstName = "Harry"
		employee.LastName = "Potter"
		employee.LeavesTaken = 10
		employee.TotalLeaves = 20
		employee.LeavesRemaining()
		l.PushFront(employee)
	}

	{
		var employee2 models.Employee // dot notation
		employee2.FirstName = "Harry"
		employee2.LastName = "Potter"
		employee2.LeavesTaken = 10
		employee2.TotalLeaves = 20
		employee2.LeavesRemaining()
		l.PushFront(employee2)
	}

	fmt.Println(l)

	//sort.Slice(l, func(i, j int) bool {
	//	return l[i] < l[j]
	//})
	//for _, v := range l {
	//	fmt.Println(v)
	//}
	main3()

}

func main3() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func main2() {
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//
	//// Migrate the schema
	//db.AutoMigrate(&models.Product{})
	//
	//// Create
	//db.Create(&models.Product{Code: "D42", Price: 100})
	//
	//// Read
	//var product models.Product
	//db.First(&product, 1)                 // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	//
	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//db.Model(&product).Updates(models.Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	//db.Delete(&product, 1)
}
