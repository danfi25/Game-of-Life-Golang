## Game of Life

This project involves creating the popular [Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life), which is a common thing to do in 
software development. The project is divided into two parts - server deployment and access through a web tab
and a standard 
terminal version of the game.


## Project Structure

The server-side of the project is divided into modules (following Uber's style guidelines), 
each of which performs its own task. The executable file is located in the `cmd` folder.

- `server`
    - `cmd` // storage of `main` packages
        - `life`
            - `main.go` // entry point to the program
    - `http`
        - `server` // HTTP server
            - `server.go` // server code
            - `handler` // registration of handler functions
                - `handler.go`
    - `internal`
        - `application` // configuration and code for invoking the application
            - `application.go`
        - `service` // service that initializes and stores the game state
            - `service.go`
    - `pkg` // storage of packages
        - `life` // game logic
            - `life_test.go`
            - `world.go`

## Running the Program

* To run the terminal version of the application, you need to go to the `terminal_version` folder and run the program with the command:
```Bash
go run main.go
```

* To run the server version of the program, you need to go to the `server_version` folder, initialize the Go module, install the Uber logger (already included in go.mod), and in the `cmd/life` folder, run `main.go`. When you open the website `localhost:8081/nextstate`, you will see the state of your game world, subsequent page updates will lead to the world state being updated.
```Bash
go mod init industrial_life
go get go.uber.org/zap
cd cmd/life
go run main.go
```