package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v2 "github.com/desmos-labs/desmos/v4/x/relationships/legacy/v2"

	profilesv4 "github.com/desmos-labs/desmos/v4/x/profiles/legacy/v4"

	v1 "github.com/desmos-labs/desmos/v4/x/relationships/legacy/v1"
)

// DONTCOVER

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
	pk     profilesv4.Keeper
}

// NewMigrator returns a new Migrator
func NewMigrator(keeper Keeper, pk profilesv4.Keeper) Migrator {
	return Migrator{
		keeper: keeper,
		pk:     pk,
	}
}

// Migrate1To2 migrates from version 1 to 2.
func (m Migrator) Migrate1To2(ctx sdk.Context) error {
	return v1.MigrateStore(ctx, m.pk, m.keeper.storeKey, m.keeper.cdc)
}

// Migrate1To2 migrates from version 2 to 3.
func (m Migrator) Migrate2To3(ctx sdk.Context) error {
	return v2.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}
