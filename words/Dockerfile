FROM node:lts-alpine3.12

WORKDIR /app

COPY package.json .
COPY package-lock.json .

RUN npm install
COPY . .

CMD ["npm", "run", "dev"]