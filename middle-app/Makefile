DBNAME:=middle_task
DATASOURCE:=root@tcp(127.0.0.1:3306)/$(DBNAME)

migrate/new:
	@goose -dir migrations mysql "$(DATASOURCE)" create "$(NAME)" sql

migrate/status:
	@goose -dir migrations mysql "$(DATASOURCE)" status

migrate/up:
	@goose -dir migrations mysql "$(DATASOURCE)" up-by-one

__migrate/down:
	@goose -dir migrations mysql "$(DATASOURCE)" down
