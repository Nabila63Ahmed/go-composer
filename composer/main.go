package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func main() {
	path := "../data.csv"
	trainingSongs := loadData(path)
	results := make(chan string)

	orderedStatisticsMap, cumulativeArrayMap := constructStats(trainingSongs)

	// songCount := 5

	length1 := 10
	currentChar1 := "a"

	length2 := 100
	currentChar2 := "a"

	length3 := 1000
	currentChar3 := "a"

	go composeRoutines(currentChar1, length1, orderedStatisticsMap, cumulativeArrayMap, results)
	go composeRoutines(currentChar2, length2, orderedStatisticsMap, cumulativeArrayMap, results)
	go composeRoutines(currentChar3, length3, orderedStatisticsMap, cumulativeArrayMap, results)

	var songs []string
	for i := 0; i < 3; i++ {
		songs = append(songs, <-results)
	}

	for _, song := range songs {
		fmt.Println(song)
	}
}

func loadData(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Cannot open file: ", err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Cannot read file: ", err)
	}

	trainingSongs := make([]string, len(records)-1)
	for i := 0; i < len(trainingSongs); i++ {
		trainingSongs[i] = records[i+1][1]
	}
	return trainingSongs
}

func constructStats(trainingSongs []string) (map[string]PairList, map[string][]float64) {
	statisticsMap := make(map[string]map[string]int)
	orderedStatisticsMap := make(map[string]PairList)
	cumulativeArrayMap := make(map[string][]float64)

	for i := 0; i < len(trainingSongs); i++ {
		for j := 0; j < len(trainingSongs[i]); j++ {
			var currentChar = string(trainingSongs[i][j])
			for k := 1; k < len(trainingSongs[i])-1; k++ {
				var nextChar = string(trainingSongs[i][k])
				_, check_map_exsists := statisticsMap[currentChar]

				if !check_map_exsists {
					statisticsMap[currentChar] = map[string]int{
						nextChar: 0,
					}
				}

				internalVal, check_internal_map := statisticsMap[currentChar][nextChar]
				if !check_internal_map {
					internalVal = 0
				}
				statisticsMap[currentChar][nextChar] = internalVal + 1

			}

		}

	}

	for key, currentMap := range statisticsMap {

		var currentSum = 0
		p := make(PairList, len(currentMap))
		a := make([]float64, len(p))

		i := 0
		for k, v := range currentMap {
			currentSum += v
			p[i] = Pair{k, v}
			i++
		}

		sort.Sort(p)

		for k := 0; k < len(p); k++ {
			var val = p[k].Value
			var newVal = float64(val) * 100 / float64(currentSum)
			if a[k] = newVal; k != 0 {
				a[k] = newVal + a[k-1]
			}
		}

		orderedStatisticsMap[key] = p
		cumulativeArrayMap[key] = a

	}
	return orderedStatisticsMap, cumulativeArrayMap
}

func composeRoutines(currentChar string, length int, orderedStatisticsMap map[string]PairList, cumulativeArrayMap map[string][]float64, results chan string) {
	song := compose(currentChar, length, orderedStatisticsMap, cumulativeArrayMap)
	results <- song
}

func compose(currentChar string, length int, orderedStatisticsMap map[string]PairList, cumulativeArrayMap map[string][]float64) string {
	var currentMap = cumulativeArrayMap[currentChar]
	var song = currentChar
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length-1; i++ {
		var randomNr = 100.0 * rand.Float64()

		for j := 0; j < len(currentMap); j++ {
			if randomNr <= currentMap[j] {
				var newCurrent = orderedStatisticsMap[currentChar][j].Key
				currentChar = newCurrent
				song += newCurrent
				currentMap = cumulativeArrayMap[currentChar]
				break
			}
		}
	}
	return song
}
