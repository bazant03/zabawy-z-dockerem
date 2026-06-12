#Etap budowania ----------------------------------
# Koplilujemy obie binarki na natywnej architekturze hosta
FROM --platform=$BUILDPLATFORM golang:alpine AS build

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

COPY --chown=appuser:appgroup ./src ./src

ARG TARGETOS TARGETARCH
# Program zostanie skompilowany na podane architektury
RUN GDOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o ./bin/main ./src/main.go 

#Etap finalny ----------------------------------
# Dla każdej architektury automatycznie zostanie stworzony obraz
FROM scratch AS final

LABEL org.opencontainers.image.source="https://github.com/bazant03/zabawy-z-dockerem" \
      org.opencontainers.description="Potezna aplikacja"

WORKDIR /app

USER 10001:10001

COPY --from=build --chown=10001:10001 /go/bin/main ./

EXPOSE 8080

ENTRYPOINT [ "./main" ]