package main

import "fmt"

func main() {
   
   nums := [6]int{1,2,3,4,5,6}
   numsSlices := nums[1:3]
   
   
   // numsArray  := make( []int ,2 ,5  )
   
   fmt.Println(numsSlices )
   // fmt.Println( do(nums[:]) )

}

func do ( nums []int ) []int {
   return nums
}

