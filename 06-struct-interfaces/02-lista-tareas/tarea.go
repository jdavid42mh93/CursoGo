package main

import "fmt"

type TaskList struct {
	task []*Task
}

func (tl *TaskList) appendTask(t *Task) {
	tl.task = append(tl.task, t)
}

func (tl *TaskList) removeTask(index int) {
	tl.task = append(tl.task[:index], tl.task[index+1:]...)
}

type Task struct {
	name      string
	desc      string
	completed bool
}

func (t *Task) toPrint() {
	fmt.Printf("Nombre: %s\nDescripcion: %s\nCompletado: %t\n", t.name, t.desc, t.completed)
}

func (t *Task) markCompleted() {
	t.completed = true
}
func main() {
	t1 := Task{
		name:      "Curso de Go",
		desc:      "Completar curso de Go este mes",
		completed: false,
	}

	t2 := Task{
		name:      "Curso de HTML",
		desc:      "Completar curso de HTML este mes",
		completed: false,
	}

	list := TaskList{}
	list.appendTask(&t1)
	list.appendTask(&t2)

	fmt.Println(list)
	//t1.toPrint()
	//t2.toPrint()

	t3 := Task{
		name:      "Curso de CSS",
		desc:      "Completar curso de CSS este mes",
		completed: false,
	}
	list.appendTask(&t3)
	fmt.Println(list)
	list.removeTask(1)
	for i, task := range list.task {
		fmt.Println(i, task.name)
	}

}
