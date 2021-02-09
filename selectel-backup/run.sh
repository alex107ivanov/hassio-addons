#!/bin/bash
set -e

echo "[Info] Starting Selectel Backup docker!"

CONFIG_PATH=/data/options.json
cryptkey=$(jq --raw-output ".cryptkey" $CONFIG_PATH)
username=$(jq --raw-output ".username" $CONFIG_PATH)
password=$(jq --raw-output ".password" $CONFIG_PATH)
path=$(jq --raw-output ".path" $CONFIG_PATH)
deleteolderthan=$(jq --raw-output ".deleteolderthan" $CONFIG_PATH)

today=`date +%Y%m%d%H%M%S`
hassconfig="/config"
hassbackup="/backup"
archfile="homeassistant_backup_$today.tar.gz"
archpath="$hassbackup/$archfile"

echo "[Info] Starting backup creating $archpath"
cd $hassconfig
tar zcf ${archpath} . -x ./*.db ./*.db-shm ./*.db-wal
echo "[Info] Finished archiving configuration"

echo "[Info] trying to upload $archpath"
/opt/upload.sh ${archpath} ${cryptkey} ${username} ${password} ${path} || exit 2

if [ "${#deleteolderthan}" -gt "0" ]; then
	echo "[Info] Deleting files older than $deleteolderthan days"
	find $hassbackup/homeassistant_backup* -mtime +$deleteolderthan -exec rm {} \;
fi

echo "[Info] Finished backup"
