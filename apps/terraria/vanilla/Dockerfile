FROM alpine:3.17.3 AS base

RUN apk add --update-cache \
    unzip

# add the bootstrap file
COPY bootstrap.sh /terraria-server/bootstrap.sh

ARG TERRARIA_VERSION=1449
ARG ARTIFACT_URL=https://terraria.org/api/download/pc-dedicated-server/terraria-server-$TERRARIA_VERSION.zip
ARG ARTIFACT_FILENAME=terraria-server.zip

ADD $ARTIFACT_URL /$ARTIFACT_FILENAME

RUN unzip /$ARTIFACT_FILENAME -d /terraria && \
    mv /terraria/$TERRARIA_VERSION/Linux/* /terraria-server && \
    #Linux subfolder does not include any config text file, oddly.
    mv /terraria/$TERRARIA_VERSION/Windows/serverconfig.txt /terraria-server/serverconfig-default.txt && \
    chmod +x /terraria-server/TerrariaServer && \
    chmod +x /terraria-server/TerrariaServer.bin.x86_64

FROM mono:6.10.0.104-slim

# documenting ports
EXPOSE 7777

# env used in the bootstrap
ENV LOGPATH=/terraria-server/logs
ENV WORLDPATH=/root/.local/share/Terraria/Worlds
ENV WORLD_FILENAME=""
ENV CONFIGPATH=/config
ENV CONFIG_FILENAME="serverconfig.txt"

VOLUME [ "${WORLDPATH}", "${CONFIGPATH}", "${LOGPATH}" ]

COPY --from=base /terraria-server/ /terraria-server/

# Set working directory to server
WORKDIR /terraria-server

ENTRYPOINT [ "/bin/sh", "bootstrap.sh" ]
