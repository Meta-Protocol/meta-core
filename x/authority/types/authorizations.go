package types

import (
	"fmt"

	"cosmossdk.io/errors"
)

var (
	OperationPolicyMessages = []string{
		"/zetachain.zetacore.crosschain.MsgRefundAbortedCCTX",
		"/zetachain.zetacore.crosschain.MsgAbortStuckCCTX",
		"/zetachain.zetacore.crosschain.MsgUpdateRateLimiterFlags",
		"/zetachain.zetacore.crosschain.MsgWhitelistERC20",
		"/zetachain.zetacore.fungible.MsgDeployFungibleCoinZRC20",
		"/zetachain.zetacore.fungible.MsgDeploySystemContracts",
		"/zetachain.zetacore.fungible.MsgRemoveForeignCoin",
		"/zetachain.zetacore.fungible.MsgUpdateZRC20LiquidityCap",
		"/zetachain.zetacore.fungible.MsgUpdateZRC20WithdrawFee",
		"/zetachain.zetacore.fungible.MsgUnpauseZRC20",
		"/zetachain.zetacore.observer.MsgAddObserver",
		"/zetachain.zetacore.observer.MsgRemoveChainParams",
		"/zetachain.zetacore.observer.MsgResetChainNonces",
		"/zetachain.zetacore.observer.MsgUpdateChainParams",
		"/zetachain.zetacore.observer.MsgEnableCCTX",
		"/zetachain.zetacore.observer.MsgUpdateGasPriceIncreaseFlags",
		"/zetachain.zetacore.lightclient.MsgEnableHeaderVerification",
	}
	AdminPolicyMessages = []string{
		"/zetachain.zetacore.crosschain.MsgMigrateTssFunds",
		"/zetachain.zetacore.crosschain.MsgUpdateTssAddress",
		"/zetachain.zetacore.fungible.MsgUpdateContractBytecode",
		"/zetachain.zetacore.fungible.MsgUpdateSystemContract",
		"/zetachain.zetacore.observer.MsgUpdateObserver",
	}
	EmergencyPolicyMessages = []string{
		"/zetachain.zetacore.crosschain.MsgAddInboundTracker",
		"/zetachain.zetacore.crosschain.MsgAddOutboundTracker",
		"/zetachain.zetacore.crosschain.MsgRemoveOutboundTracker",
		"/zetachain.zetacore.fungible.MsgPauseZRC20",
		"/zetachain.zetacore.observer.MsgUpdateKeygen",
		"/zetachain.zetacore.observer.MsgDisableCCTX",
		"/zetachain.zetacore.lightclient.MsgDisableHeaderVerification",
	}
)

// DefaultAuthorizationsList list is the list of authorizations that presently exist in the system.
// This is the minimum set of authorizations that are required to be set when the authorization table is deployed
func DefaultAuthorizationsList() AuthorizationList {
	authorizations := make(
		[]Authorization,
		len(OperationPolicyMessages)+len(AdminPolicyMessages)+len(EmergencyPolicyMessages),
	)
	index := 0
	for _, msgURL := range OperationPolicyMessages {
		authorizations[index] = Authorization{
			MsgUrl:           msgURL,
			AuthorizedPolicy: PolicyType_groupOperational,
		}
		index++
	}
	for _, msgURL := range AdminPolicyMessages {
		authorizations[index] = Authorization{
			MsgUrl:           msgURL,
			AuthorizedPolicy: PolicyType_groupAdmin,
		}
		index++
	}
	for _, msgURL := range EmergencyPolicyMessages {
		authorizations[index] = Authorization{
			MsgUrl:           msgURL,
			AuthorizedPolicy: PolicyType_groupEmergency,
		}
		index++
	}

	return AuthorizationList{
		Authorizations: authorizations,
	}
}

// SetAuthorizations adds the authorization to the list. If the authorization already exists, it updates the policy.
func (a *AuthorizationList) SetAuthorizations(authorization Authorization) {
	for i, auth := range a.Authorizations {
		if auth.MsgUrl == authorization.MsgUrl {
			a.Authorizations[i].AuthorizedPolicy = authorization.AuthorizedPolicy
			return
		}
	}
	a.Authorizations = append(a.Authorizations, authorization)
}

// RemoveAuthorizations removes the authorization from the list. It does not check if the authorization exists or not.
func (a *AuthorizationList) RemoveAuthorizations(authorization Authorization) {
	for i, auth := range a.Authorizations {
		if auth.MsgUrl == authorization.MsgUrl {
			a.Authorizations = append(a.Authorizations[:i], a.Authorizations[i+1:]...)
		}
	}
}

// GetAuthorizedPolicy returns the policy for the given message url.If the message url is not found,
// it returns an error and the first value of the enum.
func (a *AuthorizationList) GetAuthorizedPolicy(msgURL string) (PolicyType, error) {
	for _, auth := range a.Authorizations {
		if auth.MsgUrl == msgURL {
			return auth.AuthorizedPolicy, nil
		}
	}
	fmt.Println("Authorization not found", msgURL)
	// Returning first value of enum, can consider adding a default value of `EmptyPolicy` in the enum.
	return PolicyType(0), ErrAuthorizationNotFound
}

// Validate checks if the authorization list is valid. It returns an error if the message url is duplicated with different policies.
// It does not check if the list is empty or not, as an empty list is also considered valid.
func (a *AuthorizationList) Validate() error {
	checkMsgUrls := make(map[string]bool)
	for _, authorization := range a.Authorizations {
		if checkMsgUrls[authorization.MsgUrl] {
			return errors.Wrap(
				ErrInvalidAuthorizationList,
				fmt.Sprintf("duplicate message url: %s", authorization.MsgUrl),
			)
		}
		checkMsgUrls[authorization.MsgUrl] = true
	}
	return nil
}
