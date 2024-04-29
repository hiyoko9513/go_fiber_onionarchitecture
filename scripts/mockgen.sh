#!/bin/bash

declare -A directories=(
  ["internal/application/usecase"]="internal/application/usecase/mocks"
  ["internal/domain/services"]="internal/domain/services/mocks"
  ["internal/interactor"]="internal/interactor/mocks"
  ["internal/presentation/http/app/handler"]="internal/presentation/http/app/handler/mocks"
)

for input_directory in "${!directories[@]}"; do
  output_directory="${directories[$input_directory]}"

  # Ensure that the output directory exists
  mkdir -p $output_directory

  # Process each .go source file separately
  for file in $(find "$input_directory" -name '*.go'); do
    # Get filename without extension
    name=$(basename "$file" .go)
    # Generate the mock file
    mockgen -source="$file" -destination="$output_directory/${name}.go"
  done
done
