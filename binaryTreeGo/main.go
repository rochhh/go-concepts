package main

import (
	"fmt"
)

func main() {
   var it *intTree
   it = it.Insert(5)
   it = it.Insert(3)
   it = it.Insert(2)
   it = it.Insert(10)
   fmt.Println(it.Contains(99))
}

type intTree struct {
   val int 
   left ,right *intTree
}


func (it *intTree) Insert(val int) *intTree{
   if it == nil {
      return &intTree{val: val}
   } else if (val < it.val){
      it.left = it.left.Insert(val)
   } else if (val > it.val){
      it.right = it.right.Insert(val)
   }
   return it
}

func (it *intTree) Contains( val int ) bool {
   switch {
   case it == nil:
      return false
   case val < it.val:
      return it.left.Contains(val)
   case val > it.val:
      return it.right.Contains(val)
   default:
      return true
   }
}