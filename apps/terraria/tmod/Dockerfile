FROM ubuntu as install-steam

# install steam
SHELL ["/bin/bash", "-o", "pipefail", "-c"]
RUN echo steam steam/question select "I AGREE" | debconf-set-selections \
 && echo steam steam/license note '' | debconf-set-selections

# Update the repository and install SteamCMD
ARG DEBIAN_FRONTEND=noninteractive
RUN dpkg --add-architecture i386 \
 && apt-get update -y \
 && apt-get install -y --no-install-recommends ca-certificates locales steamcmd \
 && rm -rf /var/lib/apt/lists/*

# Add unicode support
RUN locale-gen en_US.UTF-8
ENV LANG 'en_US.UTF-8'
ENV LANGUAGE 'en_US:en'

# Create symlink for executable
RUN ln -s /usr/games/steamcmd /usr/bin/steamcmd

RUN mkdir -p /root/.steam/sdk64 &&  \
    ln -s /root/.steam/steamcmd/linux64/steamclient.so /root/.steam/sdk64/steamclient.so

FROM install-steam as install-terraria

WORKDIR /root/terraria-server

RUN apt-get update && \
    apt-get install -y \
    curl \
    unzip \
    libicu-dev \
    jq

# install terraria
ARG TERRARIA_VERSION=1449
ARG TERRARIA_ARTIFACT_URL=https://terraria.org/api/download/pc-dedicated-server/terraria-server-$TERRARIA_VERSION.zip
ARG TERRARIA_ZIP_FILENAME=terraria-server.zip

WORKDIR /root/terraria-server

RUN mkdir terraria

ADD $TERRARIA_ARTIFACT_URL $TERRARIA_ZIP_FILENAME

FROM install-terraria as install-tmod

WORKDIR /root/terraria-server

RUN unzip $TERRARIA_ZIP_FILENAME -d terraria-zip && \
    mv terraria-zip/$TERRARIA_VERSION/Linux/* terraria && \
    mv terraria-zip/$TERRARIA_VERSION/Windows/serverconfig.txt terraria/serverconfig-default.txt && \
    chmod +x terraria/TerrariaServer && \
    chmod +x terraria/TerrariaServer.bin.x86_64 && \
    rm $TERRARIA_ZIP_FILENAME && \
    rm -rf terraria-zip

ARG TMOD_VERSION
ARG TMOD_ARTIFACT_URL=https://github.com/tModLoader/tModLoader/releases/download/$TMOD_VERSION/tModLoader.zip
ARG TMOD_ZIP_FILENAME=tmodloader.zip

ENV CONFIGPATH=/config
ENV CONFIG_FILENAME="serverconfig.txt"
ENV MOD_PATH=/root/terraria-server/tModLoader/Mods
ENV STEAM_LOBBY=friends

# install tmod
ADD $TMOD_ARTIFACT_URL $TMOD_ZIP_FILENAME

RUN unzip $TMOD_ZIP_FILENAME -d tModLoader && \
    rm $TMOD_ZIP_FILENAME && \
    chmod u+x tModLoader/start-tModLoaderServer.sh && \
    chmod u+x tModLoader/DedicatedServerUtils/manage-tModLoaderServer.sh && \
    chmod u+x tModLoader/LaunchUtils/ScriptCaller.sh

RUN tModLoader/DedicatedServerUtils/manage-tModLoaderServer.sh update

COPY install_mods.sh install_mods.sh

COPY bootstrap.sh bootstrap.sh

ENTRYPOINT [ "/bin/bash", "bootstrap.sh" ]
