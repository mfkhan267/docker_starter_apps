FROM golang:1.20.5-alpine AS build
RUN mkdir /src
ADD ./*.go /src
ADD ./go.mod /src
ADD ./go.sum /src
WORKDIR /src
RUN go get -d -v -t
RUN GOOS=linux go build -v -o golang-demo 
RUN chmod +x golang-demo

FROM scratch
COPY --from=build /src/golang-demo /usr/local/bin/golang-demo
EXPOSE 8080
CMD ["golang-demo"]
