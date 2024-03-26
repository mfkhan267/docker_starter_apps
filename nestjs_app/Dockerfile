# Base image
FROM node:latest as build

# Create app directory
WORKDIR /app

# A wildcard is used to ensure both package.json AND package-lock.json are copied
COPY package*.json .

# Install app dependencies
RUN npm install

# Bundle app source
COPY . .

# Creates a "dist" folder with the production build
RUN npm run build

FROM node:latest 

WORKDIR /app

COPY package.json .

RUN npm install --only=production

COPY --from=build /app/dist ./dist

# RUN npm run start:prod

# Start the server using the production build
CMD [ "node", "dist/main.js" ]