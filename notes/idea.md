# es-curator

This solution shall give the user a platform to:
    * collect any type of record 
    * any logged in user can write review for a record
    * privileged users now named curators can choose review to be pinned visibly on the record
    * any logged in user can link the record to any other 
        * review/record on this platform
        * URL
    * curators can pin those links, curate them

## Techincal decisions
* frontend: react with serverside rendering
* backend: go or rust with some db to choose (docuemnt?)
* user authentication: use some ready to work SaaS or keycloak?

## Coding/architecture guidelines and principles
* we will use semantic versioning with feature branches and squash commits
* services run in a container
