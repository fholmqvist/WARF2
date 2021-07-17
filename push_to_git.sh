#!/bin/bash
message="PlantFarm. Moved some items to globals due to import cycle not allowed."
git add --all
git commit -m "$message"
git push