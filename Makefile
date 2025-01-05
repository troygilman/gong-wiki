dev:
	templ generate
	go run .

css:
	npx tailwindcss -i ./tailwind.css -o ./public/index.css --watch
