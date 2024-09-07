#!/bin/bash

set -e

main() {
    golines_output=$(golines --list-files ./)

    [[ -n "$golines_output" ]] &&
        {
            printf "%s $(tput setaf 1)%s\n$(tput sgr0)" "❌" "Go code is not formatted:"
            echo "$golines_output"
            exit 1
        } ||
        printf "%s $(tput setaf 2)%s\n$(tput sgr0)" "✅" "Formatting passed!"

    testifylint_output=$(testifylint ./...)

    [[ -n "$testifylint_output" ]] &&
        {
            printf "%s $(tput setaf 1)%s\n$(tput sgr0)" "❌" "Tests have linting issues:"
            echo "$testifylint_output"
            exit 1
        } ||
        printf "%s $(tput setaf 2)%s\n$(tput sgr0)" "✅" "Test linting passed!"

    # TODO: This fails in CI.
    # mockery

    # if [[ -n "$(git status -s -uno)" ]]; then
    #     printf "%s $(tput setaf 1)%s\n$(tput sgr0)" "❌" "mockery generated output does not match commit."
    #     printf "%s $(tput setaf 3)%s $(tput setaf 4)$(tput bold)%s$(tput sgr0) $(tput setaf 3)%s\n$(tput sgr0)" "❌" "Did you forget to run" "mockery" "?"
    #     exit 1
    # fi

    # go generate ./...

    # if [[ -n "$(git status -s -uno)" ]]; then
    #     printf "%s $(tput setaf 1)%s\n$(tput sgr0)" "❌" "Go generate output does not match commit."
    #     printf "%s $(tput setaf 3)%s $(tput setaf 4)$(tput bold)%s$(tput sgr0) $(tput setaf 3)%s\n$(tput sgr0)" "❌" "Did you forget to run" "go generate ./..." "?"
    #     exit 1
    # fi
}

main
