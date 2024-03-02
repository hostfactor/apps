#!/bin/bash

if [ -z "${TMOD_MODS}" ]; then
  echo "No mods listed"
else
  ./install_mods.sh $TMOD_MODS
fi

tModLoader/DedicatedServerUtils/manage-tModLoaderServer.sh update -u anonymous

./tModLoader/start-tModLoaderServer.sh --update -u anonymous -config "$CONFIGPATH/$CONFIG_FILENAME" -modpath "$MOD_PATH" -lobby "$STEAM_LOBBY" -nosteam
