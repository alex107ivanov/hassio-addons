#!/bin/bash

FILE=$1
CRYPTKEY=$2
SUSER=$3
SPASS=$4
SPATH=$5

if [ "$FILE.." == ".." ] || [ "$CRYPTKEY.." == ".." ] || [ "$SUSER.." == ".." ] || [ "$SPASS.." == ".." ] || [ "$SPATH" == ".." ] ; then
  echo " *** Usage: $0 filename cryptkey username password path"
  exit 1
fi

openssl enc -aes128 -salt -in $FILE -out $FILE.enc -e -k $CRYPTKEY
/opt/supload.sh -u $SUSER -k $SPASS $SPATH $FILE.enc && rm $FILE.enc
