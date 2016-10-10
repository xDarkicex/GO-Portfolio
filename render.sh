#!/usr/bin/bash
if [[ -z $2 ]]; then
    pug -p app/views/$1.pug < app/views/$1.pug
else
    pug -O "$2" -p app/views/$1.pug < app/views/$1.pug
fi