#! /bin/sh

project_path=${1%/}

project_name=$(basename $project_path)

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
