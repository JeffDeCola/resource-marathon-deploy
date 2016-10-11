#!/bin/bash
# resource-marathon-deploy readme-github-pages.sh

set -e -x

# The code is located in /resource-marathon-deploy
echo "pwd is: " $PWD
echo "List whats in the current directory"
ls -lat 

# Note: resource-marathon-deploy-updated already created becasue of yml file
git clone resource-marathon-deploy resource-marathon-deploy-updated

cd resource-marathon-deploy-updated
ls -lat 

# FOR GITHUB WEBPAGES
# BASICALLY COPY README.md to /docs/_includes/README.md
# Remove everything before the second hedading.
sed '0,/GitHub Webpage/d' README.md > docs/_includes/README.md
# update the image links (remove docs/)
sed -i 's#IMAGE](docs/#IMAGE](#g' docs/_includes/README.md

#ADD AND COMMIT
git config --global user.email "jeff@keeperlabs.com"
git config --global user.name "Jeff DeCola (concourse)"

git status
git add .
git commit -m "cp README.md to docs/_includes/README.md"
git status
#git config --list
