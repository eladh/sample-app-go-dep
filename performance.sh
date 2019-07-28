#!/usr/bin/env zsh

#--- Variables ---
project=${PWD##*/}
resultFile="timings.txt"
sleepInterval=10

export TIMEFMT="%J  %U user %S system %P cpu %*E total"

runWithProxy() {
    goProxy=${1}
    cmd=${2}
    echo "Setting GoProxy to ${goProxy}"
    echo "Setting GoProxy to ${goProxy}" >> ${resultFile}

    for i in {1..2}
    do
        echo "Ready to perform run ${i}"
        sudo rm -rf $GOPATH/pkg/mod
        rm go.sum
        export GOPROXY=${goProxy}
        { time ${cmd} get . 2>> go.stderr ; } 2>> ${resultFile}
        sleep ${sleepInterval}
    done
}

runWithProxy "" "go"
runWithProxy "https://gocenter.io" "go"
