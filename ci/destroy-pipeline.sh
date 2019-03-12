#!/bin/bash
# resource-marathon-destroy destroy-pipeline.sh

fly -t ci destroy-pipeline --pipeline resource-marathon-destroy
