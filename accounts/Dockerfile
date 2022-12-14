FROM golang
WORKDIR /usr/src/app
COPY ./go.mod .
RUN go mod tidy
COPY . .