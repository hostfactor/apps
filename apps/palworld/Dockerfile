FROM ghcr.io/hostfactor/apps/base

ARG HOME=/home/hostfactor

ARG USER=hostfactor
ARG GROUP=hostfactor
ARG UID=1000
ARG GID=1000

RUN groupadd -g $GID $GROUP && \
    useradd -M -s /bin/bash -d $HOME -g $GID -u $UID $USER

RUN steamcmd +force_install_dir $HOME/app +login anonymous +app_update 2394010 +quit

WORKDIR $HOME/app

RUN chown -R $UID:$GID .

USER $UID
