package main

import (
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

/**
创建工厂函数，设置树的值
*/
func createNode(value int) *tree {
	return &tree{value: value}
}

/**
当接收者为值接收者时，会将该对象拷贝一份到方法中 多用于查询数据
*/
func (this tree) print() {
	fmt.Println(this)
}

/**
当接收者为指针接收者时，会将对象的地址传递到方法在，多用于更新数据
*/
func (this *tree) setValue(value int) {
	this.value = value
}

/**
中序遍历
*/
func (this *tree) traverse() {
	if this == nil {
		return
	}

	this.left.traverse()
	this.print()
	this.right.traverse()
}

/**
根据切片生成二叉树
*/
func createTree(data []int) *tree {
	trees := make([]tree, len(data))
	//装箱进入节点数组
	for index, value := range data {
		tree := new(tree)
		tree.value = value
		tree.left = nil
		tree.right = nil
		trees[index] = *tree
	}
	//判断节点数组内容是否为null
	if len(trees) > 0 {
		//i代表的是根节点的索引，从0开始
		for i := 0; i < len(data)/2-1; i++ {
			//计算根节点在数组中的下一个元素是否为nil 如果不是则加入根节点的左节点
			if &trees[2*i+1] != nil {
				trees[i].left = &trees[2*i+1]
			}
			//计算根节点在数组中的下第二个元素是否为nil 如果不是则加入根节点的右节点
			if &trees[2*i+2] != nil {
				trees[i].right = &trees[2*i+2]
			}
		}
	}
	//判断当前树的最后一个根节点 因为其可能没有右节点，所以单独处理
	index := len(data)/2 - 1
	trees[index].left = &trees[2*index+1]
	//在数组为奇数时有右节点
	if len(data)%2 == 1 {
		trees[index].right = &trees[2*index+2]
	}
	return &trees[0]
}

func (this *tree) append(node tree) {

}

func main() {
	var root tree

	root = tree{value: 3}
	root.left = &tree{}
	root.right = &tree{5, nil, nil}
	root.right.left = new(tree)

	//root.traverse()

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	newRoot := createTree(data)
	fmt.Println("root--->", newRoot)
	fmt.Println("newRoot.left.right.value--->", newRoot.left.right.value)
	fmt.Println("newRoot.right.left.value--->", newRoot.right.left.value)
	newRoot.append(tree{10, nil, nil})
	newRoot.traverse()
}
