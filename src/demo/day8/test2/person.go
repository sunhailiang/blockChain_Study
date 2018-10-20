package main

import (
	"encoding/json"
	"io/ioutil"
)

type person struct {
	Name string
	Gender string
	Age int
}
func (p *person)Save()(err error){
	 data,err:=json.Marshal(p)
	 if err!=nil{
	 	return
	 }
	 err=ioutil.WriteFile("d:/person.dat",data,0755)
	 return 
}
func (p *person)Load(err error)  {
	data,err:=ioutil.ReadFile("d:/person.dat")
	if err!=nil{
		return
	}
	err=json.Unmarshal(data,p)
	return
}