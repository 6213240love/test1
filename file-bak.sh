#!/bin/bash
datetime=$(date +%Y-%m-%D)
git add .
git commit -m "update $datetime"
git push origin master

