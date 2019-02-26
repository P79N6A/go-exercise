package main

import "github.com/swa19/go-exercise/basic"

func main() {
	//test.TestParamName()
	//
	//var a test.Sequence
	//a = []int{1, 7, 3, 4, 5}
	//fmt.Println(a.String())

	//debug.Main()

	//defer
	// basic.DeferMain()

	//线程同步测试
	// basic.SyncTest()

	// basic.BufioMain()
	// basic.NetMain()

	//flag包使用测试
	//basic.FlagMain()

	//map的make初始化和赋值
	//basic.InitMap()

	//reflect
	//basic.Reflect()

	//// 判断长度
	//var password="12345a6789"
	//if len(password)<8||len(password)>16{
	//	fmt.Println(1)
	//}
	//// 判断是否为纯数字
	//if regexp.MustCompile(`^[0-9]*$`).MatchString(password){
	//	fmt.Println(2)
	//}
	//// 判断是否包含空格
	//if strings.Contains(password," "){
	//	fmt.Println(3)
	//}
	////仅支持数字,字母和字符
	//if !regexp.MustCompile(`^[\w~!@#$%^&*()+,.:;=<>?/|\-[\]\\]{8,16}$`).MatchString(password) {
	//	fmt.Println(4)
	//}
	//fmt.Println(5)

	//var a=[]int{1,2,3}
	//var b=[]int{3,4,5}
	////var result=make([]int,len(a)+len(b))
	////fmt.Println(copy(result,b))
	//fmt.Println(copy(b[len(b):len(a)],a))
	//
	//fmt.Println(b)

	//URL_REGEX:= regexp.MustCompile(`(?i)^(?:http|ftp)s?://(?:(?:[A-Z0-9](?:[A-Z0-9-]{0,61}[A-Z0-9])?\.)+(?:[A-Z]{2,6}\.?|[A-Z0-9-]{2,}\.?)|localhost|\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})(?::\d+)?(?:/?|[/?]\S+)$`)
	//fmt.Println(URL_REGEX.MatchString(""))

	//var str="abcdd"
	////a:=rune(str[0])
	//fmt.Printf("%v",string(str[0]))

	basic.TestGormConnect()
	//basic.TestGopkgGormConnect()

}
