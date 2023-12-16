#!/bin/bash

commit_msg_file=$1
commit_msg=$(cat "$commit_msg_file")

# Define the commit message pattern
commit_msg_pattern="^(feat|fix|docs|style|refactor|test|chore|ci)(\(.+\))?: .+"

# Check if the commit message matches the pattern
if ! [[ $commit_msg =~ $commit_msg_pattern ]]; then
    echo "Invalid commit message format. Please follow the pattern: <type>(<scope>): <message>"
    
    # Suggest a valid commit message format
    echo "Suggested format: feat(scope): Add new feature"
    
    # Available prefixes
    echo "Available prefixes: feat, fix, docs, style, refactor, test, chore"
    
    exit 1
fi
