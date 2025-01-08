#! /bin/bash

YELLOW='\033[1;33m'
NC='\033[0m' # No Color

#m Edit these two environment variables as you create more environment setup scripts.
export ENVIRONMENT_PROMPT="HTTP-LOGGER"
export KUBECONFIG=$HOME/.kube/http-logger

printf "Setting up isolated Kubernetes environment in: ${YELLOW}$KUBECONFIG${NC}\n"
printf "NOTE: The command format to run this script is: \". $BASH_SOURCE\"\n"

#export PS1="\[$(tput setaf 2)\]$ENVIRONMENT_PROMPT: \[$(tput setaf 7)\]\[$(tput setaf 6)\]\W\\$ \[$(tput setaf 7)\]\[$(tput sgr0)\]"
export PS1="\[$(tput setaf 2)\]\$ENVIRONMENT_PROMPT: \[$(tput setaf 3)\][\$(kubectl config view --minify --output 'jsonpath={..namespace}' 2>/dev/null || echo 'default')]\[$(tput setaf 7)\]:\$(parse_git_branch)\[$(tput setaf 6)\]\W\\$ \[$(tput setaf 7)\]\[$(tput sgr0)\]"

