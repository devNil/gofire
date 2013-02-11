#!/bin/bash
#build all less files
if [ $1 ]
then
	if [ $1 == -b ]
		then
		for filename in *.less 
		do
			lessc "$filename" "${filename%.*}.css"
		done
	fi
	if [ $1 == -c ]
		then
		for filename in *.css 
		do
			rm "$filename"
		done
	fi
else 
	echo "./build -b or -c"
fi
