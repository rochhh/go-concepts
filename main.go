package main

import "fmt"

func main() {
   h1 := person{}
   h1.name = "roch"
   h1.sound = "meow"
   h1.insaan.someVar = "i have no idea"
   fmt.Println(h1.insaan.someVar)
}

type person struct {
   name string
   age int 
   height float32
   animal // embedded struct ; call by p.sound 

   insaan struct {
      someVar string // nested struct ; call by p.insaan.someVar
   }

}

type animal struct{
   name string
   sound string 
}

