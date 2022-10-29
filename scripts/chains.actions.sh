#!/usr/bin/env bash

function chain-gen() {
  cd tools
  go build
  cd -
  rg -n 'wg-chain' /home/cong/sm/project/openresty-wg | ./tools/chain
}
