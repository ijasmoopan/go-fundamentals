package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task represents a simple task with an ID.
type Task struct {
	ID int
}

// worker simulates a worker that processes tasks from the buffered channel.
func worker(tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("Task %d processing\n", task.ID)
		// fmt.Println("Doing something........")
		// Simulate work by sleeping for a random duration
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Printf("Task %d completed\n", task.ID)
	}
}

// producer generates tasks and sends them to the buffered channel.
func producer(totalTasks int, tasks chan<- Task) {
	for i := 1; i <= totalTasks; i++ {
		task := Task{ID: i}
		fmt.Printf("Task %d produced\n", i)
		tasks <- task // Send task to the buffered channel
	}
	close(tasks) // Close the channel after all tasks are produced
}

func main() {
	rand.Seed(time.Now().UnixNano())
	totalTasks := 10 // Total number of tasks to be produced
	bufferCount := 3 // Size of the buffered channel
	workerCount := 3 // Number of workers

	var wg sync.WaitGroup

	tasks := make(chan Task, bufferCount)

	// Start workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(tasks, &wg)
	}

	// Start the producer
	go producer(totalTasks, tasks)

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All tasks have been processed.")
}
