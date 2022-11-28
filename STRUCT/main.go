package main 

import "fmt"

type people struct {
	fname string
	lname string 
	infor infor
}
type infor struct {
	address string
}

func (n *people) changelName(lnameNew string){
	n.lname = lnameNew
}

func (n people) in(){
	fmt.Println("%#v\n",n)
}

func main() {
	per := people{ fname: "leon",
				   lname:"selina",
				   infor: infor{
					address:"08 hV",
				   },}
	// per.fname = "Tran"
	// fmt.Println("%v",per)
	perPoint := &per
	perPoint.changelName("Boss")
	per.in()
}