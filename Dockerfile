# Run the build
FROM golang:alpine
ENV WORKDIR /go/src/github.com/mojlighetsministeriet/gui
COPY . $WORKDIR
WORKDIR $WORKDIR
RUN apk --update add git nodejs-npm
RUN go get -t -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
RUN cd client && npm install && npm run build

# Create the final docker image
FROM scratch
COPY --from=0 /go/src/github.com/mojlighetsministeriet/gui/client/build/es6-unbundled /client
COPY --from=0 /go/src/github.com/mojlighetsministeriet/gui/gui /
ENTRYPOINT ["/gui"]
