FROM golang:1.19 as build

WORKDIR /app

ENV GIN_MODE=release

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

FROM python:3.11

WORKDIR /app

COPY --from=build ./main .
COPY --from=build requirements.txt .

# # Install Python3 and Pip3
# RUN apt-get update -y
# RUN apt-get install software-properties-common -y
# # RUN add-apt-repository add-apt-repository "deb http://archive.ubuntu.com/ubuntu $(lsb_release -sc) universe" -y
# RUN apt-get update -y
# RUN apt-get install python3 python3-pip python3-venv -y
# RUN python3 -m venv .venv
# RUN sudo .venv/bin/activate
RUN python3 -m pip install -r requirements.txt

EXPOSE 6001


CMD [ "/main" ]
