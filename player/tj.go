package player

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

func tj(ip string) int {
	var timeout time.Duration = 30
	var oldMap = make(map[string]int64, 50)
	
	f, _ := os.Open("count.gob")
	d := gob.NewDecoder(f)
	d.Decode(&oldMap)
	var newMap = make(map[string]int64, 50)
	t := time.Now()
	for key, value := range oldMap {
		if key != ip && value > t.Unix() {
			newMap[key] = value
		}
	}
	newMap[ip] = t.Add(timeout * time.Second).Unix()

	f, err := os.OpenFile("count.gob", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	enc := gob.NewEncoder(f)
	if err := enc.Encode(newMap); err != nil {
		fmt.Println(err)
	}
	return len(newMap)
}
