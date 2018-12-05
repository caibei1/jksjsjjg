package main

import "fmt"

//快速排序
func main()  {
	//切片是引用传递
	arr := []int{11,12,3,4,5,7,7,8,19,4}
	//注意是len-1
	quckSort(arr,0,len(arr)-1)
	fmt.Println(arr)
}

func quckSort(arr []int,p,q int)  {
	//当q和p一样的时候  或者p>q  说明排序已经完成
	if p >= q {
		return
	}
	//node是参照点最终的位置
	node := step(arr,p,q)
	//比参照点小的继续排序  比参照点大的继续排序
	quckSort(arr,p,node-1)
	quckSort(arr,node+1,q)
}

func step(arr []int,p,q int) int {
	//把q作为参考点  j是已经排序的下一个位置
	j := p
	for i := p; i < q; i++{
		if arr[i] <= arr[q]{
			arr[j],arr[i] = arr[i],arr[j]
			j++
		}
	}
	//最后把参考点的位置交换
	arr[q],arr[j] = arr[j],arr[q]
	return j
}

