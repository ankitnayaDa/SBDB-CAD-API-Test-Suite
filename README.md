# SBDB-CAD-API-Test-Suite
This repository contains API test automation suite for Californian Institute of Technology’s (CIT) SBDB Close-Approach Data API.

# Understanding API test automation Framework
Test automation suite is written in go and use `go test` to run the test cases.


Tests and source code are run inside a Docker container .`Docker` needs to be installed on test machine.

Tests output a visual test report outside the Docker container

This repository contains below folders



```sh
utils/ : contains utility functions
runTestSuit.sh : to run the test automation suite
libs/  : contains test case related functions
types/ : contains types , variable and constant
Dockerfile : text document that contains all the commands a user could call on the command line to assemble an image
tests/  : contains test cases
test.sh : this scripts to execute test and generate report in container
go.mod  : collection of Go packages
go.sum  : go.sum file lists down the checksum of direct and indirect dependency required
test_report.html : visual output test results 
```

# Getting Started With This Project
To get started, clone this repo, and run `./runTestSuit.sh`.

```sh
git clone https://github.com/ankitnayaDa/SBDB-CAD-API-Test-Suite.git
cd SBDB-CAD-API-Test-Suite
./runTestSuit.sh
```
The aforementioned command outputs an test report HTML file in the same location.  
`test_report.html`

## Understanding HTML test report
At the top of the page are general stats related to the test. Below is the title of the test followed by a section called "Test Groups". Tests are organized into these groups which are represented by color-coded squares. A green square indicates that all test in that group passed. A red square indicates that at least one test in that group failed.

To view the tests in a particular test group, click on any of the test group indicators. A selected group indicator will be colored in black. The number of tests allocated per test group can be set with the groupSize command-line flag.

![image](https://user-images.githubusercontent.com/93188283/186396467-8cc8f138-a494-4a2e-84ce-f790f2b61c79.png)

To view the output of a related test, click on the title of a test on the list. If you want to expand all of the test on the list, simultaneously press shift and the test group indicator.

![image](https://user-images.githubusercontent.com/93188283/186396724-aed840e8-e70b-4e4f-8deb-0a8b6c0c2c82.png)

# Test cases covered:
1. Fetch all Close Approach object data for mars and verify if data is received or not  
2. Fetch all Close Approach object data for mars based on specific dates and verify if data is sorted based on date
3. Fetch all Close Approach object data for mars based on specific maximun and minimum distances and verify if data is sorted based on disttance
4. Verify Zero count results in mars for class IEO
5. Fetch all NEAs and comets for mars and verify if data is sorted based on distance
6. Fetch all NEOS for earth and verify if data is sorted based on distance
7. Fetch all numbered-objects for earth with filter based on dates and verify if data is sorted based on date
8. Fetch details for neo 2020 QW3 for earth and verify if data is sorted based on distance
9. Verify user able to view  details for neo 2020 QW3 for earth with limit 2
10. Verify user able to view All NEOS for earth sorted based on distance with diameter and fullname
11. Test for 400 bad request when request contained invalid keywords and/or content (details returned in the JSON payload)
12. Test 405 response when the request used an incorrect method (see the HTTP Request section) 

# Console output 

Console output while running the test suite

```sh
ankit@ankit-VirtualBox:~/SBDB-CAD-API-Test-Suite$ ./runTestSuit.sh 
cadtest
cadtest
Untagged: cad_test:latest
Deleted: sha256:d00f1b8206f7797b4f2d83cae21cc36899c055fe94551d8faa43751b0ea9303d
Deleted: sha256:6786431d32612872a117b55417bf81c681c8405da766ac0eb545df6e8165f13e
Deleted: sha256:fe061ca8cee7b8cee7db35ef16bbaa57efd6465f32da170db282accd39fcf9ea
Deleted: sha256:4810212f510574af50c045ec8cb8835c8fecac6fee0faaa1483d6c67d39a5aa5
Deleted: sha256:9f925297ebaecedac5bec800c19e9eeca79b1837e6b55a4473147a4c4950bb72
Deleted: sha256:ac75d511aba660ab42cbfe537b024ec22935776f9c24881203a0ebb4b5ddeb4e
Deleted: sha256:08ac1bfde098bcb12b3e8726e1ea66610f0d795484e77c01aeb54ccf0561115b
Deleted: sha256:ab2e2173ff48d71c9ae221fa6c9c5a629ab817a9ad3d19f2ecabab68369871f2
Deleted: sha256:e57c1bdc16d7fdc12f9abe2e5119de49a2d09da67cfbb4a2b12a121071227d1a
Deleted: sha256:4b9141e7a5a97af45c33f9649eb8a6c51b86b00571650de9e531feb2b46e5241
Deleted: sha256:b864dd217943f205f25a06c3af9228ccb9e23bd7c399faa2427ccbc5e2747a2f
Deleted: sha256:d0b956f50908aafb38869303230fbe307a2279d61d2367e5ba41c5ba1bbc975f
Deleted: sha256:eb3446d7f772b985889b878cb057044514a380e382a14d4af26fb8e8ef18fdbf
Deleted: sha256:246f58fa5c1d2c0681d29ae0786ca9cbd53ad63f60da9434c7d9b15bf24af7fc
Deleted: sha256:f52aa1bea6dc2ff5124c6720539343c71564c0bab06c4c4af2f0281ec599219e
Deleted: sha256:7c9f83553b3ad16f38559f7f5a76b01e99deac3f52d9169f23a075add5738a7d
Deleted: sha256:b642990d4c2df4c512a11b0ad86e36bd4cdded3c2bad0a3118502beaf894921b
Deleted: sha256:1ff62145660337fecdfc207d32175d168122e88217c154da43c9170acdce057b
Deleted: sha256:89f711f7ff51be5dffbf332a4d4ec9118ee3819dd9a46c60cc42fe2ee6ef1e2b
Deleted: sha256:165b6c100895a8f70b62b1fde9019fa7af9bd507d43a96008a96649fdd16d3c6
Deleted: sha256:0e1124c8449163f3c5d9a4b73b2264c4ad28e69c50a985f5dd42bcc471745512
Deleted: sha256:b12ac1a40236ee92ed51bc2cec1e9942deec0f9a31b23312deccaa95ec1c5765
Deleted: sha256:cc0220a4757bd32a0239dbfab33a8c542273b4bd0002e45067965555bab9d561
Deleted: sha256:66c363de80f3a1442da4a458e16fd3f9c2d5d6b2c53eb8d8a55fcbb3dba66bfc
Deleted: sha256:0ae12b84e6628fc860efa3f6a64392d4fd189797edb2e7911f84dab9406a0bd5
Deleted: sha256:6b5069ca208a93d81f22e6f6f994225dcfe824044ff6f91ab4af058981931cd9
Deleted: sha256:5ad872da1e696cdaa9863703ccfb23e8d591a9fa792235e8862ac2fa4da6590e
Deleted: sha256:ea134531c22bdb6f1483a84ff88ba2792c3f85f36de0c18b3180437563925c2e
Deleted: sha256:b631c15811b562ebf6ee31d1f8ba54374c7c0556abc4f70f0b70ba1ccaab4e77
Deleted: sha256:e317f21b704f04ab56af0982aa7cf4218d70844fa8847b1f212b0efd09899f08
Deleted: sha256:b84a594766881568b2a1ff0d45012c1b96e0f54d9ef70757706fcbbd900829fe
Sending build context to Docker daemon  376.3kB
Step 1/18 : FROM golang:1.16-alpine as download-dependencies
 ---> 7642119cd161
Step 2/18 : RUN apk update &&     apk add git &&     apk add build-base &&     apk add bash
 ---> Running in f7bf0288645b
fetch https://dl-cdn.alpinelinux.org/alpine/v3.15/main/x86_64/APKINDEX.tar.gz
fetch https://dl-cdn.alpinelinux.org/alpine/v3.15/community/x86_64/APKINDEX.tar.gz
v3.15.6-38-g327cfc61c2 [https://dl-cdn.alpinelinux.org/alpine/v3.15/main]
v3.15.6-38-g327cfc61c2 [https://dl-cdn.alpinelinux.org/alpine/v3.15/community]
OK: 15864 distinct packages available
(1/6) Installing brotli-libs (1.0.9-r5)
(2/6) Installing nghttp2-libs (1.46.0-r0)
(3/6) Installing libcurl (7.80.0-r2)
(4/6) Installing expat (2.4.7-r0)
(5/6) Installing pcre2 (10.40-r0)
(6/6) Installing git (2.34.4-r0)
Executing busybox-1.34.1-r3.trigger
OK: 19 MiB in 21 packages
(1/20) Installing libgcc (10.3.1_git20211027-r0)
(2/20) Installing libstdc++ (10.3.1_git20211027-r0)
(3/20) Installing binutils (2.37-r3)
(4/20) Installing libmagic (5.41-r0)
(5/20) Installing file (5.41-r0)
(6/20) Installing libgomp (10.3.1_git20211027-r0)
(7/20) Installing libatomic (10.3.1_git20211027-r0)
(8/20) Installing libgphobos (10.3.1_git20211027-r0)
(9/20) Installing gmp (6.2.1-r1)
(10/20) Installing isl22 (0.22-r0)
(11/20) Installing mpfr4 (4.1.0-r0)
(12/20) Installing mpc1 (1.2.1-r0)
(13/20) Installing gcc (10.3.1_git20211027-r0)
(14/20) Installing musl-dev (1.2.2-r7)
(15/20) Installing libc-dev (0.7.2-r3)
(16/20) Installing g++ (10.3.1_git20211027-r0)
(17/20) Installing make (4.3-r0)
(18/20) Installing fortify-headers (1.1-r1)
(19/20) Installing patch (2.7.6-r7)
(20/20) Installing build-base (0.5-r3)
Executing busybox-1.34.1-r3.trigger
OK: 210 MiB in 41 packages
(1/4) Installing ncurses-terminfo-base (6.3_p20211120-r1)
(2/4) Installing ncurses-libs (6.3_p20211120-r1)
(3/4) Installing readline (8.1.1-r0)
(4/4) Installing bash (5.1.16-r0)
Executing bash-5.1.16-r0.post-install
Executing busybox-1.34.1-r3.trigger
OK: 212 MiB in 45 packages
Removing intermediate container f7bf0288645b
 ---> 65733a9dc927
Step 3/18 : WORKDIR /sbdb-cad-api-test
 ---> Running in 88909259076c
Removing intermediate container 88909259076c
 ---> 1ec4b67342f0
Step 4/18 : RUN mkdir libs/
 ---> Running in bc0988856f40
Removing intermediate container bc0988856f40
 ---> c6e9800fcf04
Step 5/18 : RUN mkdir tests/
 ---> Running in c750b39a37c6
Removing intermediate container c750b39a37c6
 ---> 909f2d3a73cc
Step 6/18 : RUN mkdir types/
 ---> Running in 2695e6ff35ac
Removing intermediate container 2695e6ff35ac
 ---> b78ed2e170ac
Step 7/18 : RUN mkdir utils/
 ---> Running in 8b3d8767e38e
Removing intermediate container 8b3d8767e38e
 ---> c42e485d6944
Step 8/18 : COPY libs/ ./libs
 ---> 9d9a5a076ff8
Step 9/18 : COPY types/ ./types
 ---> 1b4122346296
Step 10/18 : COPY utils/ ./utils
 ---> 8b32c2515fcb
Step 11/18 : COPY tests/ ./tests
 ---> 9858b0e03b0d
Step 12/18 : COPY go.mod ./
 ---> 8af4a7c7a0ac
Step 13/18 : COPY go.sum ./
 ---> 733a751f2f47
Step 14/18 : COPY test.sh ./
 ---> 0374a7b893a0
Step 15/18 : RUN go mod download
 ---> Running in 22808fd76b75
Removing intermediate container 22808fd76b75
 ---> 098cd6c243fb
Step 16/18 : RUN pwd
 ---> Running in ec79ce74f2bd
/sbdb-cad-api-test
Removing intermediate container ec79ce74f2bd
 ---> e1dc60e41347
Step 17/18 : RUN ls
 ---> Running in 92e5bcbb74dd
go.mod
go.sum
libs
test.sh
tests
types
utils
Removing intermediate container 92e5bcbb74dd
 ---> a8ebc7d3a5dc
Step 18/18 : ENTRYPOINT [ "/bin/bash", "test.sh", "runDocker" ]
 ---> Running in 29a3985fe7d2
Removing intermediate container 29a3985fe7d2
 ---> a1aa092f8946
Successfully built a1aa092f8946
Successfully tagged cad_test:latest
total 60
-rw-rw-r--    1 root     root         33429 Aug 22 10:32 go.sum
-rw-rw-r--    1 root     root           198 Aug 22 10:32 go.mod
-rwxrwxr-x    1 root     root           998 Aug 23 13:00 test.sh
drwxr-xr-x    1 root     root          4096 Aug 24 11:03 libs
drwxr-xr-x    1 root     root          4096 Aug 24 11:03 utils
drwxr-xr-x    1 root     root          4096 Aug 24 11:03 types
drwxr-xr-x    1 root     root          4096 Aug 24 11:03 tests
go: downloading github.com/vakenbolt/go-test-report v0.9.3
go: downloading github.com/spf13/cobra v1.1.3
go: downloading github.com/spf13/cobra v1.5.0
go: downloading github.com/spf13/pflag v1.0.5
go get: added github.com/spf13/cobra v1.5.0
go get: added github.com/vakenbolt/go-test-report v0.9.3
total 16
-rw-rw-r--    1 root     root          1037 Aug 23 12:56 Check_For_Error_Response_test.go
-rw-rw-r--    1 root     root          4789 Aug 24 11:02 Get_Cad_for_Mars_test.go
drwxrwxr-x    8 1000     1000          4096 Aug 24 11:02 tmp_results
/sbdb-cad-api-test/tests
--------------------------------------------------
Running Tests.....
[go-test-report] finished in 42.639674934s
--------------------------------------------------
Current working directory=/sbdb-cad-api-test/tests
--------------------------------------------------
Files in current working directory:-
total 140
-rw-rw-r--    1 root     root          1037 Aug 23 12:56 Check_For_Error_Response_test.go
-rw-rw-r--    1 root     root          4789 Aug 24 11:02 Get_Cad_for_Mars_test.go
drwxrwxr-x    8 1000     1000          4096 Aug 24 11:02 tmp_results
-rw-r--r--    1 root     root        126912 Aug 24 11:04 test_report.html
---------------------------------------------------
mkdir: can't create directory 'tmp_results': File exists
--------------------------------------------------
Current working directory=/sbdb-cad-api-test/tests/tmp_results
--------------------------------------------------
Files in current working directory:-
total 196
drwxrwxr-x    2 1000     1000          4096 Aug 22 10:32 utils
-rwxrwxr-x    1 1000     1000           231 Aug 22 10:32 runTestSuit.sh
-rw-rw-r--    1 1000     1000         33429 Aug 22 10:32 go.sum
-rw-rw-r--    1 1000     1000           198 Aug 22 10:32 go.mod
drwxrwxr-x    2 1000     1000          4096 Aug 23 10:36 libs
drwxrwxr-x    2 1000     1000          4096 Aug 23 11:02 types
-rw-rw-r--    1 1000     1000           447 Aug 23 12:39 Dockerfile
-rwxrwxr-x    1 1000     1000           998 Aug 23 13:00 test.sh
-rw-rw-r--    1 1000     1000          3711 Aug 24 10:56 README.md
drwxrwxr-x    2 1000     1000          4096 Aug 24 11:02 tests
-rw-r--r--    1 root     root        126912 Aug 24 11:04 test_report.html
---------------------------------------------------
```
