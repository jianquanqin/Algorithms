package main

import (
	"algorithms_examples/linearList/src"
	"fmt"
	"time"
)

func main() {
	linearList := src.LinearList{} //直接初始化，可以直接使用默认值
	dynamicLinearList := src.DynamicLinearList{}

	start := time.Now()
	dynamicLinearList.InsertDynamicLinearList(1, 20)
	dynamicLinearList.InsertDynamicLinearList(2, 21)
	dynamicLinearList.InsertDynamicLinearList(3, 21)
	dynamicLinearList.InsertDynamicLinearList(4, 21)
	dynamicLinearList.InsertDynamicLinearList(5, 21)

	fmt.Println(time.Since(start))

	fmt.Println(linearList, dynamicLinearList)
}
