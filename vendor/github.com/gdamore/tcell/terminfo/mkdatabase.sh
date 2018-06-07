#!/bin/bash

# Copyright 2017 The TCell Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use file except in compliance with the License.
# You may obtain a copy of the license at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

#
# When called with no arguments, this shell script builds the Go database,
# which is somewhat minimal for size reasons (it only contains the most
# commonly used entries), and then builds the complete JSON database.
#
# To limit the action to only building one or more terminals, specify them
# on the command line:
#
#  ./mkdatabase xterm
#
# The script will also find and update or add any terminal "aliases".
# It does not remove any old entries.
#
# To add to the set of terminals that we compile into the Go database,
# add their names to the models.txt file.
#

# This script is not very efficient, but there isn't really a better way
# without writing code to decode the terminfo binary format directly.
# Its not worth worrying about.

# This script also requires bash, although ksh93 should work as well, because
# we use arrays, which are not specified in POSIX.

export LANG=C
export LC_CTYPE=C

progress()
{
	typeset -i num=$1
	typeset -i tot=$2
	typeset -i x
	typeset back
	typeset s

	if (( tot < 1 ))
	then
		s=$(printf "[ %d ]" $num)
		back="\b\b\b\b\b"
		x=$num
		while (( x >= 10 ))
		do
			back="${back}\b"
			x=$(( x / 10 ))
		done
		
	else
		x=$(( num * 100 / tot ))
		s=$(printf "<%3d%%>" $x)
		back="\b\b\b\b\b\b"
	fi
	printf "%s${back}" "$s"
}

ord()
{
	printf "%02x" "'$1'"
}

goterms=( $(cat models.txt) )
args=( $* )
if (( ${#args[@]} == 0 ))
then
	args=( $(toe -a | cut -f1) )
fi

printf "Scanning terminal definitions: "
i=0
aliases=()
models=()
for term in ${args[@]}
do
	case "${term}" in
	*-truecolor)
		line="${term}|24-bit color"
		;;
	*)
		line=$(infocmp $term | head -2 | tail -1)
		if [[ -z "$line" ]]
		then
			echo "Cannot find terminfo for $term"
			exit 1
		fi
		# take off the trailing comma
		line=${line%,}
	esac

	# grab primary name
	term=${line%%|*}
	all+=( ${term} )

	# should this be in our go terminals?
	for model in ${goterms[@]}
	do
		if [[ "${model}" == "${term}" ]]
		then
			models+=( ${term} )
		fi
	done

	# chop off primary name
	line=${line#${term}}
	line=${line#|}
	# chop off description
	line=${line%|*}
	while [[ "$line" != "" ]]
	do
		a=${line%%|*}
		aliases+=( ${a}=${term} )
		line=${line#$a}
		line=${line#|}
	done
	i=$(( i + 1 ))
	progress $i ${#args[@]}
done
echo
# make sure we have mkinfo
printf "Building mkinfo: "
go build mkinfo.go
echo "done."

# Build all the go database files for the "interesting" terminals".
printf "Building Go database: "
i=0
for model in ${models[@]}
do
	safe=$(echo $model | tr - _)
	file=term_${safe}.go
	./mkinfo -go $file $model
	go fmt ${file} >/dev/null
	i=$(( i + 1 ))
   	progress $i ${#models[@]}
done
echo

printf "Building JSON database: "

# The JSON files are located for each terminal in a file with the
# terminal name, in the following fashion "database/x/xterm.json

i=0
for model in ${all[@]}
do
	letter=$(ord ${model:0:1})
	dir=database/${letter}
	file=${dir}/${model}.gz
	mkdir -p ${dir}
	./mkinfo -nofatal -quiet -gzip -json ${file} ${model}
	i=$(( i + 1 ))
	progress $i ${#all[@]}
done
echo

printf "Building JSON aliases: "
i=0
for model in ${aliases[@]}
do
	canon=${model#*=}
	model=${model%=*}
	letter=$(ord ${model:0:1})
	cletter=$(ord ${canon:0:1})
	dir=database/${letter}
	file=${dir}/${model}
	if [[  -f database/${cletter}/${canon}.gz ]]
	then
		[[ -d ${dir} ]] || mkdir -p ${dir}
		# Generally speaking the aliases are better uncompressed
		./mkinfo -nofatal -quiet -json ${file} ${model}
	fi
	i=$(( i + 1 ))
	progress $i ${#aliases[@]}
done
echo
