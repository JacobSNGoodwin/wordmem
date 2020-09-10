FROM node:lts-alpine3.12 as builder

WORKDIR /app

COPY package*.json ./
COPY tsconfig*.json ./

RUN npm install

COPY src src

RUN if [ "${NODE_ENV}" != "development" ]; then npm run build; fi
# run this tag in docker-compose with npm run dev command


FROM node:lts-alpine3.12
WORKDIR /app

# do I actually need this?
ENV NODE_ENV=production

# Only install prod dependencies since we don't need types and stuff here
COPY package*.json ./
RUN npm install --only=prod

COPY --from=builder /app/dist ./dist

# Can add CMD here, or we can pass it when running docker container run
EXPOSE 8080

CMD ["npm", "run", "start"]