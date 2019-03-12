#!/bin/bash
# resource-marathon-deploy set-pipeline.sh

fly -t ci set-pipeline -p resource-marathon-deploy -c pipeline.yml --load-vars-from ../../../../../.credentials.yml
