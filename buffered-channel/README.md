# Buffered Channel - Example

In this example, a producer will generate tasks and send them to a buffered channel, and multiple worker goroutines will receive the tasks and process them concurrently.

The buffered channel allows the producer to send tasks without waiting for immediate consumption by workers, up to the buffer's capacity.

## Problem: Simulating Task Processing with a Buffered Channel

- A producer will generate tasks and send them to a buffered channel.
- Multiple workers will receive tasks from the channel and process them.
- Each worker will simulate some work by sleeping for a random duration.

## Explanation of the Code:

- **Task Struct :** Each task is represented by a simple Task struct that contains an ID.

- **Producer :** The producer function generates a specified number of tasks (10 in this case) and sends them to the buffered channel (tasks). The channel has a buffer size of 3, which means the producer can send up to 3 tasks before it blocks if no worker consumes the tasks.

- **Workers :** Each worker is represented by the worker function, which continuously listens on the tasks channel for new tasks.
Once a task is received, the worker processes it (simulated by time.Sleep) and then moves on to the next task.
There are 3 workers (workerCount = 3), meaning tasks are processed concurrently by up to 3 workers at a time.

- **Buffered Channel :** The `tasks := make(chan Task, bufferSize)` creates a buffered channel with a buffer size of 3. The producer can send 3 tasks without waiting for workers to consume them, but once the buffer is full, the producer blocks until a worker takes a task from the channel.

- **WaitGroup :** A sync.WaitGroup is used to wait for all worker goroutines to complete their work. Each worker calls wg.Done() after processing all tasks.

- **Closing the Channel :** The producer closes the channel after sending all tasks (close(tasks)), which signals the workers to stop receiving tasks.

## Key Takeaways:
Buffered Channel allows the producer to send tasks to the workers without immediately blocking, up to the buffer size (3 in this case).
Workers process tasks concurrently, and once the buffer is full, the producer waits until the buffer has space.
The channel is closed by the producer to signal the end of task generation, and the workers stop processing once all tasks are handled.
