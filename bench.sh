#!/bin/sh

async=$(./bin/async $1 $2 | grep 'Elapsed')

sync=$(./bin/sync $1 $2 | grep 'Elapsed')


echo $async 'vs' $sync