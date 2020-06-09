package controllers

import (
	"reflect"
	"testing"
)

func TestCreateDataMap(t *testing.T) {
	expected := make(map[string]string)
	expected["action"] = "favorite"
	expected["placeId"] = "1234"

	q := "action=favorite&placeId=1234"
	actual := createDataMap(q)

	if !reflect.DeepEqual(expected, actual) {
		t.Error("\n理想： ", expected, "\n実態： ", actual)
	}

	t.Log("TestCreateDataMap Finished!")
}
