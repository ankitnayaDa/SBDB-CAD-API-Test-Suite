# SBDB-CAD-API-Test-Suite
API test automation suite Californian Institute of Technology’s (CIT) SBDB Close-Approach Data API

# Run test cases with Below command
./runTestSuit.sh

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
