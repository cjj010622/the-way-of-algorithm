package day50

/*
矩形蛋糕的高度为 h 且宽度为 w，给你两个整数数组 horizontalCuts 和 verticalCuts，其中：

 horizontalCuts[i] 是从矩形蛋糕顶部到第  i 个水平切口的距离
verticalCuts[j] 是从矩形蛋糕的左侧到第 j 个竖直切口的距离
请你按数组 horizontalCuts 和 verticalCuts 中提供的水平和竖直位置切割后，请你找出 面积最大 的那份蛋糕，并返回其 面积 。由于答案可能是一个很大的数字，因此需要将结果 对 109 + 7 取余 后返回。
*/


func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
    horizontalCuts=append(horizontalCuts,[]int{0,h}...)
    verticalCuts=append(verticalCuts,[]int{0,w}...)
    sort.Ints(horizontalCuts)
    sort.Ints(verticalCuts)
    x,y:=0,0
    const mod int=1e9+7
    for i:=1;i<len(horizontalCuts);i++{
        x=max(x,horizontalCuts[i]-horizontalCuts[i-1])
    }
    for i:=1;i<len(verticalCuts);i++{
        y=max(y,verticalCuts[i]-verticalCuts[i-1])
    }
    return (x*y)%mod
}

func max(a,b int) int{
    if a>b{
        return a
    }

    return b
}