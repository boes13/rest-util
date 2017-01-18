package jsonapi

import (
	"fmt"
	"testing"
)

type meta struct {
	field1 string
	field2 string
}

type data struct {
	id      string
	name    string
	address string
	phone   string
}

func TestCreateDataResponse(t *testing.T) {
	dataResponse := CreateDataResponse()
	if dataResponse == nil {
		t.Error("Expected DataResponse object, got nil!")
	}

	err := dataResponse.AddMeta(meta{"field1 value", "field2 value"})
	if err != nil {
		t.Error("Expected no error, got error!")
	}

	mydata := data{
		id:      "my_id",
		address: "my_address",
		name:    "my_name",
		phone:   "my_phone",
	}
	err = dataResponse.SetData(mydata)
	if err != nil {
		t.Error("Expected no error, got error!")
	}

	mydata2 := data{
		id:      "my_id2",
		address: "my_address2",
		name:    "my_name2",
		phone:   "my_phone2",
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

	dataResponse.SetLinks("www.self.com", "www.related.com", "www.first.com", "www.last.com", "www.prev.com", "www.next.com")

	fmt.Printf("%+v", dataResponse)
}
