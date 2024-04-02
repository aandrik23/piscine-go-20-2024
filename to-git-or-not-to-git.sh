#!/bin/bash

# Fetch data from the JSON API and extract name, power, and gender of superhero with id 170
curl -s https://platform.zone01.gr/assets/superhero/all.json | jq -r '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender'
