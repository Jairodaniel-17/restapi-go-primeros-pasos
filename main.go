package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some Content",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a mi API REST con GO")
}

// metodo para obtener todas las tareas
func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

// metodo para crear una tarea
func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte una tarea valida")
	}
	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

// metodo para obtener una tarea
func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "ID invalido")
		return
	}
	for _, task := range tasks {
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}

}

// metodo para actualizar una tarea
func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "ID invalido")
		return
	}
	var updatedTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte una tarea valida")
		return
	}
	json.Unmarshal(reqBody, &updatedTask)
	for i, task := range tasks {
		if task.ID == taskID {
			updatedTask.ID = taskID
			tasks[i] = updatedTask
			fmt.Fprintf(w, "La tarea con ID %v ha sido actualizada", taskID)
			return
		}
	}
	fmt.Fprintf(w, "No se encontró la tarea con ID %v", taskID)
}

// metodo delete
func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "ID invalido")
		return
	}
	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "La tarea con ID %v ha sido eliminada", taskID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", router))

}

/*
package main

import (
	"fmt"
)

func main() {
	// Imprimir en consola

	fmt.Println("Hola, Go!")
	// Declarar variables tipo string
	var cadenaTexto string = "Esto es una cadena de texto"
	// Imprimir variable
	fmt.Println(cadenaTexto)
	// cambiar valor de variable
	cadenaTexto = "Esto es otra cadena de texto"
	// Imprimir variable
	fmt.Println(cadenaTexto)
	// Declarar variable tipo int
	var numero int = 100
	// Imprimir variable
	fmt.Println(numero)
	// mezclar variables string + int
	fmt.Println("El numero es: " + fmt.Sprint(numero))
	// arreglo de numeros enteros
	var arregloNumeros [5]int
	// asignar valor a arreglo
	arregloNumeros[0] = 1
	arregloNumeros[1] = 2
	arregloNumeros[2] = 3
	arregloNumeros[3] = 4
	arregloNumeros[4] = 5
	// imprimir arreglo
	fmt.Println(arregloNumeros)
	// map de string a int
	myMap := make(map[string]float64)
	myMap["llave1"] = 1
	myMap["llave2"] = 2
	myMap["llave3"] = 3.54
	myMap["llave4"] = 4
	myMap["llave5"] = 5
	// imprimir map
	fmt.Println(myMap)
	// imprimir valor de una llave
	fmt.Println(myMap["llave1"])
	// bucles
	for i := 0; i < 20; i++ {
		fmt.Println(i)
	}
	// foreach
	for i := range arregloNumeros {
		fmt.Println(i, arregloNumeros[i])
	}

	// foreach map
	for i := range myMap {
		fmt.Println(i, myMap[i])
	}

	// funciones
	fmt.Println("Suma: ", suma(1.5464654656, 2.34563563454))

	// struct
	type User struct {
		id       int
		name     string
		lastName string
	}

	user1 := User{id: 1, name: "Juan", lastName: "Perez"}
	fmt.Println(user1.name, user1.lastName)
	// ingresar datos por consola
	var input string
	fmt.Scanln(&input)
	user2 := User{id: 2, name: input, lastName: "Mendoza"}
	fmt.Println("Hola, ", user2)
}

// funcion suma
func suma(a float64, b float64) float64 {
	return a + b
}

*/
