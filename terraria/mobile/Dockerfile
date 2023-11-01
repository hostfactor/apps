FROM alpine:3.17.3 AS base

RUN apk add --update-cache \
    unzip

ARG TERRARIA_VERSION=1449
ARG ARTIFACT_URL=https://terraria.org/api/download/mobile-dedicated-server/terraria-server-$TERRARIA_VERSION.zip
ARG ARTIFACT_FILENAME=terraria-server-$TERRARIA_VERSION.zip

ADD $ARTIFACT_URL /$ARTIFACT_FILENAME

RUN unzip $ARTIFACT_FILENAME && \
    mkdir ./terraria-server  && \
    mv ./$TERRARIA_VERSION/Linux/* ./terraria-server && \
    rm ./terraria-server/System* && \
    rm ./terraria-server/Mono* && \
    rm ./terraria-server/monoconfig && \
    rm ./terraria-server/mscorlib.dll && \
    rm -r ./$TERRARIA_VERSION && \
    rm -r ./$ARTIFACT_FILENAME && \
    chmod +x ./terraria-server/TerrariaServer && \
    chmod +x ./terraria-server/TerrariaServer.bin.x86_64

FROM mono:6.12.0.182

EXPOSE 7777

VOLUME ["/root/.local/share/Terraria/Worlds", "/config"]

COPY --from=base /terraria-server/ /terraria-server/

ENTRYPOINT ["mono", "--server", "--gc=sgen", "-O=all", "/terraria-server/TerrariaServer.exe"]
