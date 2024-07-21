run:
	@templ generate
	@go run cmd/app_v1/main.go

templ:
	@templ generate --proxy="http://localhost:3000"
templ1:
	@templ generate --watch --proxy="http://localhost:3000" --cmd="go run ./cmd/app_v1/"
templ2:
	@templ generate --watch --proxy="http://localhost:3000" --cmd="tmp/main.exe"
templ-watch:
	@templ generate --watch --proxy="http://localhost:3000"
css-watch:
	@npx tailwindcss -i view/css/app.css -o public/css/style.css --watch
css:
	@npx tailwindcss -i view/css/app.css -o public/css/style.css
	