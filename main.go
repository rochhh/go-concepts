package main

import "fmt"

func main() {
   
   rect := shape{
      length: 10,
      width: 10,
   }
   fmt.Println(rect.area())
}

type shape struct{
   length int
   width int 
}


// struct specific method , can only be used by obj created from struct and no one else 
func (x shape) area () int{
   return x.length * x.width
}


