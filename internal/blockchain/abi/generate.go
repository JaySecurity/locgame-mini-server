package abi

//go:generate abigen --abi LOCGamePlayNFT.abi.json --pkg contracts --type LOCGamePlayNFT --out ../contracts/LOCGamePlayNFT.go
//go:generate abigen --abi ERC20.abi.json --pkg contracts --type ERC20 --out ../contracts/ERC20.go
//go:generate abigen --abi LOCGBridged.abi.json --pkg contracts --type ERC20 --out ../contracts/LOCGBridged.go