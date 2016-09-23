#!/bin/bash

loglevel=$1
file=$2

sed -i'' -e "s/{LOGLEVEL}/$loglevel/g" $file
