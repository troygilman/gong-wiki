FROM node:23.0.0-bookworm as css-builder

WORKDIR /app
COPY . .

RUN yarn install
RUN npx tailwindcss -i ./tailwind.css -o /css/index.css

FROM golang:1.24.0-bookworm as builder

WORKDIR /app

COPY . .
COPY --from=css-builder /css/index.css ./public/index.css

RUN go mod download && go mod verify

RUN go install github.com/a-h/templ/cmd/templ@v0.3.856
RUN templ generate

RUN go build -v -o /run-app .

FROM debian:bookworm

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
