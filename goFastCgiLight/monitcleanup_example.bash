cd /home/juno/git/goFastCgiLight/goFastCgiLight

pgrep cleanup || bin/cleanup
pgrep find || /usr/bin/find www -type d -empty -delete
/usr/bin/touch cleanupspace.csv



