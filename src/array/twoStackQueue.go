package main

import "fmt"

type MyQueue struct {
	StackPush []int
	StackPop []int
}

func Constructor1() MyQueue{
	return MyQueue{}
}

func (this *MyQueue) pushToPop(){
	if len(this.StackPop) == 0{
		for len(this.StackPush) > 0{
			value := this.StackPush[len(this.StackPush)-1]
			this.StackPush = this.StackPush[:len(this.StackPush)-1]
			this.StackPop = append(this.StackPop,value)
		}
	}
}

func (this *MyQueue) Push(s int){
	this.StackPush = append(this.StackPush,s)
	//this.pushToPop()
}

func (this *MyQueue) Pop() int{
	this.pushToPop()
	data := this.StackPop[len(this.StackPop)-1]
	this.StackPop = this.StackPop[:len(this.StackPop)-1]
	return data
}

func (this *MyQueue) Peek() int{
	this.pushToPop()
	return this.StackPop[len(this.StackPop)-1]
}

func (this *MyQueue) Empty() bool{
  return len(this.StackPush) == 0 && len(this.StackPop) == 0
}

func main(){
	obj := Constructor1()
	obj.Push(1)
	obj.Push(2)
	fmt.Printf("%d\n",obj.Peek())
	fmt.Printf("%d\n",obj.Pop())
	fmt.Println(obj.Empty())
	fmt.Printf("%d\n,%d\n",obj.StackPush,obj.StackPop)
}