#!/bin/bash
message="Inlined Job.Priority, fixed bug where JobService sorted before checking for jobs."
git add --all
git commit -m "$message"
git push