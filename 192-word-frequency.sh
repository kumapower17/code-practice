#! /bin/bash

# https://leetcode-cn.com/problems/word-frequency/

# Read from the file words.txt and output the word frequency list to stdout.
tr ' ' '\n' <words.txt | sed -e /^$/d | sort | uniq -c | sort -k1 -n -r | awk '{ print $2, $1 }'
