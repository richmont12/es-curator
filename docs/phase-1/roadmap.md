# phase 1

## phase 1.1 simple running sample applications
* use of some identity provider as SaaS or self deployed
* access to a document db, self deployed or SaaS
* a go backend available as container image
    * authentication with identity provider
    * loads data from some document db
* a react frontend available as container image
    * requesting data from go backend
    * authentication with identity provider

## phase 1.2 easy to privison infrastructure
* some tooling to easily create everything that we can deploy on
    * decide: k8s or anything docker swarm like?
    * also locally available

## phase 1.3 running deployment
* some tooling to easily get the platform based on the container images up
    * easy configurable