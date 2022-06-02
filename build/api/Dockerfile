FROM golang:alpine AS BASE
WORKDIR /go/src/github.com/llulioscesar/mercadolibre-x-men


ENV LANG en_US.UTF-8
ENV LC_ALL=C
ENV LANGUAGE en_US.UTF-8

FROM BASE as DEV
EXPOSE 80

COPY . .

RUN go mod tidy

# Dev Dependencies
CMD ["go", "run", "main.go"]

EXPOSE 80