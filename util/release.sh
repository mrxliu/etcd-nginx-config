#!/usr/bin/env bash
set -e

VERSION="0.1.3"

SCRIPT_PATH="$( cd "$(dirname "$0")" ; pwd -P )/$(basename $0)"
PROJECT_DIR=$(dirname `dirname $SCRIPT_PATH`)

echo "Building version ${VERSION}..."
cd $PROJECT_DIR
goxc -pv=$VERSION -d=pkg

echo "Generating downloads page for version ${VERSION}..."
BASE_URL="http:\/\/download.bentonroberts.com\/etcd-nginx-config\/${VERSION}"
SED_EXPR="s/(etcd-nginx-config_.*)\$/${BASE_URL}\/\1/"
cat pkg/${VERSION}/downloads.md | sed -r $SED_EXPR > downloads.md
