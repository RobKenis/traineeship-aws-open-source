# Getting started with Elastic Beanstalk

## Creating a Beanstalk environment
Deploy [beanstalk.yaml](beanstalk.yaml) using the CloudFormation console.

## Deploying an application
Build the application by executing `build.sh` or by building the application and making a zip file yourself.
```shell
./build.sh
```
This will create an artifact called **beanstalk.zip**, it contains 2 important files: the jar file to execute
and a Procfile to tell beanstalk which processes to execute on startup. This works in a similar way as for example
dyno's on Heroku.

Navigate to the Beanstalk console and under your environment, upload the zip file, enter a version and pray for the best.
