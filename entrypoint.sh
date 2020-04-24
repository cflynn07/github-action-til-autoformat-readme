#!/bin/sh

/root/main;
git diff
git config --global user.email "readme-bot@example.com"
git config --global user.name "README-bot"
git add ./README.md
git diff --quiet || (git add README.md && git commit -m "Updated README")
git push
