package main

// import "fmt"
import "os"

func main() {
	// var laBai string = "3 bích"
	// laBai:="3 bich"
	// laBai := create()
	// fmt.Println("hello word")
	// fmt.Println(laBai)

	// var arrayBai [3]string
	// arrayBai := make([]string, 3)
	// arrayBai[0]= "2 co"
	// arrayBai[1]= "4 co"
	// arrayBai[2]= "5 co"
	// arrayBai := []string{"2 ro", "3 ro"}
	// arrayBai = append(arrayBai,"5 ro", "6 bich")
	// fmt.Println(arrayBai)
	// arrayLaBai := arrayBai{"3 ro", "4 ro"}
	// for index, bai := range arrayLaBai{
	// 	fmt.Println(index, bai)
	// }

	arrayLaBai := createArrayBai()
	// bai1, baiconlaij := chiaBai(arrayLaBai,4)
	// bai1.in()
	// baiconlaij.in()
	// arrayLaBai.in()
	// fmt.Println(arrayLaBai.chuyenThanhString())
	// arrayLaBai.luufile("test")

	arrayLaBai = append (arrayLaBai,"9 co")

	arrayLaBai.in()

	arrayLaBai, err := taoBaiTuFile("test1")
		if err !=nil{
			os.Exit(1)
		}


	arrayLaBai.in()
	// arrayLaBai.xoasbai()
	// arrayLaBai.in()

}

func create() string{
	return "3 co"
}