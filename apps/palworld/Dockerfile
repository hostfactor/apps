FROM ghcr.io/hostfactor/apps/base-builder as builder

ENV PORT=8211

RUN $HOME/steamcmd/steamcmd.sh +force_install_dir $HOME/app +login anonymous +app_update 2394010 +quit

RUN mkdir -p app/Pal/Saved/Config/LinuxServer

WORKDIR ./app

RUN cp DefaultPalWorldSettings.ini Pal/Saved/Config/LinuxServer/PalWorldSettings.ini

COPY start.sh start.sh

CMD ["./start.sh"]
