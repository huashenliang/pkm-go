# Golang : Gin and gorm with mysql

This repository is the backend for pokemon client: https://github.com/huashenliang/pkm-client

###Before running this app:

###### Make sure go is installed
- Install Go -> https://go.dev/doc/install
- Run `go version` to make sure go 1.19 or greater is being installed

###### Make sure you have a mysql database up and running 
For using docker:
1. Run `docker pull mysql`
2. Run `docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=mysqlpw -d mysql:tag`
3. Find more info at: https://hub.docker.com/_/mysql


###To run this app
1. Create `.env` file for the following format:
  ```
  POKEMON_API_KEY='your-api-key'
  DB_USERNAME='your-uersname'
  DB_PASSWORD='your-password
  DB_NAME='your-database-name'
  DB_HOST='your-database-host'
  DB_PORT='your-database-port'
  ```
2. Make sure your database is running
3. Run `go build main.go`
4. Run the exe file `.\main.exe` or run `go run main.go`

After running the commands, you can find server started on `8080` .

Some example enpoints:
- Get - localhost:8080/getDeckList
- Get - localhost:8080/getPokemonTypes
- Post - localhost:8080/generateDeck
      - With body:       ```
      {
          "PokemonType": "Darkness",
          "DeckName": "My Deck"
      }
      ```
