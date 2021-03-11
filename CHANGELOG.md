# [v0.3.0](https://github.com/nervosnetwork/ckb-sdk-go/compare/v0.2.1...v0.3.0) (2021-03-11)


### Bug Fixes

* wrong change target ([0706e68](https://github.com/nervosnetwork/ckb-sdk-go/commit/0706e68))
* wrong witness ([184145b](https://github.com/nervosnetwork/ckb-sdk-go/commit/184145b))


### Features

* add cheque cell args generation function ([a4d6f63](https://github.com/nervosnetwork/ckb-sdk-go/commit/a4d6f63))
* add claim cheque payment ([5579e55](https://github.com/nervosnetwork/ckb-sdk-go/commit/5579e55))
* add claim cheques unsigned tx builder ([065848d](https://github.com/nervosnetwork/ckb-sdk-go/commit/065848d))
* add GetBlockEconomicState ([0edae6a](https://github.com/nervosnetwork/ckb-sdk-go/commit/0edae6a))
* add IsChequeCell function ([8af982e](https://github.com/nervosnetwork/ckb-sdk-go/commit/8af982e))
* add IssuingChequeUnsignedTxBuilder ([76c964d](https://github.com/nervosnetwork/ckb-sdk-go/commit/76c964d))
* add OutputsCapacity method to transaction ([5702cd0](https://github.com/nervosnetwork/ckb-sdk-go/commit/5702cd0))
* add ReceiverInfo ([d99b073](https://github.com/nervosnetwork/ckb-sdk-go/commit/d99b073))
* add searchLimit const ([03398d7](https://github.com/nervosnetwork/ckb-sdk-go/commit/03398d7))
* add sign issue cheque tx method ([0ccf34c](https://github.com/nervosnetwork/ckb-sdk-go/commit/0ccf34c))
* add since generator ([69a83c4](https://github.com/nervosnetwork/ckb-sdk-go/commit/69a83c4))
* add sudt payment ([4938ff8](https://github.com/nervosnetwork/ckb-sdk-go/commit/4938ff8))
* add udt live cell collector ([73b7898](https://github.com/nervosnetwork/ckb-sdk-go/commit/73b7898))
* add UnsignedTxBuilder interface ([1c5588e](https://github.com/nervosnetwork/ckb-sdk-go/commit/1c5588e))
* add ValidateChequeAddress function ([a8621ff](https://github.com/nervosnetwork/ckb-sdk-go/commit/a8621ff))
* add withdraw cheque payment ([611b7d3](https://github.com/nervosnetwork/ckb-sdk-go/commit/611b7d3))
* hash witnesses which do not in any input group ([e3a293f](https://github.com/nervosnetwork/ckb-sdk-go/commit/e3a293f))
* implement generate unsigned issuing cheque tx ([d2d09d7](https://github.com/nervosnetwork/ckb-sdk-go/commit/d2d09d7))
* more set function on SystemScript ([129dec1](https://github.com/nervosnetwork/ckb-sdk-go/commit/129dec1))
* remove get_cells_by_lock_hash RPC and RPCs under indexer module ([fc920d4](https://github.com/nervosnetwork/ckb-sdk-go/commit/fc920d4))
* support custom SystemScripts ([f851553](https://github.com/nervosnetwork/ckb-sdk-go/commit/f851553))
* support filter on ckb-indexer searchKey ([aebf2f7](https://github.com/nervosnetwork/ckb-sdk-go/commit/aebf2f7))


### BREAKING CHANGES

* Remove RPCs under indexer module
* need send a SystemScripts to `GenerateTx` method manually



# [v0.2.1](https://github.com/nervosnetwork/ckb-sdk-go/compare/v0.2.0...v0.2.1) (2020-11-25)

### Features

* [#13](https://github.com/nervosnetwork/ckb-sdk-go/pull/13): expose GenerateFullPayloadAddress function 

# [v0.2.0](https://github.com/nervosnetwork/ckb-bitpie-sdk/compare/v0.1.0...v0.2.0) (2020-11-25)


### Bug Fixes

* [#5](https://github.com/nervosnetwork/ckb-sdk-go/pull/5): fix nil pointer dereference on toCellWithStatus function
* [#7](https://github.com/nervosnetwork/ckb-sdk-go/pull/7): fix tx fee calculation bug


### Features

* [#6](https://github.com/nervosnetwork/ckb-sdk-go/pull/6): support ckb indexer
* [#8](https://github.com/nervosnetwork/ckb-sdk-go/pull/8): support ckb indexer on payment
* [#9](https://github.com/nervosnetwork/ckb-sdk-go/pull/9): add OccupiedCapacity function
* [#10](https://github.com/nervosnetwork/ckb-sdk-go/pull/10): support generate and parse short acp address
