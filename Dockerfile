FROM golang:1.16.4-alpine

WORKDIR /app
COPY . /app

RUN go mod init sitemap-generator
RUN go mod tidy

RUN go build -o sitemap-generator

ENV XML_SITEMAP_FILE=/tmp/test.xml

CMD ["/app/sitemap-generator"]
