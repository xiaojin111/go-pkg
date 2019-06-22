#!/usr/bin/env bash
set -e
set -u
set -o pipefail

CUR=`dirname $0`

. ${CUR}/build/util.sh

cd $CUR

function update_all() {
    info "[1/3] Updating all git submodules..."
    git submodule update --recursive --remote
    git submodule foreach --recursive git rev-parse HEAD

    echo
    info "[2/3] Update all go packages..."
    go get -u all

    echo
    info "[3/3] Tidy go mod packages..."
    go mod tidy
}

warn "以下操作需要科学上网后才能知晓，确认是否已开启代理?"
select yn in "Yes" "No"; do
    case $yn in
        Yes ) update_all; break;;
        No ) exit 1;;
    esac
done

echo
info "Done"
