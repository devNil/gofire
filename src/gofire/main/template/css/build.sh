#!/bin/bash
for filename in *.less 
do
	lessc "$filename" "${filename%.*}.css"
done