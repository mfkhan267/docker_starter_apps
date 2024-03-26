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

![image](https://github.com/mfkhan267/base-container-images-demo/assets/77663612/b62ee35d-4f67-4bec-b73e-846639500ee8)

# Run #

docker container run --name go_demo -d -p 8000:8080 demo

![image](https://github.com/mfkhan267/base-container-images-demo/assets/77663612/9092f5d8-973d-402d-b910-311fa6ba7f20)

![image](https://github.com/mfkhan267/base-container-images-demo/assets/77663612/c5168617-9372-4351-9969-c3c20e7e6efb)


