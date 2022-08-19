/*Cola de prioridad haciendo uso del paquete heap de Go en una lista de frutas las cuales tienen una fecha de caducidad:
Su fecha de caducidad será en meses, se irá desencolando productos que esten proximos a vencerse

*/
package main

import (
	"container/heap"
	"fmt"
)

type Elemento struct {
	prioridad      int
	nombreElemento string
}

type ElementoPQ []Elemento

func (self ElementoPQ) Len() int {
	return len(self)
}

func (self ElementoPQ) Less(i, j int) bool {
	return self[i].prioridad < self[j].prioridad
}

func (self ElementoPQ) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self *ElementoPQ) Push(x interface{}) {
	*self = append(*self, x.(Elemento))
}

func (self *ElementoPQ) Pop() (popped interface{}) {
	popped = (*self)[len(*self)-1]
	*self = (*self)[:len(*self)-1]
	return
}

func main() {
	cola := &ElementoPQ{
		{5, "Manzanas gala"},
		{3, "Fresas"},
		{7, "Ciruelas pasa"},
		{1, "Bananos"}}

	heap.Init(cola)

	heap.Push(cola, Elemento{4, "Sandias"})

	for cola.Len() != 0 {
		fmt.Println("Producto desencolado: ", heap.Pop(cola))
	}
}
