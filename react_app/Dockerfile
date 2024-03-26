FROM node:latest as build
WORKDIR /app
COPY package* yarn.lock .
RUN yarn install
COPY . .
RUN yarn run build

FROM nginx
RUN mkdir /app
COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
COPY --from=build /app/build /app