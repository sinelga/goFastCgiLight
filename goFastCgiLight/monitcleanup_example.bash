cd /home/juno/git/goFastCgiLight/goFastCgiLight

pgrep elabqueue || pgrep cleanup || bin/cleanup
pgrep find || /usr/bin/find www -type d -empty -delete

