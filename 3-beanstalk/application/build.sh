#!/usr/bin/env bash

./gradlew build

zip -r -j beanstalk.zip build/libs/beanstalk-demo-0.0.1-SNAPSHOT.jar Procfile
