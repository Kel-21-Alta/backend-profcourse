# Make container MYSQL
docker run --name profcourse-mysql -p 3300:3306 -e MYSQL_DATABASE=profcourse -e MYSQL_USER=profcourse -e MYSQL_PASSWORD=profcourse -e MYSQL_ROOT_PASSWORD=profcourse -d mysql:8.0.27
