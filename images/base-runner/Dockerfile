FROM debian:stable-slim


ARG HOME=/home/hostfactor
ARG USER=hostfactor
ARG GROUP=hostfactor
ARG UID=1000
ARG GID=1000

ENV HOME=$HOME
ENV USER=$USER
ENV GROUP=$GROUP
ENV UID=$UID
ENV GID=$GID

RUN groupadd -g $GID $GROUP && \
    useradd -M -s /bin/bash -d $HOME -g $GID -u $UID $USER

RUN mkdir -p $HOME/.steam \
 && ln -s /root/.local/share/Steam/steamcmd/linux32 $HOME/.steam/sdk32 \
 && ln -s /root/.local/share/Steam/steamcmd/linux64 $HOME/.steam/sdk64 \
 && ln -s /root/.steam/sdk32/steamclient.so $HOME/.steam/sdk32/steamservice.so \
 && ln -s /root/.steam/sdk64/steamclient.so $HOME/.steam/sdk64/steamservice.so

WORKDIR $HOME

RUN chown -R $UID:$GID .

USER $UID
