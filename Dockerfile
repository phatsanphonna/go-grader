FROM golang:1.20 as build

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o ./main ./main.go

# Python base image
FROM debian:bookworm as production

WORKDIR /root/app

ENV TZ Asia/Bangkok

COPY --from=build /app/main ./main
# COPY ./wine_inw.tar.gz .

# RUN echo "deb http://mirror.applebred.net/debian bullseye main contrib non-free" | tee -a /etc/apt/sources.list
RUN echo 'deb http://mirror.applebred.net/debian bookworm main contrib non-free' > /etc/apt/sources.list
RUN echo 'deb http://mirror.applebred.net/debian-security bookworm-security main contrib non-free' >> /etc/apt/sources.list
RUN echo 'deb http://mirror.applebred.net/debian bookworm-updates main contrib non-free' >> /etc/apt/sources.list

# RUN apt update -y
# RUN apt install -y dpkg
# RUN dpkg --add-architecture i386
RUN apt-get update -y
RUN apt-get install -y python3.11-minimal python3-pip pipx && rm -rf /var/lib/apt/lists/*

# Python3 modules
RUN pipx install numpy

CMD ./main