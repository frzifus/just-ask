# just-ask

#### Description
-------
Just-ask is a simple ssh-askpass tool. It's for confirming or
denying the use of an ssh key via ssh-agent. Also it's available for all 
platforms that golang supports.

#### Build
```bash
go get -v github.com:frzifus/just-ask
cd $GOPATH/src/git.klimlive.de/frzifus/just-ask
# linux64
GOOS=linux GOARCH=amd64 go build
# arm
GOOS=linux GOARCH=arm go build
# win64
GOOS=windows GOARCH=amd64 go build
```

#### Install and uninstall (Linux)
**Execute as privileged user**
```bash
# install
$ cp just-ask /usr/local/bin/
$ pushd /usr/local/bin
$ ln -s ssh-askpass /usr/local/bin/just-ask
$ popd

# uninstall
$ rm /usr/local/bin/just-ask
$ unlink /usr/local/bin ssh-askpass
```

