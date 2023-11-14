FROM alpine:3.17.3 AS base

RUN apk add --update-cache \
    curl unzip

# add the bootstrap file
COPY bootstrap.sh /tshock/bootstrap.sh

ARG TSHOCK_VERSION="v5.2.0"
ARG TERRARIA_VERSION="1.4.4.9"

# Download and unpack TShock
RUN set -eux; \
    arch="$(apk --print-arch)"; \
    case "$arch" in \
        'x86_64') \
            export TSHOCKZIP='TShock-5.2-for-Terraria-$TERRARIA_VERSION-linux-x64-Release.zip'; \
            ;; \
        'aarch64') \
            export TSHOCKZIP='TShock-5.2-for-Terraria-$TERRARIA_VERSION-linux-arm64-Release.zip'; \
            ;; \
        *) echo >&2 "error: unsupported architecture '$arch'."; exit 1 ;; \
    esac; \
    curl -L -o /$TSHOCKZIP https://github.com/Pryaxis/TShock/releases/download/$TSHOCK_VERSION/$TSHOCKZIP; \
    unzip $TSHOCKZIP -d /tshock ; \
    tar -xvf /tshock/*.tar -C /tshock ; \
    rm $TSHOCKZIP ; \
    chmod +x /tshock/TShock.Server ; \
    # add executable perm to bootstrap
    chmod +x /tshock/bootstrap.sh

# do not use -slim due to mysql/tshock requirements
FROM mcr.microsoft.com/dotnet/runtime:6.0

# documenting ports
EXPOSE 7777 7878

# env used in the bootstrap
ENV CONFIGPATH=/root/.local/share/Terraria/Worlds
ENV LOGPATH=/tshock/logs
ENV WORLD_FILENAME=""

# Allow for external data
VOLUME ["/root/.local/share/Terraria/Worlds", "/tshock/logs", "/plugins"]

# install nuget to grab tshock dependencies
RUN apt-get update -y && \
    apt-get install -y nuget mariadb-client jq && \
    rm -rf /var/lib/apt/lists/* /tmp/*

# copy game files
COPY --from=base /tshock/ /tshock/

# Set working directory to server
WORKDIR /tshock

# run the bootstrap, which will copy the TShockAPI.dll before starting the server
ENTRYPOINT [ "/bin/sh", "bootstrap.sh" ]