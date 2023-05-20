migrate-up:	
	migrate -database "mysql://mkceliks:159987741@tcp(localhost:3306)/yaz_gelsin?query" -path ./migrations up

migrate-down:	
	migrate -database "mysql://mkceliks:159987741@tcp(localhost:3306)/yaz_gelsin?query" -path ./migrations down

.PHONY: migrate-up migrate-down