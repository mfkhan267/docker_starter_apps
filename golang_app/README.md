# Demo Manifests And Application Used In DevOps Toolkit Videos

* [Containers Are Not VMs! Which Base Container (Docker) Images Should We Use?](https://youtu.be/82ZCJw9poxM)

* [Using Docker Multi-Stage Builds](https://youtu.be/zpkqNPwEzac)

# Setup #

git clone https://github.com/mfkhan267/base-container-images-demo

cd base-container-images-demo

# Demo #

cat Dockerfile-full

cat Dockerfile-debian

cat Dockerfile

cat Dockerfile-alpine

# Build #

docker image build -t demo_go .

![image](https://github.com/mfkhan267/docker_starter_apps/assets/77663612/bdcf5c1f-0a90-4601-87ea-f4fb2a96b23a)

# Run #

docker container run --name go_demo -d -p 8000:8080 demo

![image](https://github.com/mfkhan267/docker_starter_apps/assets/77663612/8bd3a7aa-168e-4a21-bcbf-e4b7f5ecb330)

![image](https://github.com/mfkhan267/docker_starter_apps/assets/77663612/26b38a6e-96f0-4904-9d53-0c63abe11ae5)


