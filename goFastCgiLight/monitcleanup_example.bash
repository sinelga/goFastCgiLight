cd /home/juno/git/goFastCgiLight/goFastCgiLight

/sbin/stop gonotifylight
su juno -c "pgrep elabqueue || pgrep cleanup || bin/cleanup"
su juno -c "pgrep find || /usr/bin/find www -type f -mtime -1 -size -10k -delete
su juno -c "pgrep find || /usr/bin/find www -type d -empty -delete"
su juno -c "/usr/bin/touch cleanupspace.csv"
/sbin/start gonotifylight




