#!/usr/bin/env bash

mockgen_cmd="mockgen"
$mockgen_cmd -source=x/tradebin/types/expected_keepers.go -package testutil -destination x/tradebin/testutil/expected_keepers_mocks.go
$mockgen_cmd -source=x/rewards/types/expected_keepers.go -package testutil -destination x/rewards/testutil/expected_keepers_mocks.go
