#!/bin/bash
TEST_DIR="src/test"
TEST_CONFIG_DIR="testdata"
TEMP_COVERAGE_FILE="tmp.cov"
COVERAGE_FILE="coverage.cov"
COVERAGE_HTML="coverage.html"

echo 'mode: count' > $COVERAGE_FILE #generate coverage file
for dir in $(find $TEST_DIR -maxdepth 1 -mindepth 1 -type d);
do
  go test -c `echo $dir|sed 's/src\///'` -covermode=count -coverpkg ./...
  name=`echo $dir|sed 's#.*/##'` #test name
  if [ -f "$name.test" ]
  then
    cp -r "$dir/"$TEST_CONFIG_DIR .
    ./"$name.test" -test.coverprofile $TEMP_COVERAGE_FILE
    sed -i '/_\/.*main.go/d' $TEMP_COVERAGE_FILE # remove illegal main file
    tail -q -n +2 $TEMP_COVERAGE_FILE >> $COVERAGE_FILE # append individual coverage
    rm "$name.test"
  else
    echo "no test file found in $dir"
  fi
done
rm -r $TEST_CONFIG_DIR
rm $TEMP_COVERAGE_FILE # remove temp coverage
mkdir -p bin
mv $COVERAGE_FILE bin/
cd bin
go tool cover -html=$COVERAGE_FILE -o $COVERAGE_HTML

