// Phase 01 の手動確認用エントリ（要件 6: HTTP なし）。
package main

import (
	"fmt"
	"log"

	"task-management-app/backend/internal/repository/memory"
	"task-management-app/backend/internal/usecase"
)

func main() {
	repo := memory.NewTaskRepository()
	svc := &usecase.TaskService{
		Repo:  repo,
		Clock: usecase.RealClock{},
	}

	a, err := svc.CreateTask("  First task  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created: id=%s title=%q\n", a.ID(), a.Title())

	b, err := svc.CreateTask("Second")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created: id=%s title=%q\n", b.ID(), b.Title())

	list, err := svc.ListTasks()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("list (created asc, id tie-break):")
	for _, t := range list {
		fmt.Printf("  - %s %q\n", t.ID(), t.Title())
	}

	got, err := svc.GetTask(a.ID())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("get one: %q\n", got.Title())

	updated, err := svc.UpdateTaskTitle(a.ID(), "Renamed")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("renamed: %q\n", updated.Title())

	if err := svc.DeleteTask(b.ID()); err != nil {
		log.Fatal(err)
	}

	list2, err := svc.ListTasks()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("after delete: count=%d\n", len(list2))

	_, err = svc.CreateTask("   ")
	if err != nil {
		fmt.Printf("validation (expected): %v\n", err)
	}

	_, err = svc.GetTask(b.ID())
	if err != nil {
		fmt.Printf("not found (expected): %v\n", err)
	}
}
