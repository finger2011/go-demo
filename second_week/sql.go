package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

// ErrNoRows在50行左右 ^_^
func main() {
	// id, name, age := getRowByTableAndID("user_demo", 1)
	// fmt.Println("id:", id)
	// fmt.Println("name:", name)
	// fmt.Println("age:", age)
	var users, len, err = getRows()
	if err != nil {
		fmt.Printf("err:%v", err)
		return 
	}
	fmt.Printf("\nlen:%d", len)
	fmt.Printf("\nusers:%v", users)
}

//User user
type User struct {
	// fields []string
	// value map[string]string

	Id      int
	Name    string
	// Age     int
	Address string
}

func getRows() ([]User, int, error) {
	db := mysqlConn()
	defer mysqlClose(db)
	// get rows info
	rows, err := db.Query("select * from user_demo limit 10")
	if err != nil {
		if err == sql.ErrNoRows {
			// 行数为0 是一个业务现象，所以不应该是一个需要抛出的错误
			//所以在接口返回中增加了行数字段，实际业务中感觉也经常用到这个字段
			// 故业务层在调用接口时判断下len是否为0来处理相应的业务逻辑
			return []User{}, 0, nil
			// panic(err.Error())
		} 
		// 不为0行的错误，需要抛出交给业务层去处理
		return []User{}, -1, wrapError(err)
	}
	defer rows.Close()
	//get cloumns info
	clos, err := rows.Columns()
	if err != nil {
		// 
		return []User{}, -1, wrapError(err)
	}

	users := make([]User,0,10)

	for rows.Next() {
		var user = User{}
		var val = reflect.ValueOf(&user)

		// rowdata用来接收从数据库返回的每一行row数据
		var rowdata = make([]interface{}, len(clos))
		for j := 0; j < len(clos); j++ {
			rowdata[j] = makeEmptyStr()
		}
		err = rows.Scan(rowdata...)
		if err != nil {
			// 读取row失败，因为是更底层的错误，所以需要继续往上层抛出
			return []User{}, -1, wrapError(err)
		}

		for i := 0; i < len(clos); i++ {
			// 设置user的field的值
			var cloumnNameRune = []rune(clos[i])
			if cloumnNameRune[0] >= 97 && cloumnNameRune[0] <= 122 {
				cloumnNameRune[0] -= 32
			}
			var cloumnName = string(cloumnNameRune)
			var field = val.Elem().FieldByName(cloumnName)
			
			if field.IsValid() {
				if field.CanSet() {
					if field.Kind() == reflect.Int {
						var s = rowdata[i].(*sql.NullString)
						if s.Valid {
							var intValue = s.String
							var value int
							if intValue == "" {
								value = 0
							} else {
								value, err = strconv.Atoi(intValue)
								if err != nil {
									// xxxx具体内容可以自定
									return []User{}, -1, wrapError(errors.New("xxxxx"))
									// panic(fmt.Sprintf("field %s convert string to int error:%v", cloumnName, value))
								}
							}
							if !field.OverflowInt(int64(value)) {
								field.SetInt(int64(value))
							} else {
								return []User{}, -1, wrapError(errors.New("xxxxx"))
								// panic(fmt.Sprintf("field %s overflow int:%v", cloumnName, value))
							}
						} else {
							return []User{}, -1, wrapError(errors.New("xxxxx"))
							// panic(fmt.Sprintf("field %s row value invalid:%v", cloumnName, rowdata[i]))
						}
					} else if field.Kind() == reflect.String {
						var value = ""
						if (rowdata[i].(*sql.NullString)).Valid {
							value = (rowdata[i].(*sql.NullString)).String
						}
						field.SetString(value)
					} else {
						panic(fmt.Sprintf("field %s unknow type", cloumnName))
					}
				} else {
					panic(fmt.Sprintf("field %s can not be set", cloumnName))
				}
			} else {
				// 字段不在预定义的User结构体中
				continue
				// panic(fmt.Sprintf("field %s is not valid", cloumnName))
			}
		}
		users = append(users, user)
	}
	
	return users, len(clos), nil
}

func makeEmptyStr() *sql.NullString {
	var s sql.NullString
	
	return &s
}

// func getRowByTableAndID(table string, id int) (int, string, int) {
// 	db := mysqlConn()

// 	defer mysqlClose(db)

// 	var name string
// 	var age int
// 	err := db.QueryRow("select id,name,age from user_demo where id = ?", id).Scan(&id, &name, &age)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			panic(err.Error())
// 		} else {
// 			panic(err.Error())
// 		}
// 	}
// 	return id, name, age
// }

func wrapError(err error) MyError {
	// 包装错误
	// 自定义错误信息
	var newError = MyError{err:err}
	// 做些包装动作
	return newError
}

func mysqlConn() *sql.DB {
	db, err := sql.Open("mysql", "root:@/demo")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("db connection success")
	return db
}

func mysqlClose(db *sql.DB) {
	db.Close()
	fmt.Println("\n db connection closed")
}

// MyError 自定义错误
// 从一个phper的角度上看，基于Opaque errors 的模式，虽然解耦了调用者和error，
// 但是对于一些出错信息保存的并不是特别完善，所以习惯还是error types模式
// 在写php处理错误时，继承树是类似baseException -> MysqlException -> MysqlParameterException
// 这样一个很简单的instance of就可以判断异常类型，并且能够拿到相关的出错信息
// 由于go没有继承，struct式的继承又会浪费很多重复的空间，所以不知道有什么成熟的方案？
type MyError struct {
	Message string
	File string
	Line int
	Code int
	// wrap的时候是否需要保存具体error还没考虑好，但是这里还是加上了
	err error
	// 可能需要一些环境信息之类的，没有具体扩展，所以直接扔一个env在这了
	env string
} 

func (e MyError) Error() string {
	return fmt.Sprintf("ERROR FILE:%s\nERROR LINE:%d\nERROR MESSAGE:%s\nERROR CODE:%d", e.File, e.Line, e.Message, e.Code)
}
