package jsonapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

type meta struct {
	Field1 string
	Field2 string
}

type data struct {
	Id      string
	Name    string
	Address string
	Phone   string
}

func TestCreateDataResponse(t *testing.T) {
	dataResponse := CreateDataResponse()
	if dataResponse == nil {
		t.Error("Expected DataResponse object, got nil!")
	}

	err := dataResponse.SetMeta(meta{"field1 value", "field2 value"})
	if err != nil {
		t.Error("Expected no error, got error!")
	}

	mydata := data{
		Id:      "my_id",
		Address: "my_address",
		Name:    "my_name",
		Phone:   "my_phone",
	}
	err = dataResponse.SetData(mydata)
	if err != nil {
		t.Error("Expected no error, got error!")
	}

	mydata2 := data{
		Id:      "my_id2",
		Address: "my_address2",
		Name:    "my_name2",
		Phone:   "my_phone2",
	}
	dataSlice := make([]data, 0)
	dataSlice = append(dataSlice, mydata)
	dataSlice = append(dataSlice, mydata2)
	err = dataResponse.SetData(dataSlice)
	if err != nil {
		t.Error("Expected no error, got error!")
	}

	var dataArray [2]data
	dataArray[0] = mydata
	dataArray[1] = mydata2
	err = dataResponse.SetData(dataArray)
	if err != nil {
		t.Error("Expected no error, got error!")
	}

	var someMap = make(map[int]int)
	someMap[100] = 10
	someMap[12] = 12
	err = dataResponse.SetData(someMap)
	if err != nil {
		t.Error("Expected no error, got error!")
	}

	dataResponse.SetLinks("www.self.com", "www.related.com", "www.first.com", "www.last.com", "www.prev.com", "www.next.com")

	buf, err := json.Marshal(dataResponse)
	if err != nil {
		t.Error("Expected no error, got error!")
	}

	fmt.Printf("%s", string(buf))
}
