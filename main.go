package main

import "fmt"

func main() {
   
   bk := corporate{
      name: "burger king",
      employees: 100000,
   }

   fmt.Println(bk.employees + bk.factoryClose())  

}


func (c corporate ) factoryRunner() string {
   return c.name
}

func (c corporate) factoryClose() int {
   return 2
}

type corporate struct{
   name string 
   employees int 
}

type factory interface{
   factoryRunner() string
   factoryClose() 
}
