 #!/usr/bin/env bash
 jsonBody='{"headline":"headlinevalue","description":"descriptionvalue"}'
 echo "Sending create request with body $jsonBody"
 echo "start ..."
 curl -i -H "Content-type: application/json" -d "$jsonBody" 'http://localhost:8080/records'
 echo "... end"