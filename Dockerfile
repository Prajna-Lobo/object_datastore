FROM golang:1.19

WORKDIR /app

COPY . ./

RUN go mod download
RUN go build -o /datastore app/main.go

EXPOSE 8080

CMD ["/datastore"]
