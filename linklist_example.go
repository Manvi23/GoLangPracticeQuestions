//single link list
package main

import "fmt"

type node struct {
	item int
	next *node
}

type linklist struct {
	head *node
	name string
}

func initLinklist(linklistname string) *linklist {
	return &linklist{head: nil, name: linklistname}
}

func (ll *linklist) AddNode(n int) {
	//create a node
	newNode := &node{item: n, next: nil}

	//check is list is empty
	if ll.head == nil {
		fmt.Println("Link List is Empty , Adding First Node")
		ll.head = newNode
	} else {
		fmt.Println("List is Not Empty")
		//if not the first node
		//traverse till the end
		tempNode := ll.head
		for tempNode.next != nil {
			tempNode = tempNode.next
		}
		//it wld be pointing to the last node
		tempNode.next = newNode
	}
}

func (ll *linklist) printList() {
	tempNode := ll.head
	if tempNode == nil {
		fmt.Println("List is Empty")
		return
	}

	for tempNode.next != nil {
		fmt.Println(tempNode.item)
		tempNode = tempNode.next
	}
	//printing last node
	fmt.Println(tempNode.item)
}

//find the node and delete it
func (ll *linklist) Delete(n int) {
	//check is first node
	tempNode := ll.head
	if tempNode.item == n {
		fmt.Println("Deleting the First  node")
		ll.head = tempNode.next
		//node would be garbage collected
	}

	prev := ll.head
	currentNode := ll.head.next

	fmt.Println("Deleting some intermidiate Node")
	for currentNode.next != nil {
		if currentNode.item == n {
			prev.next = currentNode.next
		}
		prev = prev.next
		currentNode = currentNode.next
	}
	//check is last node to be delted
	if currentNode.item == n {
		prev.next = nil
	}

}
func (ll *linklist) reverse() {
	fmt.Println("reversing the Linklist")
	var ptr1, ptr3 *node
	ptr2 := ll.head
	for ptr2 != nil {
		ptr3 = ptr2.next
		ptr2.next = ptr1
		ptr1 = ptr2
		ptr2 = ptr3
	}
	ll.head = ptr1
}

func main() {
	fmt.Println("I am in Main function")
	ll1 := initLinklist("linklist1")
	ll1.AddNode(23)
	ll1.AddNode(12)
	ll1.AddNode(5)
	ll1.AddNode(9)
	ll1.printList()
	//ll1.Delete(5)
	ll1.printList()
	ll1.reverse()
	ll1.printList()

}
