#!/bin/bash
#build all less files
for filename in *.less 
do
	lessc "$filename" "${filename%.*}.css"
done
