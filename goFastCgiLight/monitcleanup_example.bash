cd /home/juno/git/goFastCgiLight/goFastCgiLight

pgrep elabqueue || pgrep cleanup || bin/cleanup
/usr/bin/find www -type d -empty -delete

