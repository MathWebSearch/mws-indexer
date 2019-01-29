FROM golang:1-alpine as builder

# Build dependencies
RUN apk add --no-cache make git

# Build this updater
ADD . /go/src/github.com/MathWebSearch/mws-updater
WORKDIR /go/src/github.com/MathWebSearch/mws-updater
RUN make build-local

# Add it into MathWebSearch
FROM mathwebsearch/mathwebsearch
COPY --from=builder /go/src/github.com/MathWebSearch/mws-updater/out/mwsupdate /mws/bin/mwsupdate

# Create volumes under harvest and index
VOLUME /harvest/
VOLUME /index/

# And add a label for the docker container
ENV MWS_DOCKER_LABEL ""

CMD [ "/bin/bash", "-c", "/mws/bin/mwsupdate", "--docker-label", "$MWS_DOCKER_LABEL" ]