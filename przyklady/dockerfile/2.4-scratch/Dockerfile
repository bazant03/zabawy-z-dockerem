# Etap budowania ----------------------------------
FROM golang:alpine AS build

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

COPY --chown=appuser:appgroup ./src ./src

# Dzięki tym flagom binarka będzie troche lżejsza
RUN go build -ldflags="-s -w" -o ./bin/main ./src/main.go 

# Etap finalny ----------------------------------
# Ten program nie potrzebuje żadnych zależności, można go uruchomić w "pustym" kontenerze
FROM scratch AS final

LABEL org.opencontainers.image.source="https://github.com/bazant03/zabawy-z-dockerem" \
      org.opencontainers.description="Potezna aplikacja"

WORKDIR /app

# W systemie plików nie ma żadnej wzmianki o nazwach użytkowników, 
# ale można używać bezpośrednio uid, dla jądra to wystarczy 
USER 10001:10001

COPY --from=build --chown=10001:10001 /go/bin/main ./

EXPOSE 8080

#Jak zrobić healthcheck'a? - tutaj nie ma curl, nawet nie ma apk żeby go pobrać!

ENTRYPOINT [ "./main" ]