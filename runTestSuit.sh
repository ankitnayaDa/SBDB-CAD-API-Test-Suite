#!/usr/bin/env bash

rm -rf FinalTestResult.log result.xml
docker container rm cad_test
docker rmi cad_test
docker build -t cad_test .
docker run -v $(pwd):/sbdb-cad-api-test/tmp_results/ cad_test:latest

