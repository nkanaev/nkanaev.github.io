default: test

ouroboros.c: codegen.py
	python3 codegen.py > ouroboros.c

ouroboros.go: ouroboros.c
	cc ouroboros.c -o /tmp/ouroboros && /tmp/ouroboros > ouroboros.go

test: ouroboros.go ouroboros.c
	go run ouroboros.go | diff - ouroboros.c && echo "pass" || echo "fail"
