package main

import (
	"fmt"
	"testDatabase/class1"
)

func main() {
	/*
		fmt.Println("输出")
		z := 37
		pi := &z
		ppi := *pi
		fmt.Println(z, *pi, ppi)
		ppi++
		fmt.Println(z, *pi, ppi)
		//bufer:=make([]byte,20,60)
		grid1 := make([][]int, 3)
		for i := range grid1 {
			grid1[i] = make([]int, 3)
		}
		grid1[1][0], grid1[1][1], grid1[1][2] = 8, 6, 2
		fmt.Println(grid1)
		s:=[]string{"A","B","C","D","E","F","G"}
		t:=s[2:6]
		fmt.Println(t,s,"=",s[:4])
		s[3]="x"
		t[len(t)-1]="y"
		fmt.Println(t,s,"=",s[:4],"+",s[4:])
	*/
	var aa int = 99
	class1.Check(aa)
	class1.Classifier(1, "wo", 12.78, true)
	var JuMin []class1.JuMin
	JuMin = class1.GetData()
	fmt.Println(JuMin[0])

	/*
		var data []map[string]string
		data = class1.GetDataJuMin1()
		fmt.Print(data)
	*/
}

/*
	for row :=range JuMin{

			fmt.Println(JuMin[0].GzDanWei)

		}
	}
*/
//fmt.Print(JuMin[0])
/*
	var jm []class1.JuMin
	jm=class1.GetDataJuMin()
	fmt.Printf(jm[0].ID)
*/
//fmt.Println(class1.GetDataJuMin1())
//}
