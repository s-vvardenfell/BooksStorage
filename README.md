# BooksStorage

### How to run
```bash
docker-compose up
```
Database will be created and populated from [database/init.sql](https://github.com/s-vvardenfell/BooksStorage/blob/main/database/init.sql)<br>

### Test and debug
For testing/debugging, it is possible to use [grpcui](https://hub.docker.com/r/wongnai/grpcui) 
(set `reflection:true` to [config.yml](https://github.com/s-vvardenfell/BooksStorage/blob/main/resources/config.yml)) and [Adminer](https://hub.docker.com/_/adminer/)
<br>
Run test by usual way `go test ./...`