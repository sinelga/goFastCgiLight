apt-get autoremove --purge apt-xapian-index   !! We dont neet it?!
apt-get autoremove --purge 

 echo 524288 > /proc/sys/fs/inotify/max_user_watches
 vi /etc/sysctl.conf
 fs.inotify.max_user_watches=1900000
 vm.overcommit_memory = 1  ??
 sysctl -p


Check nginx 
strace -p 3544 -p 3545 -p 3546 -p 3547 2>&1 | grep gz

ps -eo %mem,rss,pid,args --sort rss | tail -10

dd if=/dev/zero of=/swapfile bs=1024 count=128k

mkswap /swapfile
swapon /swapfile
swapon -s
vi /etc/fstab
/swapfile       none    swap    sw      0       0 

vm.swappiness = 10  in /etc/sysctl.conf  "sysctl -p"
cat /proc/sys/vm/swappiness


GOPATH=$GOPATH:/home/juno/git/goFastCgiLight/goFastCgiLight go test -v


bin/syncpushdomaindb -locale=fi_FI -themes=porno

ZRANGEBYSCORE pushdomains -inf +inf withscores  LIMIT 0 10000

ZADD pushdomains 1 pilluseksi.com
ZADD pushdomains 1 pilluseksi.info



stop gofast && start gofast && stop nginx && start nginx



