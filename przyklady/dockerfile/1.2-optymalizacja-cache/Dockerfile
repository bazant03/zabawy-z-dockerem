FROM node:26-alpine

WORKDIR /app

# Na początek kopiujemy samą listę zależności
COPY ./package.json ./package-lock.json ./

RUN npm install 

# Dopiero teraz kopiujemy resztę, zmiana w kodzie nie powoduje unieważnienia cache warstwy z zależnościami
COPY ./src ./src

EXPOSE 8080
CMD [ "node", "./src/index.js"]