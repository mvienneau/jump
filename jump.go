package main

import (
	"fmt"
	"encoding/json"
	"sync"
	"time"
)

// A struct to help shape the global average data
type AverageData struct{
	Avg float64
	Cnt int
}

// The struct to read actions into
type Message struct {
	Action string	`json:"action"`
	Time int		`json:"time"`
}

// Struct to shape the return body
type Return struct {
	Action string	`json:"action"`
	Average float64	`json:"average"`
}

var mutex = sync.RWMutex{}

// Assumption: Using a global map here keyed off action to hold the average data.
// this could easily be swapped out with a persistent key value store. 
var globalData = map[string]AverageData{}

func addAction(s string) error {
	// Convert the string into a Struct
	json_obj := Message{}
	err := json.Unmarshal([]byte(s), &json_obj)
	if err != nil{
		return err
	}

	// A lot of reading/writing is happening here on the global map
	// Throwing a mutex around this access will garuntee that when we update the average
	// it won't be thrown off by a new value coming in and updating it mid-computation
	mutex.Lock()
	defer mutex.Unlock()
	// Get the current average for the action type
	curAvg := globalData[json_obj.Action]
	// If there isn't a current average...
	if (AverageData{}) == curAvg {
		new_avg := &AverageData{
			Avg: float64(json_obj.Time),
			Cnt: 1,
		}
		globalData[json_obj.Action] = *new_avg
		return nil

	}
	// If there is an average, compute the new avg via: avg_new = avg_old + (value_new - avg_old / samples_new)
	new_avg := &AverageData{
		Avg: float64(curAvg.Avg + ((float64(json_obj.Time) - curAvg.Avg) / float64((curAvg.Cnt + 1)))),
		Cnt: curAvg.Cnt + 1,
	}
	globalData[json_obj.Action] = *new_avg
	return nil

}

func getStats() string {
	// No data
	if len(globalData) == 0{
		return "[]"
	}
	var a []Return
	// Transform our "database" into the correct structure
	for action, value := range globalData{
		a = append(a, Return{
				Action: action,
				Average: value.Avg,
			})
	}
	b, err := json.Marshal(a)
	if err != nil{
		return "An error occured"
	}
	return string(b)
}

// I left this function in to help see how it behaves very easily with a go run
// Just adds some actions in go routines, sleeps for 5 seconds then outputs the results
func main() {
	go addAction(`{"action": "run", "time": 350}`)
	go addAction(`{"action": "jump", "time": 100}`)
	go addAction(`{"action": "jump", "time": 200}`)
	go addAction(`{"action": "run", "time": 150}`)
	go addAction(`{"action": "run", "time": 250}`)
	time.Sleep(time.Second * 5)
	stats := getStats()
	fmt.Println(stats)
}