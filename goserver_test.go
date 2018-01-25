package gotest

import (
	"errors"
	"github.com/KerryJava/goserver/base"
	"github.com/jinzhu/gorm"
	//"goserver"
	"fmt"
	"testing"
)

type User struct {
	ID    int64  `json:"userid"`
	Phone int64  `json:"phone"`
	Name  string `json:"Name"`
}

func (User) TableName() string {
	return "user"
}

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")

	}

	return a / b, nil

}

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil { //try a unit test on function
		t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第一个测试通过了") //记录一些你期望记录的信息

	}

}

//func Test_Division_2(t *testing.T) {
//	t.Error("就是不通过")
//}

func Test_Mysql(t *testing.T) {
	user := User{}
	name := "testuser"
	base.OrmDB.Where(&User{Name: name}).First(&user)

	fmt.Println(user.Name)
	fmt.Println("%v", user)
}

func Benchmark_TimeConsumingPreloadOrm(b *testing.B) {

	b.StopTimer() //调用该函数停止压力测试的时间计数

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		OrmRawQuery()
	}
}

func Benchmark_TimeConsumingOrm(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		//Division(4, 5)
		OrmQuery()
	}
}

func Benchmark_TimeConsumingFunctionRawQuery(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		//Division(4, 5)
		base.RawQuery(false)
	}
}

func Benchmark_TimeConsumingFunctionRawQueryPrepare(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		//Division(4, 5)
		base.RawQuery(true)
	}
}

func OrmQuery() {
	user := User{}
	name := "testuser"
	base.OrmDB.Where(&User{Name: name}).First(&user)
	//	fmt.Println(user.Name)
}

func OrmRawQuery() {
	user := User{}
	name := "testuser"
	var db *gorm.DB = base.OrmDB
	db.Raw("SELECT name, phone FROM user WHERE name = ?", name).Scan(&user)

	//	fmt.Println(user.Name)
}
