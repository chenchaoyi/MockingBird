MockingBird
===========

MockingBird is collection of tools to Mock external server for server backend. It will have feature mocking server and performance mocking server


## Dummy Server ##

For Palantiri use. This relies on Stan's palantiri go lib (https://git.gree-dev.net/stanislav-vishnevski/go-overseer/)

To make it run, you can either

######  set your GOPATH by 
- `echo "export GOPATH=\$HOME/go" >> ~/.bash_profile`
- `mkdir -p $GOPATH/src/git.gree-dev.net/stanislav-vishnevski && cd $GOPATH/src/git.gree-dev.net/stanislav-vishnevski`
- `git clone git@git.gree-dev.net:stanislav-vishnevski/go-overseer.git`
- do `git clone git@git.gree-dev.net:lay-zhu/MockingBird.git` wherever you want
- in your clone folder do `cd /#{FOLDERPATH}/MockingBird/dummy_server && GOPATH=~/#{FOLDERPATH}/MockingBird/dummy_server:$GOPATH CGO_ENABLED=0 go run server.go -amount 200`

###### download and execute binary directly
- to be updated