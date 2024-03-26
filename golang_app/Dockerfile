FROM golang:alpine3.19 AS build
RUN mkdir /src
ADD ./*.go /src
ADD ./go.mod /src
ADD ./go.sum /src
WORKDIR /src
RUN go get -d -v -t
RUN GOOS=linux go build -v -o golang-app
RUN chmod +x golang-app

FROM scratch
COPY --from=build /src/golang-app /usr/local/bin/golang-app
EXPOSE 8080
CMD ["golang-app"]
