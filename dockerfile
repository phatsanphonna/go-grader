FROM golang:1.19 as build

WORKDIR /app

ENV GIN_MODE=release

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# Install Python3 and Pip3
RUN apt-get update -y
RUN apt-get install software-properties-common -y
# RUN add-apt-repository add-apt-repository "deb http://archive.ubuntu.com/ubuntu $(lsb_release -sc) universe" -y
RUN apt-get update -y
RUN apt-get install python3 pipx -y
RUN pipx install numpy

EXPOSE 6001


CMD [ "/main" ]
