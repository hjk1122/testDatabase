package class1

import "fmt"

type colStruct struct {
	//colName string
	colMemo  string
	colValue string
}

var JuminCol JuMinData

//var sql="select * from jumin where Yuan='云丽' LIMIT 1"
var sql = "select * from jumin"

//表结构数据
type JuMinData struct {
	ID         colStruct
	ChkDate    colStruct
	Yuan       colStruct
	Louhao     colStruct
	MenHao     colStruct
	MenPai     colStruct
	HuKou      colStruct
	HuJi       colStruct
	XianZhuZhi colStruct
	Name       colStruct
	GuanXi     colStruct
	Sex        colStruct
	ShenFenID  colStruct
	Zhengzhi   colStruct
	Birthday   colStruct
	Old        colStruct
	Wenhua     colStruct
	Hunfou     colStruct
	Minzu      colStruct
	Telephone  colStruct
	Shiye      colStruct
	Tuixiu     colStruct
	GzDanWei   colStruct
	Techang    colStruct
	Dibahu     colStruct
	Gulaohu    colStruct
	Tefu       colStruct
	Junren     colStruct
	Canjiren   colStruct
	Memo       colStruct
}
type JuMin struct {
	ID         string
	ChkDate    string
	Yuan       string
	Louhao     string
	MenHao     string
	MenPai     string
	HuKou      string
	HuJi       string
	XianZhuZhi string
	Name       string
	GuanXi     string
	Sex        string
	ShenFenID  string
	Zhengzhi   string
	Birthday   string
	Old        string
	Wenhua     string
	Hunfou     string
	Minzu      string
	Telephone  string
	Shiye      string
	Tuixiu     string
	GzDanWei   string
	Techang    string
	Dibahu     string
	Gulaohu    string
	Tefu       string
	Junren     string
	Canjiren   string
	Memo       string
}

func InitCol() {
	JuminCol = JuMinData{colStruct{"编号", ""},
		colStruct{"核准时间", ""},
		colStruct{"园", ""},
		colStruct{"楼号", ""},
		colStruct{"门号", ""},
		colStruct{"门牌", ""},
		colStruct{"户口状态", ""},
		colStruct{"户籍地", ""},
		colStruct{"现住址", ""},
		colStruct{"姓名", ""},
		colStruct{"与户主关系", ""},
		colStruct{"性别", ""},
		colStruct{"身份证号", ""},
		colStruct{"政治面目", ""},
		colStruct{"出生年月", ""},
		colStruct{"年龄", ""},
		colStruct{"文化程度", ""},
		colStruct{"婚否", ""},
		colStruct{"民族", ""},
		colStruct{"家庭电话", ""},
		colStruct{"是否失业", ""},
		colStruct{"是否退休", ""},
		colStruct{"工作单位", ""},
		colStruct{"特长", ""},
		colStruct{"是否低保户", ""},
		colStruct{"是否孤老户", ""},
		colStruct{"是否特扶家庭", ""},
		colStruct{"是否军人", ""},
		colStruct{"是否残疾人", ""},
		colStruct{"备注", ""}}
}
func GetDataJuMin() []JuMin {
	newXorm()
	var jm []JuMin
	var err error
	//var flag bool
	//flag,err=engine.Where(sql).Get(&jm)
	err = engine.SQL(sql).Find(&jm)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	} else {
		return jm
	}

}
func GetDataJuMin1() []map[string]string {
	newXorm()
	var data []map[string]string
	//sql := "SELECT * FROM `jumin`"
	sql := "select  * from jumin where Yuan='云丽' LIMIT 1"
	var err error
	data, err = engine.QueryString(sql)
	if err != nil {
		fmt.Printf(err.Error())
		return nil
	}
	return data
}

//获得居民信息
func GetDataJuMinStr() []map[string]string {
	newXorm()

	var data []map[string]string
	//sql := "SELECT `Id`,`Name` FROM `info`"
	sql := "SELECT * FROM `JuMinData`"
	var err error
	if err != nil {
		data, err = engine.QueryString(sql)
	}

	return data
}
