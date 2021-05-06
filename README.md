# clido
A command line based task manager/TODO list with a persistent database. Written in Go, it uses [Bolt DB](https://github.com/boltdb/bolt) as the key-value database.
To install, make sure you have Go installed, clone the repo, go to the root and run `go install .`. Used as `clido <command> <args>`.
 Commands supported currently include,
- `list` lists all the tasks not yet completed, in the order of decreasing priorities,`--late` flag lists all the tasks you're running late on
- `add <string>` adds a task `string` to the list of tasks you want to complete, `--priority` flag assigns a priority to the task
- `do <number>` marks the task number `number` in the list as done, can accept multiple numbers
- `completed` lists all the tasks completed today 
- `remove <number>` removes the task number `number` in the list as done, can take multiple numbers
