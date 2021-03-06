package helper

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"go-stock/app/global/structs"
)

func PrintLog(param interface{}) {
	byteData, _ := jsoniter.Marshal(param)
	show := string(byteData)
	fmt.Println("PrintLog:" + show)
	panic(show)
}

func StructToMap(myStruct interface{}) (myMap map[string]interface{}, err error) {
	myMap = make(map[string]interface{})

	byteData, err := jsoniter.Marshal(myStruct)
	if err != nil {
		return
	}

	if err = jsoniter.Unmarshal(byteData, &myMap); err != nil {
		return
	}

	return
}

func ResponseFail(err error) (res *structs.APIResponse) {

	return
}
