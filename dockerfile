FROM golang:1.19

WORKDIR /app

ENV GIN_MODE=release

COPY . .
RUN go mod download

RUN apt-get install python3 python3-pip -y
RUN pip3 install -r requirements.txt

EXPOSE 6001

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

CMD [ "/main" ]
