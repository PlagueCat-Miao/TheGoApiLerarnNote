package Pointer

import (
	"fmt"
	"log"
)
import "testing"

type  Name struct {

	name string
	ch  (chan int)

}
func input(a Name){
	a.name= "k"
}
func inputP(a* Name){
	a.name= "kP"
}


func TestPoint(t *testing.T) {
	t.Log("Hellcat: Pointer 请求测试")
	var n  Name;
	n1 := new(Name)
	var n2,n3 *Name
	n2=&n;
	n3=n1;
	n2.name="hell"
	n3.name="cat"

	fmt.Println("n",n,"n1:",n1)
	var fn,fn2  Name;
	input(fn)
	inputP(&fn2)
	fmt.Println("n",fn.name,"n1:",fn2.name)

	var gfn Name;
	gfn.ch = make(chan int ,10)
    end := make(chan int ,10)
	go func(gfn Name, end chan int){
	var x int
	x = <-gfn.ch
	log.Println("go func copy-gfn.ch : ", x);
	fmt.Println("可见类Name虽然是副本，但是副本的chan是指针 并指向同一值，故能相互影响")
	fmt.Println("副本chan: end 同理也是指针，故能与函数外相互影响")
	end <- 4
	}(gfn,end)

	log.Println("start: gfn.ch<- ", 2)
	gfn.ch <-2

	<-end
	t.Log("Hellcat: Pointer 测试完成")
}

