# Todos
a command-line tool to manage tasks for a project.

## Install
```shell
go get -u github.com/hibiken/todos

go install github.com/hibiken/todos
```


## Get Started
```shell
# First initialize todos in current directory
todos init

# Add some todos
todos add "First thing I need to do today"
todos add "Second thing I need to do today"

# List todos
todos ls

# To see done todos, pass 'all' flag
todos ls --all
todos ls -a

# Mark todo as done
todos done [id]

# Mark todo as undone
todos undone [id]

# Deletes todo
todos delete [id]

# This deletes only done todos
todos prune

# This deletes all todos
todos clearall

# Run help to get this info
todos help
todos --help
```