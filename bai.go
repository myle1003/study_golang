package main

import "fmt"
import "strings"
import "io/ioutil"
import "log"
// import "os"
import "math/rand"
import "time"


type arrayBai []string

func createArrayBai() arrayBai{
	kq := arrayBai{}
	nhieunuoc := []string{"co", "ro", "chuon", "bich"}
	nhieunut := []string{"2","3","4","5"}
	for _, nuoc := range nhieunuoc{
		for _, nut := range nhieunut{
			bai :=nut + " "+nuoc
			kq = append(kq,bai)
		}
	}
	return kq
}
 
func chiaBai(n arrayBai, sl int) (arrayBai, arrayBai){
	return n[:sl],n[sl:]
}

func (n arrayBai) in(){
	fmt.Println(n)
}

func (n arrayBai) chuyenThanhString() string{
	return strings.Join(n, ",")
}

func (n arrayBai) luufile(tenFile string) error{
	data := []byte(n.chuyenThanhString())
	return ioutil.WriteFile(tenFile,data,0666)
}

func taoBaiTuFile(tenFile string) (arrayBai,error){
	data,err :=ioutil.ReadFile(tenFile)
	if err!= nil {
		log.Println("doc file ko dc ", err);
		// os.Exit(1);
		return arrayBai{},err
	}
	chuoiBai := string(data)
	mang := strings.Split(chuoiBai,",")
	arrayLaBai :=arrayBai(mang)
	return arrayLaBai,nil
}

func (n arrayBai) xoasbai(){
	rand.Seed(time.Now().UnixNano())
	for index := range n {
		rn := rand.Intn(len(n))
		// doi choi n[index] n[rn]
		n[index], n[rn] = n[rn], n[index]
	}
}