#! /bin/bash
#chmod +x!! must be done
#*/5 * * * * /home/juno/git/goFastCgi/goFastCgi/crongofastscript.bash
#*/5 * * * * flock -n /tmp/crongofastscript.lock -c /home/juno/git/goFastCgiLight/goFastCgiLight/crongofastscript.bash

cd /home/juno/git/goFastCgiLight/goFastCgiLight
#pgrep elabqueue && echo `date >>/tmp/elabqueue`
#bin/paragraphshandler -locale=fi_FI -themes=finance -quant=500
#pgrep elabqueue || bin/elabqueue
#pgrep elabqueue || bin/cleanupspace -hits=1 -created=500

#bin/elabqueue
#bin/cleanupspace -hits=1 -created=300
#pgrep elabqueue || bin/elabqueue && bin/cleanupspace -hits=1 -created=300 && bin/orphans -updated=1800
pgrep elabqueue || pgrep cleanup || bin/elabqueue
pgrep elabqueue || pgrep cleanup || bin/newlimitdb

#bin/newdomain -locale=fi_FI -themes=porno -domain=tissit.tv -expire=3600
#bin/newdomain -locale=fi_FI -themes=porno -expire=120

