# Run the build
FROM mojlighetsministeriet/go-polymer-faster-build
ENV WORKDIR /go/src/github.com/mojlighetsministeriet/gui
COPY . $WORKDIR
WORKDIR $WORKDIR
RUN go get -t -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build
RUN cd client && yarn install
RUN cd client && yarn build

# Create the final docker image
FROM scratch
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=0 /go/src/github.com/mojlighetsministeriet/gui/client/build/es6-unbundled /client
COPY --from=0 /go/src/github.com/mojlighetsministeriet/gui/gui /
ENTRYPOINT ["/gui"]
