apt-get autoremove --purge apt-xapian-index   !! We dont neet it?!
apt-get autoremove --purge 


Check nginx 
strace -p 3544 -p 3545 -p 3546 -p 3547 2>&1 | grep gz