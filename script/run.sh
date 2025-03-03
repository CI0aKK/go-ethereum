#!/bin/sh
cd ..
make geth
nohup ./build/bin/geth --syncmode "snap" --http --http.addr "0.0.0.0" --http.port 8545 --http.api "eth,net,web3,allen" --bootnodes=enode://2b252ab6a1d0f971d9722cb839a42cb81db019ba44c08754628ab4a823487071b5695317c8ccd085219c3a03af063495b2f1da8d18218da2d6a82981b45e6ffc@65.108.70.101:30303,enode://4aeb4ab6c14b23e2c4cfdce879c04b0748a20d8e9b59e25ded2a08143e265c6c25936e74cbc8e641e3312ca288673d91f2f93f8e277de3cfa444ecdaaf982052@157.90.35.166:30303 --authrpc.jwtsecret=~/ethereum/jwt.hex >> /root/go/github/allen/go-ethereum/log/geth_log.txt 2>&1 &