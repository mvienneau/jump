package main

import (
	"testing"
	"reflect"
)

func TestAddAction(t *testing.T) {
	t.Log("Testing addAction for 3 Items, 2 of the same type")
	err := addAction(`{"action": "run", "time": 350}`)
	if err != nil{
		t.Errorf("Unexpected Error in addAction", err)
	}
	g_data := map[string]AverageData{}
	g_data["run"] = AverageData{
		Avg: 350,
		Cnt: 1,
	}
	eq := reflect.DeepEqual(g_data, globalData)
	if !eq{
		t.Fail()
	}

	err = addAction(`{"action": "run", "time": 250}`)
	if err != nil{
		t.Errorf("Unexpected Error in addAction", err)
	}

	g_data["run"] = AverageData{
		Avg: 300,
		Cnt: 2,
	}
	eq = reflect.DeepEqual(g_data, globalData)
	if !eq{
		t.Fail()
	}

	err = addAction(`{"action": "jump", "time": 50}`)
	if err != nil{
		t.Errorf("Unexpected Error in addAction", err)
	}

	g_data["jump"] = AverageData{
		Avg: 50,
		Cnt: 1,
	}
	eq = reflect.DeepEqual(g_data, globalData)
	if !eq{
		t.Fail()
	}
}

func TestBadData(t *testing.T) {
	t.Log("Testing Error on Bad Data")

	err := addAction(`{"action": "jump", "time": "Not a valid time"}`)
	if err == nil{
		t.Errorf("Expected to error on bad data, did not")
	}

	err = addAction(`this isn't valid data either`)
	if err == nil{
		t.Errorf("Expected to error on bad data, did not")
	}
}

func TestGetStats(t *testing.T) {
	t.Log("Testing get stats")
	// Empty our data store
	globalData = map[string]AverageData{}
	s := getStats()
	// Test for empty data
	if s != "[]"{
		t.Errorf("Expected empty array, got ", s)
	}

	// Add something to data store
	globalData["run"] = AverageData{
		Avg: 350,
		Cnt: 1,
	}
	s = getStats()
	// Test for correct shape and numbers
	if s != `[{"action":"run","average":350}]`{
		t.Errorf("getStats did not transform the data correctly")
	}
	


}