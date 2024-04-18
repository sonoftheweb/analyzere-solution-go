# Go implementation of analyzere problem

I decided to do the go version to show my skills in the language. The code below implements the problem statement as
described.

### Docker

- Build the Docker image: `docker build -t compute .`
- Run the Docker container with the command: `cat input.txt | docker run -i --rm compute 100 5000`

### Non Docker

- Compile the Go code: `go build compute.go`
- Run the program with the command: `cat input.txt | ./compute 100 5000`

Make sure in both scenarios, you have the input.txt file in the same directory as the compiled binary or the Dockerfile.