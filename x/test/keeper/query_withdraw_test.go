package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "zyx/testutil/keeper"
	"zyx/testutil/nullify"
	"zyx/x/test/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestWithdrawQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNWithdraw(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetWithdrawRequest
		response *types.QueryGetWithdrawResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetWithdrawRequest{
				TxHash: msgs[0].TxHash,
			},
			response: &types.QueryGetWithdrawResponse{Withdraw: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetWithdrawRequest{
				TxHash: msgs[1].TxHash,
			},
			response: &types.QueryGetWithdrawResponse{Withdraw: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetWithdrawRequest{
				TxHash: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Withdraw(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestWithdrawQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.TestKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNWithdraw(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllWithdrawRequest {
		return &types.QueryAllWithdrawRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.WithdrawAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Withdraw), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Withdraw),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.WithdrawAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Withdraw), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Withdraw),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.WithdrawAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Withdraw),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.WithdrawAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
