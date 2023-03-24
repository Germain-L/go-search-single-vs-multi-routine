#!/bin/sh

async=$(./async $1 $2 | grep 'Elapsed')

sync=$(./sync $1 $2 | grep 'Elapsed')


echo $async 'vs' $sync