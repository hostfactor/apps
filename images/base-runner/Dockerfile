FROM debian:stable-slim

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
