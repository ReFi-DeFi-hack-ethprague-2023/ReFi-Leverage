package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/umee-network/umee/v2/x/leverage/types"
)

// Flag constants
const (
	FlagDenom = "denom"
)

// GetQueryCmd returns the CLI query commands for the x/leverage module.
func GetQueryCmd(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdQueryAllRegisteredTokens(),
		GetCmdQueryParams(),
		GetCmdQueryBorrowed(),
		GetCmdQueryBorrowedValue(),
		GetCmdQuerySupplied(),
		GetCmdQuerySuppliedValue(),
		GetCmdQueryReserveAmount(),
		GetCmdQueryCollateral(),
		GetCmdQueryCollateralValue(),
		GetCmdQueryExchangeRate(),
		GetCmdQuerySupplyAPY(),
		GetCmdQueryBorrowAPY(),
		GetCmdQueryMarketSize(),
		GetCmdQueryTokenMarketSize(),
		GetCmdQueryBorrowLimit(),
		GetCmdQueryLiquidationThreshold(),
		GetCmdQueryLiquidationTargets(),
		GetCmdQueryMarketSummary(),
		GetCmdQueryTotalCollateral(),
		GetCmdQueryTotalBorrowed(),
	)

	return cmd
}

// GetCmdQueryAllRegisteredTokens creates a Cobra command to query for all
// the registered tokens in the x/leverage module.
func GetCmdQueryAllRegisteredTokens() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "registered-tokens",
		Args:  cobra.NoArgs,
		Short: "Query for all the current registered tokens",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			resp, err := queryClient.RegisteredTokens(cmd.Context(), &types.QueryRegisteredTokens{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryParams creates a Cobra command to query for the x/leverage
// module parameters.
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the x/leverage module parameters",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			resp, err := queryClient.Params(cmd.Context(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryBorrowed creates a Cobra command to query for the amount of
// total borrowed tokens for a given address.
func GetCmdQueryBorrowed() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "borrowed [addr]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for the total amount of borrowed tokens for an address",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryBorrowedRequest{
				Address: args[0],
			}
			if d, err := cmd.Flags().GetString(FlagDenom); len(d) > 0 && err == nil {
				req.Denom = d
			}

			resp, err := queryClient.Borrowed(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	cmd.Flags().String(FlagDenom, "", "Query for a specific denomination")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryBorrowedValue creates a Cobra command to query for the USD
// value of total borrowed tokens for a given address.
func GetCmdQueryBorrowedValue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "borrowed-value [addr]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for the total USD value of borrowed tokens for an address",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryBorrowedValueRequest{
				Address: args[0],
			}
			if d, err := cmd.Flags().GetString(FlagDenom); len(d) > 0 && err == nil {
				req.Denom = d
			}

			resp, err := queryClient.BorrowedValue(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	cmd.Flags().String(FlagDenom, "", "Query for value of only a specific denomination")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQuerySupplied creates a Cobra command to query for the amount of
// tokens supplied by a given address.
func GetCmdQuerySupplied() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "supplied [addr]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for the total amount of tokens supplied by an address",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QuerySuppliedRequest{
				Address: args[0],
			}
			if d, err := cmd.Flags().GetString(FlagDenom); len(d) > 0 && err == nil {
				req.Denom = d
			}

			resp, err := queryClient.Supplied(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	cmd.Flags().String(FlagDenom, "", "Query for a specific denomination")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQuerySuppliedValue creates a Cobra command to query for the USD value of
// total tokens supplied by a given address.
func GetCmdQuerySuppliedValue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "supplied-value [addr]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for the USD value of tokens supplied by an address",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QuerySuppliedValueRequest{
				Address: args[0],
			}
			if d, err := cmd.Flags().GetString(FlagDenom); len(d) > 0 && err == nil {
				req.Denom = d
			}

			resp, err := queryClient.SuppliedValue(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	cmd.Flags().String(FlagDenom, "", "Query for value of only a specific denomination")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryReserveAmount creates a Cobra command to query for the
// reserved amount of a specific token.
func GetCmdQueryReserveAmount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reserved [denom]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for the amount reserved of a specified denomination",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryReserveAmountRequest{
				Denom: args[0],
			}

			resp, err := queryClient.ReserveAmount(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryCollateral creates a Cobra command to query for the amount of
// total collateral tokens for a given address.
func GetCmdQueryCollateral() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collateral [addr]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for the total amount of collateral tokens for an address",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryCollateralRequest{
				Address: args[0],
			}
			if d, err := cmd.Flags().GetString(FlagDenom); len(d) > 0 && err == nil {
				req.Denom = d
			}

			resp, err := queryClient.Collateral(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	cmd.Flags().String(FlagDenom, "", "Query for a specific denomination")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryCollateralValue creates a Cobra command to query for the USD
// value of total collateral tokens for a given address.
func GetCmdQueryCollateralValue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collateral-value [addr]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for the total USD value of collateral tokens for an address",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryCollateralValueRequest{
				Address: args[0],
			}
			if d, err := cmd.Flags().GetString(FlagDenom); len(d) > 0 && err == nil {
				req.Denom = d
			}

			resp, err := queryClient.CollateralValue(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	cmd.Flags().String(FlagDenom, "", "Query for value of only a specific denomination")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryExchangeRate creates a Cobra command to query for the
// exchange rate of a specific uToken.
func GetCmdQueryExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exchange-rate [denom]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for the exchange rate of a specified denomination",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryExchangeRateRequest{
				Denom: args[0],
			}

			resp, err := queryClient.ExchangeRate(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryAvailableBorrow creates a Cobra command to query for the
// available amount to borrow of a specific denom.
func GetCmdQueryAvailableBorrow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "available-borrow [denom]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for the available amount to borrow of a specified denomination",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryAvailableBorrowRequest{
				Denom: args[0],
			}

			resp, err := queryClient.AvailableBorrow(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQuerySupplyAPY creates a Cobra command to query for the
// supply APY of a specific uToken.
func GetCmdQuerySupplyAPY() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "supply-apy [denom]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for the supply APY of a specified denomination",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QuerySupplyAPYRequest{
				Denom: args[0],
			}

			resp, err := queryClient.SupplyAPY(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryBorrowAPY creates a Cobra command to query for the
// borrow APY of a specific token.
func GetCmdQueryBorrowAPY() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "borrow-apy [denom]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for the borrow APY of a specified denomination",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryBorrowAPYRequest{
				Denom: args[0],
			}

			resp, err := queryClient.BorrowAPY(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryMarketSize creates a Cobra command to query for the
// Market Size of a specific token.
func GetCmdQueryMarketSize() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "market-size [denom]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for the USD market size of a specified denomination",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryMarketSizeRequest{
				Denom: args[0],
			}

			resp, err := queryClient.MarketSize(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryTokenMarketSize creates a Cobra command to query for the
// Market Size of a specific token, in token denomination instead of USD.
func GetCmdQueryTokenMarketSize() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token-market-size [denom]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for the market size of a specified denomination measured in base tokens",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryTokenMarketSizeRequest{
				Denom: args[0],
			}

			resp, err := queryClient.TokenMarketSize(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryBorrowLimit creates a Cobra command to query for the
// borrow limit of a specific borrower.
func GetCmdQueryBorrowLimit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "borrow-limit [addr]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for the borrow limit of a specified borrower",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryBorrowLimitRequest{
				Address: args[0],
			}

			resp, err := queryClient.BorrowLimit(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryLiquidationThreshold creates a Cobra command to query a
// liquidation threshold of a specific borrower.
func GetCmdQueryLiquidationThreshold() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liquidation-threshold [addr]",
		Args:  cobra.ExactArgs(1),
		Short: "Query a liquidation threshold of a specified borrower",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryLiquidationThresholdRequest{
				Address: args[0],
			}

			resp, err := queryClient.LiquidationThreshold(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryMarketSummary creates a Cobra command to query for the
// Market Summary of a specific token.
func GetCmdQueryMarketSummary() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "market-summary [denom]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for the market summary of a specified denomination",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryMarketSummaryRequest{
				Denom: args[0],
			}

			resp, err := queryClient.MarketSummary(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryLiquidationTargets creates a Cobra command to query for
// all eligible liquidation targets
func GetCmdQueryLiquidationTargets() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liquidation-targets",
		Args:  cobra.ExactArgs(0),
		Short: "Query for all borrower addresses eligible for liquidation",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryLiquidationTargetsRequest{}

			resp, err := queryClient.LiquidationTargets(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryTotalCollateral creates a Cobra command to query for the
// total collateral amount of a specific token.
func GetCmdQueryTotalCollateral() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-collateral [denom]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for the total amount of collateral of a uToken denomination",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			req := &types.QueryTotalCollateralRequest{
				Denom: args[0],
			}
			resp, err := queryClient.TotalCollateral(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryTotalBorrowed creates a Cobra command to query for the
// total borrowed amount of a specific token.
func GetCmdQueryTotalBorrowed() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-borrowed [denom]",
		Args:  cobra.ExactArgs(1),
		Short: "Query for the total amount borrowed of a token denomination",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			req := &types.QueryTotalBorrowedRequest{
				Denom: args[0],
			}
			resp, err := queryClient.TotalBorrowed(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(resp)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
