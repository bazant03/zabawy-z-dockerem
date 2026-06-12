# Budujemy nasz obraz na podstawie innego, już gotowego obrazu
FROM node:26-alpine

# Ustawiamy katalog /app jako katalog roboczy, jeśli nie istnieje, Docker go utworzy
WORKDIR /app

# Kopiujemy pliki projektu do środka obrazu 
COPY ./ ./

# Instalujemy zależności
RUN npm install 

# Informujemy, że aplikacja wykorzystuje dany port - to tylko dokumentacja
EXPOSE 8080

# Polecenie które zostanie wykonane po uruchomieniu kontenera
CMD [ "node", "./src/index.js"]