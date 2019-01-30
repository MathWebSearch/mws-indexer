# mws-indexer

[![Build Status](https://travis-ci.org/MathWebSearch/mws-indexer.svg?branch=master)](https://travis-ci.org/MathWebSearch/mws-indexer)

A script that can safely update the index used by MathWebSearch. 
Written in golang, and packaged via docker. 

## Building

The code can be built using standard `golang` tools. 
Furthermore, a Makefile is provided which can be used as:

```bash
make
```

A static `mwsupdate` executable (both for the current architecture and cross-compiled for others) can then be found in the `out` directory. 

## Update Procedure

```
Usage of mwsupdate:
  -docker-label string
        label of docker container to restart
  -harvest-dir string
        Path to harvest directory (default "/data/")
  -index-dir string
        Path to index directory (default "/index/")
  -mws-index string
        mws-index executable (default "/mws/bin/mws-index")
```

The script performs the following update mechanism:

0. Run `git pull` inside the harvest directory (iff it is a git repository)
1. Generate a new index from the harvest directory, using the `mws-index` executable. 
2. If successfull, replace the index directory with the newly generated one
3. If successfull and if a `docker-label` is provided, restart all docker containers with the given label. 

## Dockerfile

This script is intended to be used from inside of docker. 
It is available as the automated build [mathwebsearch/mws-indexer](https://hub.docker.com/r/mathwebsearch/mws-indexer) on DockerHub. 

It can be run as follows:

```bash
    docker run -t -i --rm -e MWS_DOCKER_LABEL="mws-container-label" -v data-volume:/data/ -v index-volume:/index/ -v /var/run/docker.sock:/var/run/docker.sock mathwebsearch/mws-indexer
```

## LICENSE

Licensed under GPL 3.0. 
