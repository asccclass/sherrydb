
run:
	DBMS=mysql DBSERVER=db4free.net DBPORT=3306 DBNAME=flowmaster DBLOGIN=andyliu DBPASSWORD=ooooo go run example.go
test:
	go test
