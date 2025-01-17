main_package_path = ./cmd/server
binary_name = server

clean:
	go clean
	rm -r -f ./bin

dev:
	wgo -verbose -file=.go -file=.gohtml -file=.sql -file=.css \
		-xdir ./internal/models \
		-xfile ./assets/static/css/main.css \
		make sqlc :: \
		make tailwind :: \
		go run ${main_package_path}

build:
	@make clean
	@make sqlc
	@make tailwind_minify
	go build -o ./bin/${binary_name} ${main_package_path}

build_linux:
	@make clean
	@make sqlc
	@make tailwind_minify
	GOARCH=amd64 GOOS=linux go build -o ./bin/${binary_name}-linux ${main_package_path}

run:
	go run ${main_package_path}

install:
	go build ${main_package_path}

sqlc:
	sqlc generate

tailwind:
	tailwindcss -i ./input.css -o ./assets/static/css/main.css

tailwind_minify:
	tailwindcss -i ./input.css -o ./assets/static/css/main.css --minify

install_deps:
	@go install github.com/bokwoon95/wgo@latest
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
# ==================================================================================== #
# DATABASE
# ==================================================================================== #
database/backup:
	sqlite3 battlog.db .dump > testdata/battlog_backup.sql

database/restore:
	sqlite3 battlog.db < testdata/battlog_backup.sql

database/initialize:
	sqlite3 battlog.db < ./sql/migrations/create_battery.sql
	sqlite3 battlog.db < ./sql/migrations/create_battery_history.sql

