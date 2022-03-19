
选择排序算法
1. 选择排序算法
算法描述：从未排序数据中选择最大或者最小的值和当前值交换 O(n^2).

算法步骤
选择一个数当最小值或者最大值，进行比较然后交换
循环向后查进行
package sort

import "fmt"

//获取切片里面的最大值
func SelectMax(arr []int) int {
    length := len(arr)
    if length <= 1 {
        return arr[0]
    }
    max := arr[0]
    for i := 1; i < length; i++ {
        if arr[i] > max {
            max = arr[i]
        }
    }
    return max
}

//切片排序
func SelectSort(arr []int) []int {
    length := len(arr)
    if length <= 1 {
        return arr
    }
    for i := 1; i < length; i++ {
        min := i
        for j := i + 1; j < length; j++ {
            if arr[min] > arr[j] {
                min = j
            }
        }
        if i != min {
            arr[i], arr[min] = arr[min], arr[i]
        }
    }
    return arr
}

//选择排序
func main() {
    arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
    max := SelectMax(arr)
    selectsort := SelectSort(arr)
    fmt.Println(max)
    fmt.Println(selectsort)
}
赏
