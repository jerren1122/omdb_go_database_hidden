#overview
This is a POC of nesting 1 api interaction inside another

The surfaced API is written in Go, The hidden API interaction is written in Ruby

The surfaced API(Go) passes a movie name to the hidden API Interaction(Ruby) to OMDB(Open Movie Database) OMDB retrieves the results for us and the ruby script chooses to only return the year the movie was made. 

This allows us to build a network of API Connections that are not dependent on knowing how the hidden
API's function. 

#building the go project
This link is very helpful: https://medium.com/the-andela-way/build-a-restful-json-api-with-golang-85a83420c9da
Additionally if you have go installed 
utilizing 'go build' and './go-rest-api' (as this is the directory name) will locate the main package and execute accordingly. 

#debugging
VSCode has a debugging functionality for Go. We can run the go script in the debugger and then interact with the API via another tool like postman this will allow us to hit the breakpoints. 

To interact in Postman once the API is up and running we would contact: 
http://localhost:8080/get-movie (as this is running on your local)
for the body we would pass {
                           "Movie": "Twister" 
                           }
                           
#error handling 

Go Layer: At the go layer we have implemented error handling in order to ensure that a movie is passed in the correct JSON Format 

Ruby Layer: The Ruby Layer checks to ensure that a movie was found if it is not it returns an error message
This information is passed to the Go Layer which returns that as output to the end user. 
                           
                           
