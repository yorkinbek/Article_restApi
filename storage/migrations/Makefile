init:
	migrate create -ext sql -dir storage/migrations -seq -digits 2 Init_db_artice
Up:
	migrate -path ./storage/migrations -database 'postgres://yorqin:yorqinbek@localhost:5432/article?sslmode=disable' up
Down:
	migrate -path ./storage/migrations -database 'postgres://yorqin:yorqinbek@localhost:5432/article?sslmode=disable' up

.PHONY: init Up Down