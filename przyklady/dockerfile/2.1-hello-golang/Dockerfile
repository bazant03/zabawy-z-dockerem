FROM golang:alpine

# Pobieramy program curl i czyścimy cache menadżera apk
RUN apk add --update curl && \
    rm -rf /var/cache/apk/*

COPY ./src ./src

# Kompilujemy program, domyslnie jest on linkowany statycznie
RUN go build -o ./bin/main ./src/main.go 

EXPOSE 8080

# Polecenie, które zostanie wykonane wewnątrz kontera, aby stwierdzić czy aplikacja działa poprawnie
HEALTHCHECK --interval=10s --start-period=2s --retries=3 \
            CMD curl -f http://localhost:8080 || exit 1

ENTRYPOINT [ "./bin/main" ]
