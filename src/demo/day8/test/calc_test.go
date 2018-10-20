package main

import "testing"
//必须要以Test开头
func TestAdd(t *testing.T) {
	//给你要测试的函数放入参数
	r := add(2, 4)
	//根据返回结果手动抛出代码异常
if r != 6 {
t.Fatalf("add(2,4) err")
}
t.Logf("test add succ")
}
