package algorithms

import (
	"fmt"
	"strings"
)

// Hackerrank Equalizer Algorithm
func EqualizeArray(arr []int32) int32 {
	// Write your code here

	//make a map of int32
	m := make(map[int32]int32)

	//iterate of the slice of integers and do a tally to generate its value
	for _, value := range arr {
		if val, exist := m[value]; exist {
			m[value] = val + int32(1)
		} else {
			m[value] = 1
		}
	}

	//variable declaration
	var maxVal, maxKey, del, sameVal int32

	//determine the max val in map
	for _, num := range m {
		if num > maxVal {
			maxVal = num
		}
	}

	//dataStore 
	var dataStore []int32

	

	//find no/list of keys with the same value and count them
	for key, val := range m {
		if val == maxVal {
			dataStore = append(dataStore, key)
			sameVal++
		}
	}

	//delete the lower key
	for key, numbs := range m {
		if numbs < maxVal {
			del += numbs
		}//if u have more than one key than has the same hightest value delete the ones with the lower key
		if sameVal > 1 {
			//find the hightest key within them
			for i := 0; i < len(dataStore); i++ {
				if dataStore[i] > maxKey {
					maxKey = dataStore[i]
				}
			}
			//if the value is same with the highest value and the key is less than the largest key count
			if numbs == maxVal && key < maxKey {
				del += numbs
			}

		}

	}
	
	return del

}

func CheckGemStone(n []string) int32 {

	//1.join strings
	strJoin := strings.Join(n, "")

	//2.slipt the strings into single alphabets
	strSplit := strings.Split(strJoin, "")
	strJn2 := strings.Join(strSplit, ", ")
	strSplit2 := strings.Split(strJn2, ", ")

	//3.iterate over the slice of strings and check if an alphabet is contains

	m := make(map[string]int)

	fmt.Println(m)
	//take each string and cross-check with letters
	for j := 0; j < len(n); j++ {
		for i := 0; i < len(strSplit2); i++ {
			if strings.Contains(n[j], strSplit[i]) {
				m[strSplit[i]]++
			}
		}
	}

	max := int(0)
	count := int(0)

	for _, val := range m {
		if val > max {
			max = val
			count++
		}

	}

	fmt.Println(count)

	return int32(count)

}
