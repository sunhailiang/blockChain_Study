package main

import "testing"

func TestSave(t *testing.T) {
	per := person{
       Name:"harry",
       Age:18,
       Gender:"男",
	}
	err:=per.Save()
	if err!=nil{
         t.Fatalf("save preson failed,err：%v",err)
	}
}

func TestLoad(t *testing.T)  {

}
