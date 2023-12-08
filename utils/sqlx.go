package utils

import (
	"context"
	"reflect"

	"github.com/zeromicro/go-zero/core/logx"
)

func GetUpdateSql(a any) (string, []interface{}) {

	var setString = "set "
	var values []interface{}
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a) //获取reflect.Type类型

	/**if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}**/
	kd := typ.Kind() //获取到a对应的类别

	if kd != reflect.Struct {
		//fmt.Println("expect struct")
		logx.WithContext(context.Background()).Error("错误的类型,%s", kd.String())
		return setString, nil
	}
	//获取到该结构体有几个字段
	num := val.NumField()
	//fmt.Printf("该结构体有%d个字段\n", num) //4个

	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		tagVal := typ.Field(i).Tag.Get("db")
		//如果该字段有tag标签就显示，否则就不显示
		/**if tagVal == "" {
			//fmt.Printf("Field %v:tag=%v\n", typ.Field(i).Name, tagVal)

			continue
		}**/
		//fmt.Printf("Field %v\n", typ.Field(i).Name)
		//获取到struct标签，需要通过reflect.Type来获取tag标签的值
		name := typ.Field(i).Name
		fieldValue := reflect.ValueOf(a).FieldByName(name)

		child := val.Field(i).Type

		if reflect.String == child().Kind() {
			if len(fieldValue.String()) > 0 {

				setString = setString + tagVal + "= ?, "
				values = append(values, fieldValue)

			}
		}

		if reflect.Int64 == child().Kind() {
			if fieldValue.Int() > 0 {
				if tagVal != "id" {
					setString = setString + tagVal + "= ?, "
					values = append(values, fieldValue.Int())
				}

			}

		}

		//结构体不处理？但是sqlNullString怎么办？
		/**if reflect.Struct == child().Kind() {

		}**/

	}

	setString = setString[0 : len(setString)-1]
	return setString, values

}
