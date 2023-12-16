mysql-start:
	docker start mysqlfuncom

mysql-stop:
	docker stop mysqlfuncom

mysql:
	docker run -it --rm mysql mysql -h172.17.0.1 -uroot -p
