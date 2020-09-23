package nv3

import (
	"context"

	cid "github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/specs-actors/actors/builtin/reward"
)

type rewardMigrator struct {
}

func (m *rewardMigrator) MigrateState(ctx context.Context, store cbor.IpldStore, head cid.Cid) (cid.Cid, error) {
	var st reward.State
	if err := store.Get(ctx, head, &st); err != nil {
		return cid.Undef, err
	}

	// The baseline function initial value and growth rate are changed.
	// As an approximation to working out what the baseline value would be at the migration epoch
	// had the baseline parameters been this way all along, this just sets the immediately value to the
	// new (higher) initial value, essentially restarting its calculation.
	// The baseline and realized cumsums, and effective network time, are not changed.
	// This boils down to a step change in the baseline function, as if it had been defined piecewise.
	// This will be a bit annoying for external analytical calculations of the baseline function.
	newBaselineInitialValue := reward.BaselineInitialValueV3

	if st.ThisEpochBaselinePower.GreaterThan(newBaselineInitialValue) {
		return cid.Undef, xerrors.Errorf("unexpected baseline power %v higher than new initial value %v",
			st.ThisEpochBaselinePower, newBaselineInitialValue)
	}

	st.ThisEpochBaselinePower = newBaselineInitialValue

	newHead, err := store.Put(ctx, &st)
	return newHead, err
}
