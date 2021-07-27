package main

import "github.com/kenriortega/go-poo-course/modulo1_y_2/course"

func main() {
	gocourse := course.NewCourse("Go desde 0", 23.33, true)
	gocourse.SetClasses(
		map[uint]string{
			1: "Introduccion",
		},
	)
	gocourse.PrintClasses()
}
