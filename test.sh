#!/usr/bin/env bash

go test --coverprofile=coverage.out ./...

covered=0
total=0
while IFS='' read -r line || [[ -n "$line" ]]; do
    IFS=' ' read -r -a array <<< "$line"
    total=$(($total+${array[1]}))
    if [ "${array[2]}" = "1" ]; then
        covered=$(($covered+${array[1]}))
    fi 
done < "coverage.out"

echo ""
echo "global coverage:" $(awk "BEGIN { pc=100*${covered}/${total}; i=int(pc); print (pc-i<0.5)?i:i+1 }")"% of statements"

rm coverage.out