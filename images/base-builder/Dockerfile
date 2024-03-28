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

ARG HOME=/home/hostfactor

ARG USER=hostfactor
ARG GROUP=hostfactor
ARG UID=1000
ARG GID=1000

RUN groupadd -g $GID $GROUP && \
    useradd -M -s /bin/bash -d $HOME -g $GID -u $UID $USER

RUN mkdir -p $HOME/.steam/sdk64 &&  \
    ln -s $HOME/.steam/steamcmd/linux64/steamclient.so $HOME/.steam/sdk64/steamclient.so

WORKDIR $HOME

RUN chown -R $UID:$GID .

USER $UID
