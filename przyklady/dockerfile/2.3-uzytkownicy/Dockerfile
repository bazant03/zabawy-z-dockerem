#Etap budowania ----------------------------------
FROM  golang:alpine AS build

# Tworzymy nowego użytkownika i grupę aby nie korzystać root'a
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Nowy użytkownik powinien być równierz właścicielem plików "aplikacji"
COPY --chown=appuser:appgroup ./src ./src

RUN go build -o ./bin/main ./src/main.go 

#Etap finalny ----------------------------------
FROM alpine AS final

LABEL org.opencontainers.image.source="https://github.com/bazant03/zabawy-z-dockerem" \
      org.opencontainers.description="Potezna aplikacja"

# To akurat musi wykonać root
RUN apk add --update curl && \
    rm -rf /var/cache/apk/*

WORKDIR /app

# Analogicznie jak w pierwszym etapie
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

COPY --from=build --chown=appuser:appgroup /go/bin/main ./ 

EXPOSE 8080

HEALTHCHECK --interval=10s --start-period=2s --retries=3 \
            CMD curl -f http://localhost:8080 || exit 1

ENTRYPOINT [ "./main" ]
