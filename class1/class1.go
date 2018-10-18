package class1

import (
	"fmt"
)

func Check(x int) {
	var i interface{} = x
	var s interface{} = []string{"left", "right"}
	j := i.(int)
	fmt.Printf("%T->%d\n", j, j)
	if i, ok := i.(int); ok {
		fmt.Printf("%T->%d\n", i, j)
	}
	if s, ok := s.([]string); ok {
		fmt.Printf("%T->%q\n", s, s)
	}
}

//判断输入值的类型
func Classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("参数 #%d 是bool型\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int8, int16, int32, int64:
			fmt.Printf("param #%d is an int\n", i)
		case uint, uint8, uint16, uint32, uint64:
			fmt.Printf("param #%d is an unsigned int\n", i)
		case nil:
			fmt.Printf("param #%d is nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d's type is unknow\n", i)
		}
	}
}

/*
func TypeChange(items...interface{}) string
{
var	resurt string
switch x.(type) {
case int, int8, int16, int32, int64:

}

}
*/
/*
var MA []byte('{"name":"Massachusetts"}')
	"area":27336,
	"water":25.7
	"senators":["John Kerry", "Scott Brown"]}')'
func JsonString() {
var object interface{}
if err := json.Unmarshal(MA, &object);err!=nil{
	fmt.Println(err)
}else {
	jsonObject := object.(map[string]interface{})
	{
		fmt.Println(jsonObjectAsString(jsonObject))
	}
}
*/
