#! /bin/bash
#chmod +x!! must be done
#*/5 * * * * /home/juno/git/goFastCgi/goFastCgi/cronparagraphshandler.bash
#*/2 * * * * flock -n /tmp/cronparagraphshandler.lock -c /home/juno/git/goFastCgiLight/goFastCgiLight/cronparagraphshandler.bash

cd /home/juno/git/goFastCgiLight/goFastCgiLight
bin/paragraphshandler -locale=fi_FI -themes=porno -quant=3000