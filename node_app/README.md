# DockerStarter

Download / Clone the Repository to a project folder

cd docker_starter

cd node_app

npm init # Accept Defaults

npm install express

node index.js # Test your app at localhost:3000 # You now have your First Hello World App running locally

Git Ignore File for NODE https://github.com/github/gitignore/blob/main/Node.gitignore

Docker Ignore File for NODE https://github.com/BretFisher/node-docker-good-defaults/blob/main/.dockerignore

docker build -t node_app . # Build your Docker Image Locally

docker run --rm -p 8080:3000 node_app # Run your first Docker Container with the above Image

The Containerize App is accessible over localhost:8080





