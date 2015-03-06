GameOfLife
===========
GameOfLife is a simple implementation of Conway's Game of Life writen in Go with a web frontend.

Usage
--------
Download the binary distribution from the [releases](https://github.com/abijr/GameOfLife/releases) page, uncompress and run. Access the frontend from [localhost:8080](http://localhost:8080).

Development
-----------
The package is "go-gettable", get the sources with the command (assuming you have the Go installed and configured):

```
go get github.com/abijr/GameOfLife
```

Afterwards you need to pull the web interface runtime and development dependencies with:

```
npm install # this should be run from the GameOfLife/static/ folder
```

TODO
-----
* Functionality to clear the "board".
* 'Evolve' automatically (add an interval).
* Use Websockets to get following "generations".
* Write a qml front end.
