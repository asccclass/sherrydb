
run:
	DBMS=mysql DBSERVER=db4free.net DBPORT=3306 DBNAME=flowmaster DBLOGIN=andyliu DBPASSWORD=e2a87d75 go run example.go
test:
	go test
