# Global arguments
ARG APP_GO_PATH="/go/src/github.com/regiocom/buildpipelinebeat"
ARG GO_VERSION="1.15.2"

# Buildcontext
FROM scratch as buildcontext
COPY . .


# App builder image
FROM golang:${GO_VERSION}-alpine as build

ARG APP_GO_PATH

# Install mage and dependencies
WORKDIR /tools
RUN apk update; \
    apk upgrade; \
    apk add --no-cache ca-certificates bash git openssh build-base; \
    git clone https://github.com/magefile/mage; \
    cd mage; \
    go run bootstrap.go;

# Prepare building the app
WORKDIR ${APP_GO_PATH}
COPY --from=buildcontext . .

# build the app
RUN mage -v build


#################### the final appimage ####################
FROM alpine as image

# arguments
ARG APP_GO_PATH

# create folders and install ca-certificates
RUN mkdir -p /etc/buildpipelinebeat/; \
    mkdir -p /app/; \
    apk add --no-cache ca-certificates

# Copy config and binary
COPY --from=buildcontext docker/buildpipelinebeat.yml /etc/buildpipelinebeat/config.yml
COPY --from=build ${APP_GO_PATH}/buildpipelinebeat /app/buildpipelinebeat

# startup parameter
ENV PATH="/app:${PATH}"

ENTRYPOINT [ "buildpipelinebeat", "-e", "-c", "/etc/buildpipelinebeat/config.yml" ]
