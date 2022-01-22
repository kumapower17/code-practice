#! /bin/bash

# grep -E is equivalent to egrep
grep -E -e '^(\([0-9]{3}\) |[0-9]{3}-)[0-9]{3}-[0-9]{4}$' file.txt
