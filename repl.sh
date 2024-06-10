#!/usr/bin/env sh

ls ./templates/*.templ ./services/*.go ./static/* ./types/* ./utils/* | entr -rcs 'DEBUG_APP=1 go run main.go'
