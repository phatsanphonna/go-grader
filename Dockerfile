FROM golang:1.19 as build

WORKDIR /app

ENV GIN_MODE=release

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./main

FROM python:3.10.12 as production

WORKDIR /app

COPY --from=build /app/main .

RUN apt update -y
RUN apt install software-properties-common -y
RUN apt install pipx -y

RUN pipx install numpy

EXPOSE 6001

CMD [ "./main" ]
