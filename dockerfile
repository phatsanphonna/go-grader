FROM golang:1.19

WORKDIR /app

ENV GIN_MODE=release

COPY . .
RUN go mod download

# Install Python3 and Pip3
RUN apt-get update -y
RUN apt-get install software-properties-common -y
# RUN add-apt-repository add-apt-repository "deb http://archive.ubuntu.com/ubuntu $(lsb_release -sc) universe" -y
RUN apt-get update -y
RUN apt-get install python3 python3-pip -y
RUN pip3 install -r requirements.txt

EXPOSE 6001

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

CMD [ "/main" ]
