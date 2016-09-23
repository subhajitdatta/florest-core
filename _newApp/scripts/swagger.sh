#!/bin/bash

cd bin
# iterate throug all packages
for i in $(echo $1|sed 's/,/ /g')
do
swagger -apiPackage=$i -mainApiFile="$i/swagger.go" -format="swagger"
done

# iterate through all valid json files
for i in ./*/index.json;
do
sed -i'' -e 's/"bool"/"boolean"/g' "$i" # replace bool with boolean
name=`echo "$i"|sed 's/\/index.json//'|sed 's/\.\///'`
path=`grep "basePath" $i  | sed 's/.*://' | sed 's/"//g' | sed 's/,//' | sed 's/ //'| sed 's/\/.*\///'`
mkdir -p "swagger/$path/$name"
mv "$i" "swagger/$path/$name/apidocs.json" # move to version dir
rm -r "$name" # remove directory
done

# remove the index file
rm index.json

