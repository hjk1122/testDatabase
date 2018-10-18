package class1

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

type Jumin struct {
	Id     string
	Name   string
	Louhao string
	Menhao string
}

func newXorm() {
	//var engine *xorm.Engine
	var err error
	//var str string
	engine, err = xorm.NewEngine("mysql", "root:12345678@(localhost:3306)/test?charset=utf8")
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "prefix_")
	if err != nil {
		fmt.Printf(err.Error())
	}
	engine.SetTableMapper(tbMapper)
}
func GetData() []JuMin {
	newXorm()

	var juMin []JuMin
	//sql := "SELECT `Id`,`Name` FROM `info`"
	sql := "select * from jumin LIMIT 1"
	engine.SQL(sql).Find(&juMin)
	//engine.Where(sql).Get(info)
	return juMin
}
func GetData1() []Jumin {
	newXorm()

	var info []Jumin

	//sql := "SELECT `Id`,`Name` FROM `info`"
	sql := "select * from jumin LIMIT 1"
	engine.SQL(sql).Find(&info)
	//engine.Where(sql).Get(info)
	//info[0].MenHao="1"
	return info
}
