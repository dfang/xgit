#!/bin/sh

# https://www.conventionalcommits.org/en/v1.0.0/
# https://github.com/conventionalcommit/commitlint#available-rules

# alternative
# https://github.com/go-semantic-release/semantic-release#how-does-it-work
# https://github.com/hazcod/semantic-commit-hook

if ! type commitlint >/dev/null 2>/dev/null; then
	echo ""
    echo "commitlint could not be found"
    echo "try again after installing commitlint or add commitlint to PATH"
	echo ""
    exit 2;
fi

commitlint lint --message $1
