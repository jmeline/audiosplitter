#!/bin/bash

# Path to the directory containing m4a files
directory="ParadiseTranceTOP50VOCALTRANCE2014"

# Loop through all .m4a files in the directory
for file in "$directory"/*.m4a; do
  # Check if the file exists
  if [ -f "$file" ]; then
    # Get the base name of the file without the extension
    base_name=$(basename "$file" .m4a)

    # Output file name with .mp3 extension
    output_file="$directory/$base_name.mp3"

    # Convert m4a to mp3 using ffmpeg
    ffmpeg -i "$file" -acodec libmp3lame -ab 192k "$output_file"

    echo "Converted $file to $output_file"
  fi
done
