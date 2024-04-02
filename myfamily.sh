#!/bin/bash

# Fetch the superhero data from the JSON API and extract the family of the superhero with the given ID
family=$(curl -s https://platform.zone01.gr/assets/superhero/all.json | jq -r --arg HERO_ID "$HERO_ID" '.[] | select(.id == ($HERO_ID|tonumber)) | .relatives')

# Remove quotes from the family string
family=$(echo "$family" | tr -d '"')

# Display the family
echo "$family"
