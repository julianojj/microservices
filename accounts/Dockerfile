FROM golang
WORKDIR /usr/src/app
COPY ./go.mod /usr/src/app/
RUN go mod tidy
COPY . .
