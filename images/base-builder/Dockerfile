FROM steamcmd/steamcmd

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

RUN mkdir -p $HOME/.steam/sdk32 && mkdir -p $HOME/.steam/sdk64 \
 && mv /root/.local/share/Steam/steamcmd/linux32/* $HOME/.steam/sdk32 \
 && mv /root/.local/share/Steam/steamcmd/linux64/* $HOME/.steam/sdk64

RUN chown -R $UID:$GID $HOME

WORKDIR $HOME

USER $USER

ENTRYPOINT []
