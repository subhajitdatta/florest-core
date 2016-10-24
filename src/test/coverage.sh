#!/bin/bash
BDD_TEST_DIR="src/test"
UNIT_TEST_DIR="src"
TEST_CONFIG_DIR="testdata"
TEMP_COVERAGE_FILE="tmp.cov"
COVERAGE_FILE="coverage.cov"
COVERAGE_HTML="coverage.html"

# generate coverage file
echo 'mode: count' > $COVERAGE_FILE #generate coverage file

# coverage for bdd tests
# looks upto given level of depth directory
for dir in $(find $BDD_TEST_DIR -maxdepth 1 -mindepth 1 -type d);
do
  # generate binary
  go test -c `echo $dir|sed 's/src\///'` -covermode=count -coverpkg ./...
  # get test name
  name=`echo $dir|sed 's#.*/##'`
  # run the binary only if successfully generated 
  if [ -f "$name.test" ]
  then
    # run binary
    cp -r "$dir/"$TEST_CONFIG_DIR .
    ./"$name.test" -test.coverprofile $TEMP_COVERAGE_FILE
    # remove illegal main file
    sed -i '/_\/.*main.go/d' $TEMP_COVERAGE_FILE 
    # append individual coverage
    tail -q -n +2 $TEMP_COVERAGE_FILE >> $COVERAGE_FILE
    # remove binary
    rm "$name.test"
  else
    echo "no test file found in $dir"
  fi
done

# coverage for unit tests
# looks upto given level of depth directory, skips files staring with '.' or '_' or 'test'
for dir in $(find $UNIT_TEST_DIR -maxdepth 10 -mindepth 1 -not -path './.*' -not -path '*/_*' -not -path '*/test*' -type d);
do
  # generate binary
  go test -c `echo $dir|sed 's/src\///'` -covermode=count
  # get test name
  name=`echo $dir|sed 's#.*/##'`
  # run the binary only if successfully generated
  if [ -f "$name.test" ]
  then
    # run binary
    ./"$name.test" -test.coverprofile $TEMP_COVERAGE_FILE
    # add full import path, needed for go tool cover
    sed -i -e 's/^/github.com\/jabong\/florest-core\/src\//' $TEMP_COVERAGE_FILE
    # append individual coverage
    tail -q -n +2 $TEMP_COVERAGE_FILE >> $COVERAGE_FILE
    # remove binary
    rm "$name.test"
  else
    echo "no test file found in $dir"
  fi
done

# remove temp files and generate html
rm -r $TEST_CONFIG_DIR
# remove temp coverage
rm $TEMP_COVERAGE_FILE 
# create bin if not present
mkdir -p bin
mv $COVERAGE_FILE bin/
cd bin
# generate html file
go tool cover -html=$COVERAGE_FILE -o $COVERAGE_HTML

