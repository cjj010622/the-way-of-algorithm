package day24

/*
序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。

设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序列化为最初的二叉搜索树。

编码的字符串应尽可能紧凑。
*/

type Codec struct {
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root==nil{
        return ""
    }

    data:=strings.Builder{}
    var dfs func(root *TreeNode)
    dfs=func(root *TreeNode){
        if root==nil{
            return
        }

        data.WriteString(strconv.Itoa(root.Val))
        data.WriteByte(' ')
        dfs(root.Left)
        dfs(root.Right)
    }

    dfs(root)
    return data.String()[0:data.Len()-1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    if data==""{
        return nil
    }

    vals:=strings.Split(data," ")
    i:=0
    var dfs func(int,int) *TreeNode
    dfs=func(mi,mx int) *TreeNode{
        if i==len(vals){
            return nil
        }

        x,_:=strconv.Atoi(vals[i])
        if x<mi||x>mx{
            return nil
        }

        i++
        root:=&TreeNode{Val:x}
        root.Left=dfs(mi,x)
        root.Right=dfs(x,mx)
        return root
    }
    return dfs(math.MinInt64,math.MaxInt64)
}