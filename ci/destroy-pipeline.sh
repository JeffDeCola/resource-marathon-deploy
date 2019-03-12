#!/bin/bash
# resource-marathon-deploy destroy-pipeline.sh

fly -t ci destroy-pipeline --pipeline resource-marathon-deploy
