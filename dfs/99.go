package dfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 99. 恢复二叉搜索树
func recoverTree(root *TreeNode) {
	// 第一种方法：中序遍历，然后找到两个逆序对，交换两个逆序对的第一个和第二个的值
	var nums []int
	var inorder func(node *TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)

	x, y := findTwoSwapped(nums)
	fix(root, 2, x, y)
}

func findTwoSwapped(nums []int) (int, int) {
	index1, index2 := -1, -1
	n := len(nums)
	// 要么有两个 要么只有一个
	// 两个的情况，找i和j+1
	// 一个的情况，找i和i+1
	for i := 0; i < n-1; i++ {
		if nums[i+1] < nums[i] {
			index2 = i + 1 //再赋值
			if index1 == -1 {
				index1 = i // 一个情况
			} else {
				break
			}
		}
	}
	return nums[index1], nums[index2]
}

func fix(root *TreeNode, count, x, y int) {
	if root == nil {
		return
	}

	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else {
			root.Val = x
		}
		count--
		if count == 0 {
			return
		}
	}
	fix(root.Left, count, x, y)
	fix(root.Right, count, x, y)
}

func recoverTree2(root *TreeNode)  {
    var x, y, pred, predecessor *TreeNode

    for root != nil {
        if root.Left != nil {
            // predecessor 节点就是当前 root 节点向左走一步，然后一直向右走至无法走为止
            predecessor = root.Left
            for predecessor.Right != nil && predecessor.Right != root {
                predecessor = predecessor.Right
            }

            // 让 predecessor 的右指针指向 root，继续遍历左子树
            if predecessor.Right == nil {
                predecessor.Right = root
                root = root.Left
            } else { // 说明左子树已经访问完了，我们需要断开链接
                if pred != nil && root.Val < pred.Val {
                    y = root
                    if x == nil {
                        x = pred
                    }
                }
                pred = root
                predecessor.Right = nil
                root = root.Right
            }
        } else { // 如果没有左孩子，则直接访问右孩子
            if pred != nil && root.Val < pred.Val {
                y = root
                if x == nil {
                    x = pred
                }
            }
            pred = root
            root = root.Right
        }
    }
    x.Val, y.Val = y.Val, x.Val
}
