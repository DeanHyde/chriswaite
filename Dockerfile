FROM golang

# Fetch dependencies
RUN go get github.com/tools/godep

# Add project directory to Docker image.
ADD . /go/src/github.com/DeanHyde/chriswaite

ENV USER dean
ENV HTTP_ADDR :8888
ENV HTTP_DRAIN_INTERVAL 1s
ENV COOKIE_SECRET wGTXafrku2fTihQV

# Replace this with actual PostgreSQL DSN.
# ENV DSN postgres://dean@localhost:5432/chriswaite?sslmode=disable

WORKDIR /go/src/github.com/DeanHyde/chriswaite

RUN go get

RUN go build

EXPOSE 8888
CMD ./chriswaite