package main

import "fmt"


func merge(left, right []int) []int {
    ret := make([]int, len(left) + len(right))
    i, j, k := 0, 0, 0
    for ; i < len(left) && j < len(right); k++ {
        if left[i] < right[j] {
            ret[k] = left[i]
            i++
        } else {
            ret[k] = right[j]
            j++
        }
    }
    for ; i < len(left); i++ {
        ret[k] = left[i]
        k++
    }
    for ; j < len(right); j++ {
        ret[k] = right[j]
        k++
    }
    return ret
}


func mergesort(a []int) {
    if len(a) <= 1 {
        return
    } 
    var m int = len(a) / 2
    mergesort(a[:m])
    mergesort(a[m:])
    copy(a, merge(a[:m], a[m:]))
}

func main() {
    a := []int {4, 6, 1, 3, 5, 0, 1, 2}
    mergesort(a)
    fmt.Println(a)
}
