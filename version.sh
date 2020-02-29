#!/bin/bash

sed -i s/version\ =\ .*$/version\ =\ \"$1\"/g version.go
sed -i s/\"Version\":.*$/\"Version\":\"$1\"/g version.json
