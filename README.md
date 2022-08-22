# SBDB-CAD-API-Test-Suite
API test automation suite Californian Institute of Technology’s (CIT) SBDB Close-Approach Data API

# Run test cases with Below command
./runTestSuit.sh

# Test cases covered:
1. Verify user able to view All Close Approach object data for mars 
2. Verify user able to view All Close Approach object data for mars based on specific dates with sort based on date
3. Verify user able to view All Close Approach object data for mars based on specific maximun and minimum distances with sort based on dist
4. Verify Zero count results in mars for class IEO
5. Verify user able to view All NEAs and comets for mars based on distance
6. Verify user able to view All NEOS for earth sorted based on distance
7. Verify user able to view All numbered-objects for earth sorted based on distance with filter based on dates
8. Verify user able to view details for neo 2020 QW3 for earth sorted based on distance
9. Verify user able to view  details for neo 2020 QW3 for earth with limit 2
10. Verify user able to view All NEOS for earth sorted based on distance with diameter and fullname
11. Test for 400 bad request when request contained invalid keywords and/or content (details returned in the JSON payload)
12. Verify 405 response when the request used an incorrect method (see the HTTP Request section) 
