#!/bin/bash
message="You need to actually run maintenance in order to have it run."
git add --all
git commit -m "$message"
git push