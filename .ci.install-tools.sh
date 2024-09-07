#!/bin/bash

main() {
    local tools
    tools=$(go list -e -f '{{join .Imports " "}}' ./internal/tools/tools.go)
    IFS=' ' read -r -a tools <<<"$tools"
    unset IFS

    printf "%s %s\n" "⚙️ " "Installing dependencies..."
    
    for tool in "${tools[@]}"; do
        go install "$tool@latest"
    done

    printf "%s $(tput setaf 2)%s\n$(tput sgr0)" "✅" "Installed dependencies!"
}

main
