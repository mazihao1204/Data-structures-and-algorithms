package main

import "fmt"

//请设计一个栈，除了常规栈支持的pop与push函数以外，还支持min函数，该函数返回栈元素中的最小值。
//执行push、pop和min操作的时间复杂度必须为O(1)。

type MinStack struct {
	stackData []int
	stackMin []int
}

func Constructor() MinStack{
	return MinStack{
		stackData: []int{},
		stackMin: []int{},
	}
}

//先压入stackData,然后判断stackMin是否为空
//为空：压入stackMin,不为空，比较stackData与stackMin栈顶的元素
func (this *MinStack) Push(x int){
	this.stackData = append(this.stackData, x)
	if len(this.stackMin) == 0{
		this.stackMin = append(this.stackMin,x)
	}else {
		dataLen := len(this.stackData)
		minLen := len(this.stackMin)
		if this.stackMin[minLen-1] >= this.stackData[dataLen-1]{
			this.stackMin = append(this.stackMin,x)
		}else {
			this.stackMin = append(this.stackMin,this.stackMin[minLen-1])
		}
	}
}

//
func (this *MinStack) Pop(){
	//value := this.stackData[len(this.stackData)-1]
	//this.stackData = this.stackData[:len(this.stackData)-1]
	//if value == this.stackMin[len(this.stackMin)-1]{
	//	this.stackMin = this.stackMin[:len(this.stackMin)-1]
	//}
	this.stackMin = this.stackMin[:len(this.stackMin)-1]
	this.stackData = this.stackData[:len(this.stackData)-1]
}

func (this *MinStack) Top() int{
	return this.stackData[len(this.stackData)-1]
}

func (this *MinStack) GetMin() int{
	return this.stackMin[len(this.stackMin)-1]
}

func main(){
	obj := Constructor()
	obj.Push(3)
	obj.Push(4)
	obj.Push(5)
	obj.Push(1)
	obj.Push(2)
	obj.Push(1)
	fmt.Printf("%d\n",obj.stackData)
	fmt.Printf("%d\n",obj.stackMin)
	obj.Pop()
	fmt.Printf("%d\n",obj.stackData)
	fmt.Printf("%d\n",obj.stackMin)
}

//nonempty  返回一个新的slice,slice中的元素都是非空字符串
func noneempty(strings []string) []string{
	i := 0
	for _, s := range strings{
		if s != ""{
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}