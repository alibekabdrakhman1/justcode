package main

import (
	"context"
	"fmt"
	"github.com/alibekabdrakhman/justcode/lecture9/models"
	"github.com/alibekabdrakhman/justcode/lecture9/repository"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	repo, err := repository.NewStorage(ctx)
	if err != nil {
		log.Fatal(err)
	}
	student := models.Student{
		Name:  "Imash",
		Class: "11a",
	}
	subject := models.Subject{
		Teacher:  "Apai",
		Students: []string{"Imash", "Erma"},
	}
	teacher := models.Teacher{
		Name:    "Apai",
		Classes: []string{"11a", "10e"},
	}
	// ya znauy chto nado error vozvrawat poslednim, tupanul i len stalo izmenyat vse(( sorry
	err, id := repo.Student.Create(ctx, student)
	if err != nil {
		log.Fatal(err)
	}
	err, student = repo.Student.GetById(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(student)

	err, id = repo.Subject.Create(ctx, subject)
	if err != nil {
		log.Fatal(err)
	}
	err, subject = repo.Subject.GetById(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(subject)

	err, id = repo.Teacher.Create(ctx, teacher)
	if err != nil {
		log.Fatal(err)
	}
	err, teacher = repo.Teacher.GetById(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(teacher)
}
