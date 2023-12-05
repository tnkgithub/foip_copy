FROM node:16-alpine AS builder

WORKDIR /app

COPY package.json ./
COPY package-lock.json ./

RUN npm i

COPY . .

RUN npm run build

FROM nginx:1.23.1-alpine

RUN rm -rf /usr/share/nginx/html

COPY --from=builder /app/dist /usr/share/nginx/html
COPY ./proxy/nginx.conf /etc/nginx/nginx.conf
