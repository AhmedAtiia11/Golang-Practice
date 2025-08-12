package main

import (
	"fmt"
	"sync"
	"studentapp/grade"
	"studentapp/stats"
)

func main() {
	s := grade.Student{Name: "Ahmed", Grades: []float64{50, 85, 88, 92}}
	ch := make(chan string) 
	var wg sync.WaitGroup
	wg.Add(len(s.Grades))
	for i, g := range s.Grades {
		go func(index int, score float64) {
			defer wg.Done()
			result := fmt.Sprintf("Grade %d: %.2f (Pass: %v)", index, score, grade.IsPassing(score))
			ch <- result
		}(i, g) // Closure fix
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for msg := range ch {
		fmt.Println(msg)
	}
	fmt.Printf("Average: %.2f\n", stats.Average(s.Grades))
	fmt.Printf("Max Grade: %.2f\n", stats.MaxGrade(s.Grades))}