#!/usr/bin/env bash

rm -rf test_report.html
docker stop cadtest
docker container rm cadtest
docker rmi cad_test
docker build -t cad_test .
docker run -v $(pwd):/sbdb-cad-api-test/tests/tmp_results/ --name cadtest cad_test:latest

