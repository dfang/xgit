#!/bin/sh

# https://www.conventionalcommits.org/en/v1.0.0/
# https://github.com/conventionalcommit/commitlint#available-rules

if ! type commitlint >/dev/null 2>/dev/null; then
	echo ""
    echo "commitlint could not be found"
    echo "try again after installing commitlint or add commitlint to PATH"
	echo ""
    exit 2;
fi

commitlint lint --message $1
