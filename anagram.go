package main

import (
	"fmt"
	"reflect"
)

func main(){
	// Good Pattern to compare multiple inputs values. 
	// The function that accepts two slices should return TRUE if every value in the slice
	// has correspend square value in the second slice. The frequence value must be the same. 
	var arg1 []int = []int{1,2,3,2}
	var arg2 []int = []int{9,1,4,4}
	ok := ValidFrequencyCounter(arg1, arg2)
	fmt.Println("The numbers on arg2 whould be square of arg1\n")
	fmt.Println(ok)
	fmt.Println("\n")
	//---------------------------------------
	var arg3 = "marcio"
	var arg4 = "mcoari"
	ok = ValidAnagram(arg3, arg4)
	fmt.Println("The letters in arg3 must contain in arg4\n")
	fmt.Println(ok)
	fmt.Println("\n")
	//---------------------------------------
	var arg = []int{1,2,3,3,3,4,5,6,6,7}
	ret := CountUniqueValues1(arg)

	fmt.Println("It identify and return total number of unique values\n")
	fmt.Printf("Total unique values = %d\n", ret)
	//----------------------------------------

/*	var n = []int{2,6,9,2,1,8,5,6,3}
	var num = 3
	ret := maxSubSliceSum(n,num)
	fmt.Printf("The max values between give numbers = %d\n", ret)*/
    //-----------------------------------------
    var s = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19}
    var i = 14
    retValue := search(s,i)
	fmt.Printf("Number %d found = %v\n",i, retValue)

	//-----------------------------------------
	// Calculating Fibonnaci number
	fibonnaci(10)

}

func ValidFrequencyCounter(arg1 []int , arg2[]int) bool{

	if len(arg1) != len(arg1){
		return false }

	var frequencyCount1 = make(map[int]int)
	var frequencyCount2 = make(map[int]int)
	
	// Build the first map aggregating the total # per each value in the slice "frequencyCount1" 
	for _, v := range arg1 {
		if _, ok := frequencyCount1[v]; ok {
			frequencyCount1[v] += 1
		} else {
			frequencyCount1[v] = 1
		}
	}

	// Build the second map aggregating the total # per each value in the slice "frequencyCount2" 
	for _, v := range arg2 {
		if _, ok := frequencyCount2[v]; ok {
			frequencyCount2[v] += 1
		} else {
			frequencyCount2[v] = 1
		}
	}

	fmt.Println(frequencyCount1)
	fmt.Println(frequencyCount2)

	// For each value in the first slice, make sure a square value is found on the second map "frequencyCount2"
	for v , _ := range frequencyCount1 { // v = map[key], the _ would be the map[]value.

		if _, ok := frequencyCount2[v * v]; !ok{
			return false
		}
		if frequencyCount1[v] != frequencyCount2[v * v]{
			return false
		}
	}

	return true
}


func ValidAnagram(arg1 string, arg2 string) bool {

	if len(arg1) != len(arg2){
		return false
	}
	var val1 = make(map[int]int)
	var val2 = make(map[int]int)
	var s1 = []rune(arg1)
	var s2 = []rune(arg2)

	fmt.Println("s1 = ",s1)
	fmt.Println("s2 = ",s2)
	fmt.Println("--------------------")

	for i , v := range s1 {
		if _ , ok := val1[int(v)] ; ok {
			val1[int(v)] +=1
		}else {
			val1[int(v)] =1
		}
		fmt.Printf("s1 value in loop %v is  = %d \n", i ,v)

	}
	for _, v := range s2 {
		if _ , ok := val2[int(v)] ; ok {
			val2[int(v)] +=1
		}else {
			val2[int(v)] =1
		}
	}

	fmt.Println("val1 = ",val1)
	fmt.Println("val2 = ",val2)

	if ok := reflect.DeepEqual(val1,val2); !ok { // compare two maps for equality
		return false
	}

	return true
}

// Remove a single element from slice in constant time.
func remove(s []int, i int) []int {

	s[i] = s[len(s)-1]  // Copy last element to index i.
	s[len(s)-1] = 0     // Erase last element (write zero value).
	return s[:len(s)-1] // Truncate slice.
}


func CountUniqueValues(arg []int) int {
	// arg contains ordered int values, [1,1,1,2,3] should return 3

	if len(arg) == 0 {
		return 0
	}
	var freq = make(map[int]int)
	for _ , v := range arg{
		if _ , ok := freq[v]; ok {
			freq[v] +=1
		}else {
			freq[v] = 1
		}
	}
	for k, v := range freq {
		fmt.Printf(" k = %d and v = %d \n",k,v)
	}
	return len(freq)
}

func CountUniqueValues1(arg []int) int {
	 var i = 0

	 if len(arg) == 0 {
	 	return 0
	 }
	 for j := 1; j < len(arg); j++ {
	 	if arg[i] != arg[j] {
	 		i++
			arg[i] = arg[j]
		}
	 }
	 fmt.Println(arg)
	 i++ // Adding 1 since i started from Zero
	return i
}


func maxSubSliceSum(arg []int, num int) int {
	var maxnum = 0
	var tempsum = 0
	for i:= 0 ; i < num ; i ++ {
		maxnum += arg[i]
	}

	for i := num ; i < len(arg); i++ {
		tempsum = tempsum - arg[i - num] + arg[i]
		fmt.Printf("[%d] - tempsum = %d \n", i, tempsum)
		if tempsum > maxnum {
			maxnum = tempsum
		}
	}
	return maxnum
}

func search(arr []int, num int) bool {

	var initNum = 0
	var lastNum = len(arr)

	for initNum < lastNum {

		tmp := (lastNum + initNum) /2

		if  num > tmp {
			initNum = tmp

		} else if num < tmp {
			 lastNum = tmp

		} else if tmp == num {
			return true
		}
	}
	return false
}

// Iterative version of Fibonacci
func fibonnaci(n int) int {
	var result int
	fmt.Printf("Starting Fibonnaci calculatin .....\n")
	for i, first, second := 0, 0, 1; i <= n; i, first, second = i+1, first+second, first {
		if i == n {
			result = first
		}
		fmt.Printf("Fibonnaci of i = %v, first = %v, second %v \n", i, first, second)
	}
	return result
}
