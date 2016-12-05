#!/usr/bin/bash
export EMAIL=xDarkicex@gmail.com
export SMTPHOST=smtp.gmail.com
export SMTPPORT=587
export SMTPPASSWORD='Vh402152Go!'
export ENV=production
sass --watch ./app/assets/stylesheets/:./public/assets/stylesheets/ --style compressed &
PID=$!
sleep 5
kill $PID

tsc --outDir ./public/assets/scripts/application/ ./app/assets/typescripts/application/*.ts
tsc --outDir ./public/assets/scripts/users/ ./app/assets/typescripts/users/*.ts
tsc --outDir ./public/assets/scripts/blog/ ./app/assets/typescripts/blog/*.ts
tsc --outDir ./public/assets/scripts/examples/ ./app/assets/typescripts/examples/*.ts
tsc --outDir ./public/assets/scripts/ ./app/assets/typescripts/*.ts

/home/ubuntu/bin/PortfolioGo