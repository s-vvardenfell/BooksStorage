# Description
Service provides the following procedures to call: 
* getting a list of books by author name 
* getting a list of authors by book title

MySQL is used as a database, gRPC is used as a protocol.

### How to run
* Use docker: `docker-compose docker-compose.yml up --build`
  <br>Database will be created and populated from [database/init.sql](https://github.com/s-vvardenfell/BooksStorage/blob/main/database/init.sql)<br>
* Use [Makefile](https://github.com/s-vvardenfell/BooksStorage/blob/main/Makefile) commands

### Test and debug
* Use [grpcui](https://hub.docker.com/r/wongnai/grpcui) 
(set `reflection:true` to [config.yml](https://github.com/s-vvardenfell/BooksStorage/blob/main/resources/config.yml)) and [Adminer](https://hub.docker.com/_/adminer/)
* Run tests by usual way: `export DSN=user:pass@/kvadoru && go test ./...` or from Makefile: `make tests`
