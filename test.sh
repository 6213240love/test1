#!/bin/bash
echo "nihao"
eho "nihao" || { echo "command failed"; exit 1; }
echo "success"
