#!/bin/bash

# Check if HERO_ID environment variable is set
if [ -z "$HERO_ID" ]; then
  echo "HERO_ID environment variable is not set."
  exit 1
fi

# Fetch data from the URL and parse JSON using jq
family=$(curl -s https://platform.zone01.gr/assets/superhero/all.json | jq -r ".[] | select(.id == \"$HERO_ID\") | .relatives")

# Remove quotes from the relatives field
family=$(echo "$family" | sed 's/"//g')

# Display the family information
echo "Family of Hero ID $HERO_ID:"
echo "$family"
