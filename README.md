# gophercise-quiz
Gophercise exercise no.1 (Quiz Game)

## Brief
Part 1 of the Gophercises course, a quiz game written in Go.

## How to run
1. Clone repository
2. Navigate to directory
3. In directory, execute `go run main.go`
4. There are two flags:
   - `--path` allows you to use other files as the problem sheet by typing its path. It must be a .csv file with two columns. The first column is the text while the second column is the expected answer. By default, it will point to the packaged `problems.csv`.
   - `--time` allows you to define a custom time limit for the entire quiz (in seconds). By default, it's set to 30 seconds.
