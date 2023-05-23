

1. use go run cmd/main/main.go to run the program, make sure you are connected to the db instance.

2. for db setup -

pull mysql docker image - 

docker pull mysql

start the container for mysql - 

docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=suryakant1234 -e MYSQL_DATABASE=godb -p 3306:3306 -d mysql

Note - above command has the info which is already used in myrepo


3. u can use the postman collection to check the payload and other info

