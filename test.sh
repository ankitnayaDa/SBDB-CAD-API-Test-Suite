#!/usr/bin/env bash


copyFilesToServer(){
   sleep 2

   echo "--------------------------------------------------"
   echo "Current working directory="`pwd`
   echo "--------------------------------------------------"
   echo "Files in current working directory:-"
   ls -lrt
   echo "---------------------------------------------------"

   mkdir tmp_results
   mv test_report.html tmp_results/
   cd tmp_results
   echo "--------------------------------------------------"
   echo "Current working directory="`pwd`
   echo "--------------------------------------------------"
   echo "Files in current working directory:-"
   ls -lrt
   echo "---------------------------------------------------"
}

runGoTest(){
    ls -ltr
    go get -u github.com/vakenbolt/go-test-report/
    cd tests/
    ls -ltr
    pwd
    echo "--------------------------------------------------"
    echo "Running Tests....."
    go test -json | go-test-report
}

runDocker() {
    runGoTest
    copyFilesToServer
}
$1

