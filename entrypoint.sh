#!/bin/sh

git config --global --add safe.directory /github/workspace
git config --global user.email "readme-bot@example.com"
git config --global user.name "README-bot"

/root/main;

git diff
echo "git add README.md && git commit -m \"Updated README\""
git add README.md && git commit -m "Updated README"
echo "git push"
git push
