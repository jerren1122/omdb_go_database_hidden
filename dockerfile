#inherit homebrew
FROM golang
RUN go get -u github.com/gorilla/mux
CMD git clone https://github.com/jerren1122/omdb_go_database_hidden.git; cd omdb_go_database_hidden; go build; ./omdb_go_database_hidden;