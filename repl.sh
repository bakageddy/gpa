#!/usr/bin/env sh

ls ./templates/*.templ ./services/*.go ./static/* ./types/* ./utils/* | entr -rcs 'go run main.go'
