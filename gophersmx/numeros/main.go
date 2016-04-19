// Crear un programa que declare dos funciones anónimas. Una que cuente del 0 al 200 y otra que cuente de 200 a 0.
// Mostrar cada número con un identificador único para cada goroutine.
// Crear goroutines a partir de estas funciones.
// No permitir que main regrese hasta que ambas goroutines hayan terminado de correr.

package main

// Agregar imports.
import (
   "runtime"
   "sync"
   "fmt"
)

func init() {

	// Alojar un procesador para que el scheduler lo use.
	runtime.GOMAXPROCS(1)
}

func main() {

	// Declarar el wait group e iniciar el contador en 2.
    var wg sync.WaitGroup
    wg.Add(2)

	// Declarar una función anónima y crear una goroutine.
	go func(from string, wg *sync.WaitGroup) {
		// Declarar un loop que cuente del 200 al 0 e imprimir cada valor
        for i := 200; i >= 0; i-- {
           fmt.Println(from, ":", i)
        }
        // Decrementar la cuenta del wait group.
        wg.Done()
	}("200 al 0", &wg)
    
    // Declarar una función anónima y crear una goroutine.
	go func(from string, wg *sync.WaitGroup){
		// Declarar un loop que cuente del 0 al 200 e imprimir cada valor
        for i := 0; i <= 200; i++ {
           fmt.Println(from, ":", i)
        }
		// Decrementar la cuenta del wait group.
		wg.Done()
	}("0 al 200", &wg)

    wg.Wait()
	fmt.Println("Done")
}