package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	. "temp/db"
	. "temp/model"
)

func AddStudent(context *gin.Context) {
	// 结构体的对象和表记录
	db := GetDB()
	// 添加老师对象
	t1 := Teacher{BaseModel: BaseModel{ID: 1}, Name: "yuan", Tno: 1001, Pwd: "123"}
	db.Create(&t1)
	fmt.Println(t1.ID)
	t2 := Teacher{BaseModel: BaseModel{ID: 2}, Name: "rain", Tno: 1002, Pwd: "234"}
	db.Create(&t2)
	t3 := Teacher{BaseModel: BaseModel{ID: 3}, Name: "eric", Tno: 1002, Pwd: "345"}
	db.Create(&t3)

	// 批量创建班级
	c1 := Class{Name: "软件一班", Num: 78, TeacherID: 1}
	c2 := Class{Name: "软件二班", Num: 88, TeacherID: 2}
	c3 := Class{Name: "计算机科学与技术1班", Num: 56, TeacherID: 3}
	c4 := Class{Name: "计算机科学与技术2班", Num: 32, TeacherID: 1}

	classes := []Class{c1, c2, c3, c4}
	db.Create(&classes)

	for i, class := range classes {
		fmt.Println(":::", i, class)
	}

	// 批量创建课程

	course01 := Course{Name: "计算机网络", Credit: 3, Period: 16, TeacherID: 1}
	course02 := Course{Name: "数据结构", Credit: 2, Period: 24, TeacherID: 1}
	course03 := Course{Name: "数据库", Credit: 2, Period: 16, TeacherID: 2}
	course04 := Course{Name: "数字电路", Credit: 3, Period: 12, TeacherID: 2}
	course05 := Course{Name: "模拟电路", Credit: 1, Period: 8, TeacherID: 2}
	courses := []Course{course01, course02, course03, course04, course05}
	db.Create(&courses)

	s1 := Student{Name: "李四", Sno: 2002, Pwd: "123", ClassID: 2}
	db.Create(&s1)
	fmt.Println("s1:::", s1)
	fmt.Println("s1:::", s1.Class)
	fmt.Println("s1:::", s1.Courses)

	db.Where("name in ?", []string{"数据结构", "数据库"}).Find(&courses)
	fmt.Println("courses:", courses)

	s2 := Student{Name: "王武", Sno: 2002, Pwd: "123", ClassID: 2, Courses: courses}
	db.Create(&s2)

	s3 := Student{Name: "赵六", Sno: 2002, Pwd: "123", ClassID: 2}
	db.Create(&s3)
	db.Model(&s3).Association("Courses").Append(courses)
	context.String(200, "OK")
}

func GetAllStudent(context *gin.Context) {
	db := GetDB()
	var students []Student
	db.Preload("Class").Preload("Courses").Find(&students)
	/*context.JSON(200, gin.H{
		"stus": students,
	})*/
	context.HTML(200, "student.html", gin.H{
		"students": students,
	})

}
