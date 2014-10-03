#! /bin/bash

rm ../www/fi_FI/porno/test.com/index.html
rm ../www/fi_FI/porno/test.com/test.html.gz
rm ../www/fi_FI/porno/test.com/test2.html.gz

export GOPATH=$GOPATH:/home/juno/git/goFastCgiLight/goFastCgiLight
cd createpage 
go test -v
#cd ../htmlfileexist
#go test -v


