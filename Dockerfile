FROM golang:latest

RUN useradd sandbox
USER sandbox

WORKDIR /var/www/server

COPY /secrets/ddns-web-passwd .
COPY /secrets/ddnskey .
COPY go.mod .
COPY main.go .
COPY ddns-update .

RUN go build .

EXPOSE 8143

CMD [ "./ddns-update-server" ]
