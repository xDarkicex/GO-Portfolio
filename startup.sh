#!/usr/bin/bash
/home/ubuntu/.rvm/gems/ruby-2.3.0/bin/sass --watch ./app/assets/stylesheets/:./public/assets/stylesheets/ --style compressed &
# PID=$!
# sleep 5
# kill $PID

/usr/local/bin/tsc --outDir ./public/assets/scripts/application/ ./app/assets/typescripts/application/*.ts
/usr/local/bin/tsc --outDir ./public/assets/scripts/users/ ./app/assets/typescripts/users/*.ts
/usr/local/bin/tsc --outDir ./public/assets/scripts/blog/ ./app/assets/typescripts/blog/*.ts
/usr/local/bin/tsc --outDir ./public/assets/scripts/examples/ ./app/assets/typescripts/examples/*.ts
/usr/local/bin/tsc --outDir ./public/assets/scripts/ ./app/assets/typescripts/*.ts

sudo /home/ubuntu/bin/PortfolioGo & 
SPID=$!
echo $SPID > pid.txt