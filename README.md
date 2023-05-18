

1. use go run cmd/main/main.go to run the program, make sure you are connected to the db instance.

2. for db setup -

pull mysql docker image - 

docker pull mysql

start the container for mysql - 

docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=suryakant1234 -e MYSQL_DATABASE=godb -p 3306:3306 -d mysql

Note - above command has the info which is already used in myrepo


3. u can use the postman collection to check the payload and other info

learning-

1. gorilla mux includes more boiler plate as compared to gin and fiber which is made more developer friendly
2. you can also pass context to gorm similar to mongo db drivers using db.WithContext(ctx)

2b. you can use gorm tags to give metainfo like primary key, database types for sql db and db columnnames

3. when you are using mysql you can give some parametes like parseTime=true so that the you get the time from db in go format rather than in mysql system format
4. for debugging you can you runtime/trace package and use trace.start and stop to cathc traces in a file
5. you can see gorm documentation for other complex queries and aggregations function
6. gorm.io is a better version for github.com/jinzhu/gorm with improved apis
7. you can use gorm.model to add ID primary key and other audit fields

7b. gorm does error handling in a different way like it does not export an err out of its apis explicity but set error property in returned *DB instance which can be checked using .ERROR on it
