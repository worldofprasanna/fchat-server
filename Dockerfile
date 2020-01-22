FROM golang:latest

LABEL maintainer="Imran Basha <ibasha66@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 4040

CMD ["./main"]
