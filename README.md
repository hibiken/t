# t
a command-line tool to manage todos for a project.

## Install
```shell
go get -u github.com/hibiken/t

go install github.com/hibiken/t
```


## Get Started
```shell
# First initialize t in the current directory, just like `git init` :)
t init

# Add some todos
t add "First thing I need to do today"
t add "Second thing I need to do today"

# List todos
t ls

# To see done todos, pass 'all' flag
t ls --all
t ls -a

# Mark todo as done
t done [id]

# Mark todo as undone
t undone [id]

# Deletes todo
t delete [id]

# This deletes only done todos
t prune

# This deletes all todos
t clearall

# Run help to get this info
t help
t --help
```