FROM ghcr.io/hostfactor/apps/base-builder as builder

RUN $HOME/steamcmd/steamcmd.sh +force_install_dir $HOME/app +login anonymous +app_update 2394010 +quit

RUN mkdir -p app/Pal/Saved/Config/LinuxServer

COPY derp.ini app/Pal/Saved/Config/LinuxServer/PalWorldSettings.ini

CMD ["./app/PalServer.sh"]
