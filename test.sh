#!/usr/bin/env bash


copyFilesToServer(){
   go get -v -u github.com/Srikrishnabh93/go-junit-report
   cat FinalTestResult.log | go-junit-report > result.xml
   sleep 2

   echo "--------------------------------------------------"
   echo "Current working directory="`pwd`
   echo "--------------------------------------------------"
   echo "Files in current working directory:-"
   ls -lrt
   echo "---------------------------------------------------"

   mkdir tmp_results
   cp FinalTestResult.log tmp_results/
   cp result.xml tmp_results/
   cd tmp_results
   echo "--------------------------------------------------"
   echo "Current working directory="`pwd`
   echo "--------------------------------------------------"
   echo "Files in current working directory:-"
   ls -lrt
   echo "---------------------------------------------------"
}

runGoTest(){
    go test -v tests/* -p 1 2>&1 |  tee -a FinalTestResult.log
}

runDocker() {
    runGoTest
    copyFilesToServer
}
$1

