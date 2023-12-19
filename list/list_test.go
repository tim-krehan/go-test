package list

import "testing"

func TestAdd(t *testing.T){
	add_error := Add("test")
	if add_error != nil {
		t.Fatalf("%v", add_error)
	}
}