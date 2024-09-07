#!/bin/bash

main() {
    cd "$(git rev-parse --show-toplevel)/internal/tools" ||
        printf "$(tput setaf 1)%s\n$(tput sgr0)" "Could not change directory"

    local tools
    tools=$(go list -e -f '{{join .Imports " "}}' ./tools.go)
    IFS=' ' read -r -a tools <<<"$tools"
    unset IFS

    printf "%s %s\n" "⚙️ " "Installing dependencies..."
    go mod tidy

    for tool in "${tools[@]}"; do
        go install "$tool@latest"
    done

    printf "%s $(tput setaf 2)%s\n$(tput sgr0)" "✅" "Installed dependencies!"
}

main
