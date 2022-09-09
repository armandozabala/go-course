package main

import "fmt"

// Lista de Task
type TaskList struct {
	tasks []*Task
}

// Agregar elementos a la lista
func (tl *TaskList) appendTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) deleteTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

// Struc
type Task struct {
	name      string
	desc      string
	completed bool
}

// metodo
func (t *Task) toPrint() {
	fmt.Printf("Nombre: %s\nDescription: %s\nCompletado: %t\n", t.name, t.desc, t.completed)
}

// metodo
func (t *Task) markCompleted() {
	t.completed = true
}

func main() {
	t1 := Task{
		name:      "Curso de GO",
		desc:      "Completar",
		completed: false,
	}

	t2 := Task{
		name:      "Curso de Java",
		desc:      "Completado",
		completed: true,
	}

	lista := TaskList{}
	//enviar ref y no la copia
	lista.appendTask(&t1)
	lista.appendTask(&t2)

	fmt.Println(lista)

	t3 := Task{
		name:      "Curso de CSS",
		desc:      "por completar",
		completed: true,
	}

	lista.appendTask(&t3)
	fmt.Println(lista)
	//t1.toPrint()
	//t2.toPrint()

	lista.deleteTask(1)

	for i, task := range lista.tasks {
		fmt.Println(i, task.name)
	}
}
