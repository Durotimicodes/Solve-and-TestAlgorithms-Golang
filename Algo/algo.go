package algorithms

import (
	"fmt"
	"math"
	"strconv"
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
		} //if u have more than one key than has the same hightest value delete the ones with the lower key
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

// Todo
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

// Todo
func CountingValleys(steps int32, path string) {
	// Write your code here

	var nStr int

	fmt.Println(path)
	for i := 0; i < int(steps); i++ {
		if string(path[i]) == "U" {
			nStr += 1
		} else if string(path[i]) == "D" {
			nStr += -1
		}
	}

	fStr := 0 + nStr + 0
	fmt.Println(fStr)

}

func FindDigits(n int32) int32 {
	// Write your code here

	//convt num to string
	ns := strconv.Itoa(int(n))

	//split the numbers
	spStr := strings.Split(ns, ", ")
	jn := strings.Join(spStr, ",")

	count := int32(0)

	//iterate over the numb and count how many is divisible by the original num
	for i := 0; i < len(jn); i++ {
		sn, _ := strconv.Atoi(string(jn[i]))
		ln, _ := strconv.Atoi(string(jn[len(jn)-1]))

		if sn != 0 {
			if n%int32(sn) == 0 {
				count++
				fmt.Println("not equal to zero", count)
			}
		} else if sn == 0 && sn != ln {
			sn++
		}

	}

	fmt.Println(count)
	return count

}

// sherlock and square
func Squares(a int32, b int32) int32 {
	// Write your code here

	//declare and initialize a count variable
	count := int32(0)

	//generate the range by incrementing the value of "a" and stop iteration @ "b"
	for a := a; a <= b; a++ {
		q := math.Sqrt(float64(a))
		if q == float64(int64(q)) {
			count++
		}
	}

	return count
}

func KandD(k []int, d []int) []int {
	//Write your code here

	result := []int{}

	i := 0
	for i < len(k) {
		for j := 0; j < len(d); j++ {
			result = append(result, k[i]+d[j])
		}
		i++
	}

	fmt.Println(result)
	return result
}


func GetMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	/*
	 * Write your code here.
	 */

	//check for the max length btw the two products
	klen := len(keyboards)
	dlen := len(drives)

	var prod_1 []int32
	if klen < dlen {
		x := 0
		for x < klen {
			for y := 0; y < dlen; y++ {
				prod_1 = append(prod_1, keyboards[x]+drives[y])
			}
			x++
		}
	} else if klen > dlen {
		k := 0
		for k < dlen {
			for j := 0; j < klen; j++ {
				prod_1 = append(prod_1, drives[k]+keyboards[j])
			}
			k++
		}
	} else {
		c := 0
		for c < klen {
			for d := 0; d < klen; d++ {
				prod_1 = append(prod_1, keyboards[c]+drives[d])
			}
			c++
		}
	}

	moneySpent := int32(0)

	for i := 0; i < len(prod_1); i++ {
		if prod_1[i] < b || prod_1[i] == b {
			moneySpent = prod_1[i]
		} else if prod_1[i] > b {
			moneySpent = -1
		}
	}

	fmt.Println(moneySpent)


	return moneySpent

}
