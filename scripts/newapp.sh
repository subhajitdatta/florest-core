#!/bin/sh

florest_path=$1
project_path=${2%/}

project_name=$(basename $project_path)

# create project dir, if not present
mkdir -p $project_path

# exit if failed to create project dir
if [ ! -d $project_path ]; then
  echo "failed to create directory:$project_path"
  exit 1
fi

# Copy the new app source code
echo building new app
cp -r -f "$florest_path/_newApp"/* $project_path

# Copy the libs
echo copying dependent libs
cp -r -f "$florest_path/_libs"/* "$project_path/_libs/"
mkdir -p "$project_path/_libs/src/github.com/jabong/florest-core"
cp -r -f "$florest_path/src" "$project_path/_libs/src/github.com/jabong/florest-core/"

config_dir="$project_path/config/$project_name"
swagger_file="$project_path/src/hello/swagger.go"
log_config_file="$project_path/config/logger/logger.json"

# Rename the newly created project to the project name published
mv "$project_path/config/newApp" $config_dir

# Replace the APP_NAME with the project name in conf file
sed -i'' -e "s/{APP_NAME}/$project_name/g" "$config_dir/conf.json"

# Replace the APP_NAME with the project name in swagger file
sed -i'' -e "s/{APP_NAME}/$project_name/g" $swagger_file

# Replace the APP_NAME with the project name in logger conf file
sed -i'' -e "s/{APP_NAME}/$project_name/g" $log_config_file

# Test the app if asked
if [ $# -eq 3 ] && [ $3 = "testapp" ]; then
  cd $project_path
  echo "starting test for new app $project_path"
  make test
  if [ $? -ne 0 ]; then
    exit 1
  fi
  # test is successful, now start the app and hit a request
  echo "run and verify new app with a request"
  make deploy
  if [ $? -ne 0 ]; then
    exit 1
  fi
  # compilation successful, start the app
  cd bin
  ./$project_name &
  # add some sleep for server to start
  sleep 15 
  curl "http://localhost:8080/$project_name/v1/hello" > hello.log
  diff "$florest_path/src/test/apptestdata/hello.log" hello.log
fi
