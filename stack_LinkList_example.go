//stack
package main

import "fmt"

type stack interface {
	Push(item int)
	Pop() int
	PrintStack()
}

type node struct {
	item int
	next *node
}

type stackList struct {
	top       *node
	stackname string
}

func initStack(st string) *stackList {
	return &stackList{stackname: st, top: nil}
}

func (st *stackList) push(i int) {
	fmt.Println("Pushed", i)
	//create a node
	newNode := &node{item: i, next: nil}
	//check if first node or stack is empty
	if st.top == nil {
		st.top = newNode
		return
	}

	newNode.next = st.top
	st.top = newNode
}

func (st *stackList) pop() (i int) {
	//fmt.Println("Popping the value")
	fmt.Println("Popped", st.top.item)
	itempopped := st.top.item
	st.top = st.top.next
	return itempopped
}

func (st *stackList) PrintStack() {
	tempNode := st.top
	for tempNode.next != nil {
		fmt.Println(tempNode.item)
		tempNode = tempNode.next
	}
	fmt.Println(tempNode.item)
}

func main() {
	fmt.Println("In Stack Main function")
	stack1 := initStack("mystack")
	stack1.push(1)
	stack1.push(2)
	stack1.push(3)
	stack1.push(4)
	stack1.push(5)
	stack1.PrintStack()
	stack1.pop()
	stack1.PrintStack()

}
