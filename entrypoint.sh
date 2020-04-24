#!/bin/sh

/root/main;
git diff
git config --global user.email "readme-bot@example.com"
git config --global user.name "README-bot"
echo "git add README.md && git commit -m \"Updated README\""
git add README.md && git commit -m "Updated README"
echo "git push"
git push
