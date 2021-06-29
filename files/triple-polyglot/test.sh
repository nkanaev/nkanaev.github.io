#!/bin/sh

run() {
    #sh -c "$1 code.txt "
    sh -c "$1 code.txt" | cmp - output.txt && echo "$1: pass" || echo "$1: fail"
}

run "perl"
run "python"
run "python3"
run "ruby"
