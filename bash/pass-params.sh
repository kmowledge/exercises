#! /usr/bin/env bash

#function definition
pass_args() {
echo "You provided $# facts about yourself!"
echo "Your name is $1."
echo "Your age is $2."
echo "${1}123 is a weak password"
}

#function call
pass_args "$1" $2
#If you pass the name argument as 'name surname' it will be caught as 1 param only via:  "$1"
