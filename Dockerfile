FROM golang:1.24.0-bookworm as builder

WORKDIR /usr/src/app

RUN go install github.com/a-h/templ/cmd/templ@v0.3.856

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN templ generate
RUN go build -v -o /run-app .

FROM node:23.0.0-bookworm

WORKDIR /usr/src/app
COPY . .

RUN yarn install
RUN npx tailwindcss -i ./tailwind.css -o ./public/index.css

FROM debian:bookworm

COPY --from=builder /run-app /usr/local/bin/
CMD ["run-app"]
