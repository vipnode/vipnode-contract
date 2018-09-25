TRUFFLEBIN := ./node_modules/.bin/truffle

build: deps contracts

test: deps contracts
	$(TRUFFLEBIN) test

run: deps contracts
	npm run dev

deploy-contract:
	$(TRUFFLEBIN) migrate -f 2 --network rinkeby --reset

deps: node_modules/

contracts: build/contracts/*

build/contracts/%.json: contracts/%.sol
	$(TRUFFLEBIN) compile

node_modules/: package.json
	npm install
	touch node_modules/

gopkg:
	# Grab a solc binary from https://github.com/ethereum/solidity/releases
	abigen --sol contracts/VipnodePool.sol --pkg vipnodepool --out go/vipnodepool/VipnodePool.go
