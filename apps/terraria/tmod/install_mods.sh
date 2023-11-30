#!/usr/bin/env sh

enabled_path=/root/terraria-server/tModLoader/Mods/enabled.json

mkdir -p "$(dirname $enabled_path)"

move_mod () {
  filename=$(basename "$1")
  mv "$1" /root/terraria-server/tModLoader/Mods/"$filename"
}

build_enabled_content() {
  enabled_content="[]"
  find /data/steamMods/steamapps/workshop/content/1281930 -name "*.tmod" | while read -r file
  do
    move_mod "$file"
  done

  for file in /root/terraria-server/tModLoader/Mods/*.tmod; do
    filename=$(basename "$file")
    filename_no_ext="${filename%.*}"
    enabled_content="$(echo "$enabled_content" | jq '. += ["'"$filename_no_ext"'"]')"
  done

  echo "$enabled_content"
}

steamcmd_install="steamcmd +force_install_dir /data/steamMods +login anonymous"

for var in "$@"
do
  steamcmd_install="$steamcmd_install +workshop_download_item 1281930 $var"
done

steamcmd_install="$steamcmd_install +quit"

echo "Installing mods"

$steamcmd_install

echo "Done!"

enabled_content=$(build_enabled_content)
echo "$enabled_content" > $enabled_path
