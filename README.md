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
  -harvest-dir string
        Path to harvest directory (default "/data/")
  -harvests2json string
        harvests2json executable (default "/mws/bin/harvests2json")
  -index-dir string
        Path to index directory (default "/index/")
  -mws-index string
        mws-index executable (default "/mws/bin/mws-index")
  -tema
        Generate indexes for tema-search
  -tema-index-dir string
        Path to tema index directory (default "/temaindex/")
```

The script performs the following update mechanism:

0. Run `git pull` inside the harvest directory (iff it is a git repository)
1. Generate a new index from the harvest directory, using the `mws-index` executable. 
2. If successful, replace the index directory with the newly generated one
3. If requested, repeat steps 1 + 2 for a tema-search index. 

## Dockerfile

This script is intended to be used from inside of docker. 
It is available as the automated build [mathwebsearch/mws-indexer](https://hub.docker.com/r/mathwebsearch/mws-indexer) on DockerHub. 

It can be run as follows for a plain mathwebsearch:

```
docker run -t -i --rm -v data-volume:/data/ -v index-volume:/index/ -v /var/run/docker.sock:/var/run/docker.sock mathwebsearch/mws-indexer
```

It can be run as follows for temasearch:

```
docker run -t -i --rm -v data-volume:/data/ -v index-volume:/index/ -v tema-index-volume:/temaindex/ -v /var/run/docker.sock:/var/run/docker.sock mathwebsearch/mws-indexer --tema
```


## LICENSE

Licensed under GPL 3.0. 
