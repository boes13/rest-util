package jsonapi

import (
	"fmt"
	"testing"
)

type any struct {
	id string
	name string
}

func TestCreateErrorResponse(t *testing.T) {
	errResp := CreateErrorResponse()
	if errResp == nil {
		t.Error("Expected ErrorResponse object created, got nil!")
		return
	}

	urls := []string {"http://google.com", "http://www.yahoo.com"}
	links := CreateErrorLinks(urls)
	if len(links) != 2 {
		t.Errorf("Expected 2 error links, got %d!", len(links))
	}

	errSource := CreateErrorSource("/path/to/error", "user_id")
	if errSource == nil {
		t.Error("Expected ErrorSource object, got nil!")
	}

	err := errResp.AddError("myid", links, "mystatus", "code", "title", "detail", errSource, nil)
	if err != nil {
		t.Error("Found err:", err)
		return
	}

	err = errResp.AddError("myid2", links, "mystatus2", "code2", "title2", "detail2", errSource, any{"meta_id2", "meta_name2"})
	if err != nil {
		t.Error("Found err:", err)
		return
	}
	fmt.Printf("%+v\n", errResp)
}

