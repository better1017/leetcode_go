package one

// 206. 反转链表

// ListNode 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

// 示例 1：
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]

// 示例 2：
//输入：head = [1,2]
//输出：[2,1]

// 示例 3：
//输入：head = []
//输出：[]

// 提示：
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000

// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode { //206.反转链表 - 迭代
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
func reverseList2(head *ListNode) *ListNode { //206.反转链表 - 递归
	if head == nil || head.Next == nil {
		return head
	}

	root := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return root
}
