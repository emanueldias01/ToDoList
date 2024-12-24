FROM golang:1.23.2

EXPOSE 8080

WORKDIR /app

COPY ./model /app/model
COPY ./database /app/database/
COPY ./dto /app/dto/
COPY ./repository /app/repository/
COPY ./service /app/service/
COPY ./controller /app/controller/
COPY ./route /app/route/

COPY ./go.mod /app/
COPY ./go.sum /app/
COPY ./main.go /app/

CMD ["go", "run", "main.go"]
