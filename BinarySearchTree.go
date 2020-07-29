package main

type BinarySearchTree struct {
	Data  int
	Left  *BinarySearchTree
	Right *BinarySearchTree
}

func (b *BinarySearchTree) search(num int) *BinarySearchTree {

	if b == nil {
		return nil
	}
	for b != nil {
		if b.Data == num {
			return b
		}
		if b.Data > num {
			b = b.Left
		} else {
			b = b.Right
		}

	}
	return nil
}

/**
二叉树 插入，考虑数据重复的情况，把数据重复的情况当做插入值比当前值 大  处理，插入right测
*/
func (b *BinarySearchTree) insert(num int) *BinarySearchTree {
	if b == nil {
		return &BinarySearchTree{Data: num}
	}
	for b != nil {
		if b.Data > num {
			if b.Left == nil {
				b.Left = &BinarySearchTree{Data: num}
				break
			}
			b = b.Left
		} else {
			if b.Right == nil {
				b.Right = &BinarySearchTree{Data: num}
				break
			}
			b = b.Right
		}
	}
	return b
}

/**
二叉树 删除，考虑元素相同的情况， 前面插入的时候，会把相同的元素插入到 right 测，所以删除的时候，也需要对right测进行删除
树节点的删除需要考虑一下情况：
1、如果节点没有子节点，那么只需要 把其父节点对该节点的引用 设置为 nil
2、
*/
func (b *BinarySearchTree) delete(num int) {

}

func main() {

}
