package public

import "bytes"

func FillLastGroup(sourceData []byte, Size int) []byte {
	setLen := Size - len(sourceData)%Size
	setVal := setLen
	setRes := bytes.Repeat([]byte{byte(setVal)}, setLen)
	sourceData = append(sourceData, setRes...)
	return sourceData
}

func CutFillData(sourceData []byte) []byte {
	fillLen := sourceData[len(sourceData)-1]
	resData := sourceData[:(len(sourceData) - int(fillLen))]
	return  resData
}
