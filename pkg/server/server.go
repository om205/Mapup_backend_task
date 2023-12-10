package server

import (
	"encoding/json"
	"net/http"
	"sort"
	"sync"
	"time"
)

type RequestPayload struct {
	ToSort [][]int `json:"to_sort"`
}

type ResponsePayload struct {
	SortedArrays [][]int `json:"sorted_arrays"`
	TimeNs       int64   `json:"time_ns"`
}

func sortArraysSequential(arrays [][]int) ([][]int, int64) {
	startTime := time.Now()

	var sortedArrays [][]int
	for _, arr := range arrays {
		sortedArray := make([]int, len(arr))
		copy(sortedArray, arr)
		sort.Ints(sortedArray)
		sortedArrays = append(sortedArrays, sortedArray)
	}

	endTime := time.Now()
	timeTaken := endTime.Sub(startTime).Nanoseconds()

	return sortedArrays, timeTaken
}

func sortArraysConcurrent(arrays [][]int) ([][]int, int64) {
	startTime := time.Now()

	var sortedArrays [][]int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, arr := range arrays {
		wg.Add(1)
		go func(arr []int) {
			defer wg.Done()

			sortedArray := make([]int, len(arr))
			copy(sortedArray, arr)
			sort.Ints(sortedArray)

			mutex.Lock()
			sortedArrays = append(sortedArrays, sortedArray)
			mutex.Unlock()
		}(arr)
	}

	wg.Wait()

	endTime := time.Now()
	timeTaken := endTime.Sub(startTime).Nanoseconds()

	return sortedArrays, timeTaken
}

func writeJSONResponse(w http.ResponseWriter, payload ResponsePayload) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}
