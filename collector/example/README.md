# Example
CKB Go SDK examples

## Introduction

The [example.go](./example.go) provides some examples shows how to use this sdk to interact with CKB's RPC, transaction, and Nervos DAO(deposit and withdraw).

## Transaction Examples

These examples show how to transfer CKB, in different situations.

- `SendCkbExample`: sign and send CKB from single-sig address.
- `SendCkbByLightClientExample`: similiar to `SendCkbExample` but interact with [`ckb-light-client`](https://github.com/nervosnetwork/ckb-light-client)
- `SendCkbFromMultisigAddressExample`: sign and send CKB from multi-sig address with multiple private keys.
- `SendChainedTransactionExample`: sign and send offchain transaction

## SUDT Examples

[SUDT](https://github.com/nervosnetwork/rfcs/blob/master/rfcs/0025-simple-udt/0025-simple-udt.md)  (Simple User Defined Tokens) is a token specification on CKB. 
You can think SUDT is an analog of ERC20 on Ethereum.

Anyone can issue his own SUDT, or transfer a specific kind of SUDT if he has enough SUDT amount. SUDT smart contract should be used in type script in these transactions.

- `IssueSudtExample`: shows how to issue SUDT
- `SendSudtExample`: shows how to send issued SUDT

## Nervos DAO Examples

Nervos DAO is a smart contract, with which users can interact the same way as any smart contracts on CKB. Nervos DAO has deposit and withdraw (phase1 and phase2).

- Deposit: Users can send a transaction to deposit CKB into Nervos DAO at any time. CKB includes a special Nervos DAO type script in the genesis block.
- Withdraw: Users can send a transaction to withdraw deposited CKB from Nervos DAO at any time (but a locking period will be applied to determine when exactly the tokens can be withdrawn). 

Check [Nervos DAO RFC](https://github.com/nervosnetwork/rfcs/blob/master/rfcs/0023-dao-deposit-withdraw/0023-dao-deposit-withdraw.md) for more details

- `DepositDaoExample`: shows how to deposit DAO
- `WithdrawDaoExample`: shows how to withdraw deposited CKB from Nervos DAO

## OmniLock Examples

Omnilock is a new lock script designed for interoperability. Check https://blog.cryptape.com/omnilock-a-universal-lock-that-powers-interoperability-1 for more details

- `SendCkbOmnilockExample`: similar to `SendCkbExample`, but use Omnilock to sign. 
- `SendCkbMultisigOmnilockExample`: similar to `SendCkbFromMultisigAddressExample`, but use Omnilock to sign.
