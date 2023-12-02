#!/usr/bin/env sh

enabled_path=/root/terraria-server/tModLoader/Mods/enabled.json
workshop_dir=/data/steamMods/steamapps/workshop/content/1281930

mkdir -p "$(dirname $enabled_path)"

find_tmod_file() {
  tmod_files=$(find "$1" -type f -name *.tmod -exec ls -1rv "{}" +)
  echo $tmod_files | awk '{print $1}'
}

move_tmod_file() {
  tmod_file=$(find_tmod_file "$1")

  if [ -z "$tmod_file" ]; then
    return 0
  fi
  filename=$(basename "$tmod_file")
  mv "$tmod_file" /root/terraria-server/tModLoader/Mods/"$filename"
}

build_enabled_content() {
  enabled_content="[]"
  ls $workshop_dir | while read -r mod_folder; do
    move_tmod_file "$workshop_dir/$mod_folder"
  done

  for mod_folder in /root/terraria-server/tModLoader/Mods/*.tmod; do
    filename=$(basename "$mod_folder")
    filename_no_ext="${filename%.*}"
    enabled_content="$(echo "$enabled_content" | jq '. += ["'"$filename_no_ext"'"]')"
  done

  echo "$enabled_content"
}

steamcmd_install="steamcmd +force_install_dir /data/steamMods +login anonymous"

for var in "$@"; do
  steamcmd_install="$steamcmd_install +workshop_download_item 1281930 $var"
done

steamcmd_install="$steamcmd_install +quit"

echo "Installing mods"

$steamcmd_install

echo "Done!"

ls $workshop_dir | while read -r mod_folder; do
  echo "Using mod $(find_tmod_file "$workshop_dir/$mod_folder")"
done

enabled_content=$(build_enabled_content)
echo "$enabled_content" >$enabled_path
