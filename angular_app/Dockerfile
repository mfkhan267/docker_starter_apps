FROM node:latest as build
WORKDIR /app
COPY package*.json .
RUN npm install
COPY . .
RUN npm run build

FROM nginx
RUN mkdir /app
COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
COPY --from=build /app/dist/angular_app/ /app
