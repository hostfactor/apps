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

FROM install-steam as install-terraria

WORKDIR /root/terraria-server

RUN apt-get update && \
    apt-get install -y \
    curl \
    unzip \
    libicu-dev

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

ARG TMOD_VERSION=2023.08.3.4
ARG TMOD_ARTIFACT_URL=https://github.com/tModLoader/tModLoader/releases/download/v$TMOD_VERSION/tModLoader.zip
ARG TMOD_ZIP_FILENAME=tmodloader.zip

# install tmod
ADD $TMOD_ARTIFACT_URL $TMOD_ZIP_FILENAME

RUN unzip $TMOD_ZIP_FILENAME -d tModLoader && \
    rm $TMOD_ZIP_FILENAME && \
    chmod u+x tModLoader/start-tModLoaderServer.sh && \
    chmod u+x tModLoader/LaunchUtils/ScriptCaller.sh

RUN steamcmd +force_install_dir /root/terraria-server/workshop +login anonymous +app_update 1281930 +quit

CMD ["/root/terraria-server/tModLoader/start-tModLoaderServer.sh"]
