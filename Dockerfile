FROM library/golang:1.13-alpine

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go install -a std

ENV APP_DIR $GOPATH/src/github.com/blackironj/action-test
RUN mkdir -p $APP_DIR
WORKDIR $APP_DIR

ADD go.mod .
ADD go.sum .

RUN go mod download

ADD . $APP_DIR
# Compile the binary and statically link
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

EXPOSE 8080
# Set the entrypoint
ENTRYPOINT ./main
