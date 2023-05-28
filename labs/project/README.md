## Project

1. Task: Create a Todo List Manager program that allows users to `create`, `view`, and `complete` tasks. Each task will have a `name`, a `description`, and a `status` (completed or not completed).

2. Setting up the project: Start by creating a new file called `todolist.go` using a code editor of your choice and define the main package.

3. Define the `Task struct`: Below the package declaration, define a struct called Task with `name`, `description`, and `status` fields

4. Implement the `NewTask` method, the function should create a new task and adds it to a slice of tasks.

5. Implement the `CompleteTask` method.

6. The program should: 
  a. be able to create several tasks and store them in a slice. 
  b. Call the `NewTask` method with the values to add it. You don't need to ask the user for input from the command line.
  c. Use the Task index to mark the task as completed.
  d. Finally, iterate over all the tasks and display their details. 

7. Build and run the program
