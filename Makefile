all: generate test bench

generate:
	go generate

test:
	go test -v .

bench:
	go test -v -bench=.

emotes:
	curl -s 'https://www.twitchmetrics.net/emotes' | grep samp | sed 's;.*samp>\(.*\)</samp.*;\1;' > newemotes.txt
	curl -s 'https://api.betterttv.net/3/cached/emotes/global' | jq '. | map(.code)' | head -n -1 | tail -n +2 | tr -d '"' | tr -d ',' | tr -d ' ' >> newemotes.txt
	curl -s 'https://api.betterttv.net/3/emotes/shared/top?offset=0&limit=100' | jq '. | map(.emote.code)' | head -n -1 | tail -n +2 | tr -d '"' | tr -d ',' | tr -d ' ' >> newemotes.txt
	curl -s 'https://api.betterttv.net/3/emotes/shared/top?offset=100&limit=100' | jq '. | map(.emote.code)' | head -n -1 | tail -n +2 | tr -d '"' | tr -d ',' | tr -d ' ' >> newemotes.txt
	curl -s 'https://gist.githubusercontent.com/oliveratgithub/0bf11a9aff0d6da7b46f1490f86a71eb/raw/ac8dde8a374066bcbcf44a8296fc0522c7392244/emojis.json' | jq '.emojis | map(.shortname)' | head -n -1 | tail -n +2 | tr -d '"' | tr -d ',' | tr -d ' ' | grep ':' >> newemotes.txt
	sort newemotes.txt | uniq > emotes.txt
	rm newemotes.txt
