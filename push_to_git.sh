#!/bin/bash
message="Running the game with args [git message...] runs generators and pushes changes to Git with [message]."
git add --all
git commit -m "$message"
git push