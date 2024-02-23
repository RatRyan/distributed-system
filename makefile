db:
	sudo docker-compose up database

up:
	sudo docker-compose up --build -d --remove-orphans

down:
	sudo docker-compose down
