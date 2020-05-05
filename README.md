- scheduler will be only thing we can assume has access to vault..
- provision will go through scheduler
- there will a scheduler per cluster
- there will be an sqs queue per cluster
- we need to figure out how to authenticate scheduler with s3/sqs outside of AWS
- the scheduler will be given a plan (via s3 through the queue) specific to a namespace and the cluster the scheduler lives in
- the scheduler will need to translate some sort of templating language for vault secrets
- the scheduler will need to be able to deploy to the kubernetes cluster it is in given the deployment plan
- the scheduler will need to be able to trigger an azure function with a response of success/failure 