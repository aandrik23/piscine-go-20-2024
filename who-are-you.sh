#!/bin/bash

# Fetch data from the JSON API and extract the name of the superhero with ID 70
curl -s https://platform.zone01.gr/assets/superhero/all.json | jq -r '.[] | select(.id == 70) | .name'

