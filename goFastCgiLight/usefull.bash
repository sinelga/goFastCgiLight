apt-get autoremove --purge apt-xapian-index   !! We dont neet it?!
apt-get autoremove --purge 


Check nginx 
strace -p 3544 -p 3545 -p 3546 -p 3547 2>&1 | grep gz

ps -eo %mem,rss,pid,args --sort rss | tail -10

dd if=/dev/zero of=/swapfile bs=1024 count=256k
mkswap /swapfile
swapon /swapfile
swapon -s
vi /etc/fstab
/swapfile       none    swap    sw      0       0 

vm.swappiness = 10  in /etc/sysctl.conf  "sysctl -p"
cat /proc/sys/vm/swappiness

