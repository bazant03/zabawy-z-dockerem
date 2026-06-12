#Etap budowania - w nim kompilujemy program
FROM  golang:alpine AS build

COPY ./src ./src

RUN go build -o ./bin/main ./src/main.go 

#Etap finalny - tutaj budujemy finalny obraz w którym umieszczamy tylko skompilowaną binarkę
FROM alpine AS final

RUN apk add --update curl && \
    rm -rf /var/cache/apk/*

WORKDIR /app

# Kopiujemy plik z poprzedniego etapu
COPY --from=build /go/bin/main ./ 

EXPOSE 8080

HEALTHCHECK --interval=10s --start-period=2s --retries=3 \
            CMD curl -f http://localhost:8080 || exit 1

ENTRYPOINT [ "./main" ]
