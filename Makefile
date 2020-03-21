all: generate test bench

generate:
	go generate

test:
	go test -v .

bench:
	go test -v -bench=.

emotes:
	curl -s 'https://www.twitchmetrics.net/emotes' -o - | grep samp | sed 's;.*samp>\(.*\)</samp.*;\1;' > newemotes.txt
	curl -s 'https://api.betterttv.net/3/cached/emotes/global' | jq '. | map(.code)' | head -n -1 | tail -n +2 | tr -d '"' | tr -d ',' | tr -d ' ' >> newemotes.txt
	curl -s 'https://api.betterttv.net/3/emotes/shared/top?offset=0&limit=100' | jq '. | map(.emote.code)' | head -n -1 | tail -n +2 | tr -d '"' | tr -d ',' | tr -d ' ' >> newemotes.txt
	curl -s 'https://api.betterttv.net/3/emotes/shared/top?offset=100&limit=100' | jq '. | map(.emote.code)' | head -n -1 | tail -n +2 | tr -d '"' | tr -d ',' | tr -d ' ' >> newemotes.txt
	sort newemotes.txt | uniq > emotes.txt
	rm newemotes.txt
