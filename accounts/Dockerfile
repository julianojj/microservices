FROM golang
WORKDIR /usr/internal/app
COPY ./go.mod /usr/internal/app/
RUN go mod tidy
COPY . .
