#!/bin/bash
# resource-marathon-deploy update_concourse.sh

fly -t ci set-pipeline -p hello-go -c ci/pipeline.yml --load-vars-from ci/.credentials.yml
