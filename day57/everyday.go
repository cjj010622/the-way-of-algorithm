package day57

/*
给定一个二叉树：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL 。

初始状态下，所有 next 指针都被设置为 NULL 。
*/

func connect(root *Node) *Node {
	if root==nil{
		return root
	}
	q:=[]*Node{root}
	for len(q)>0{
		var p *Node
		for n:=len(q);n>0;n--{
			node:=q[0]
			q=q[1:]
			if p!=nil{
				p.Next=node
			}
			p=node
			if node.Left!=nil{
				q=append(q,node.Left)
			}
			if node.Right!=nil{
				q=append(q,node.Right)
			}
		}
	}
	return root
}