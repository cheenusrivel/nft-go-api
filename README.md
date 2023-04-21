# nft-go-api

Golang API with gin-gorm framework

The API will collect National Registration Identity Card (NRIC) and wallet address from WebApp

Store into RDBS (PostgreSQL) the unique NRIC and wallet address

POST API Response with a Receipt produce by hash the API body (Hash is used due to one way encrytion)

docker-compose.yaml script for PostgreSQL DB

<h2>Setup</h2>
Run the below commands in VSCode terminal to install Golang necessary artifacts

* $ go get -u gorm.io/gorm - Install gin-gorm framework
* $ go get gorm.io/driver/postgres -  Install postgres DB driver
* $ go get github.com/gin-contrib/cors - Install cors library to allow local origin
* $ docker-compose up --build - To install postgres DB in docker adn running in Port 5432. Make sure, Docker Desktop is running on local before running this command.

<h2>Run nft-go-api</h2>
Run the below commands in VSCode terminal to start the application

* $go run main.go - Application launches with URL http://localhost:8080/

<h2> Output </h2>

<h3>Get Command</h3>

![image](https://user-images.githubusercontent.com/88041827/233524524-0f3bce03-88e2-4209-ad2d-4ee2b5b0fcd3.png)

<h3>DB Table</h3>

![image](https://user-images.githubusercontent.com/88041827/233525141-6e479ec3-95e6-4e46-b0a4-4c861a0a5c4c.png)





