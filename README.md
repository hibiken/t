# t
a command-line tool to manage todos for a project.

## Install
```shell
go get -u github.com/hibiken/t

go install github.com/hibiken/t
```


## Get Started
```shell
# Add some todos
t add "First thing I need to do today"
t add "Second thing I need to do today"

# List todos
t list
t ls # alias of 'list'

# To see done todos, pass 'all' flag
t ls --all
t ls -a

# Mark todo as done
t done

# Mark todo as undone
t undone [id]

# Deletes todo
t delete [id]

# This deletes only done todos
t prune

# This deletes all todos
t clear

# Run help to get this info
t help
t --help
```