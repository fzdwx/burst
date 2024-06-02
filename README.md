# burst

Expose ports to the server quickly

## Installation

```sh
go install github.com/fzdwx/burst@main
```

## Usage

```sh
# on server
burst serve

# export local port 8888 to server
burst client ap tcp:::8888
```
