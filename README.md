# "Calculator"

This is sources of calculator, written on Golang.
App can do base calculations:

1. Summirizing
2. Substructing
3. Multiplying
4. Deviding
5. Get number to power
6. Take root from number

### LOCAL RUN:

There are two ways to run this application locally:

1. run docker container
2. run source file main.go.

To run project via docker do next steps:<br>

1. Install docker: https://docs.docker.com/engine/install/
2. In CMD in directory of project run next commands:<br>
   1. <i><b>sudo docker build . -t calculator</b></i><br>
   2. <i><b>sudo docker run -i calculator</b></i>

To run project using source files run next commands in CMD in directory of project :<br>

1. <i><b>go build -o calculator ./src/main/main.go</b></i><br>
2. <i><b>go run ./calculator</b></i><br>
