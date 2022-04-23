# syntax=docker/dockerfile:1

FROM golang:1.17

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p logs
RUN mkdir -p temp

RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init

RUN go build -o ./teamdev-backend

EXPOSE 8000

CMD [ "./teamdev-backend" ]
