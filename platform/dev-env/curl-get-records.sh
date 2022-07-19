 #!/usr/bin/env bash
 echo "Sending get request"
 echo "start ..."
 curl -i -H "Content-type: application/json" 'http://localhost:8080/records'
 echo 
 echo "... end"