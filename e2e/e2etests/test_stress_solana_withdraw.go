func TestStressSolanaWithdraw(r *runner.E2ERunner, args []string) {
	require.Len(r, args, 2)

	withdrawSOLAmount := utils.ParseBigInt(r, args[0])
	numWithdrawalsSOL := utils.ParseInt(r, args[1])

	privKey := r.GetSolanaPrivKey()
	r.Logger.Print("Starting stress test of %d SOL withdrawals", numWithdrawalsSOL)

	totalApproveAmount := new(big.Int).Mul(withdrawSOLAmount, big.NewInt(int64(numWithdrawalsSOL)))
	tx, err := r.SOLZRC20.Approve(r.ZEVMAuth, r.SOLZRC20Addr, totalApproveAmount)
	require.NoError(r, err)
	receipt := utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
	utils.RequireTxSuccessful(r, receipt, "approve_sol")

	// Fetch the current nonce
	nonce, err := r.ZEVMClient.PendingNonceAt(r.Ctx, r.ZEVMAuth.From)
	require.NoError(r, err)

	durationsChan := make(chan float64, numWithdrawalsSOL)
	defer close(durationsChan)

	var eg errgroup.Group

	for i := 0; i < numWithdrawalsSOL; i++ {
		i := i
		currentNonce := nonce + uint64(i) // Increment nonce for each transaction

		eg.Go(func() error {
			startTime := time.Now()

			auth := r.ZEVMAuth
			auth.Nonce = big.NewInt(int64(currentNonce)) // Set unique nonce

			tx, err := r.SOLZRC20.Withdraw(auth, []byte(privKey.PublicKey().String()), withdrawSOLAmount)
			if err != nil {
				r.Logger.Print("Error in withdrawal %d: %v", i, err)
				return nil
			}

			receipt := utils.MustWaitForTxReceipt(r.Ctx, r.ZEVMClient, tx, r.Logger, r.ReceiptTimeout)
			if receipt.Status != 1 {
				r.Logger.Print("Error in transaction %d: receipt status is not successful: %s", i, tx.Hash().Hex())
				return fmt.Errorf("transaction failed")
			}

			r.Logger.Print("Index %d: Starting SOL withdraw, tx hash: %s", i, tx.Hash().Hex())

			cctx := utils.WaitCctxMinedByInboundHash(r.Ctx, tx.Hash().Hex(), r.CctxClient, r.Logger, r.ReceiptTimeout)
			if cctx.CctxStatus.Status != crosschaintypes.CctxStatus_OutboundMined {
				return fmt.Errorf(
					"Index %d: Withdraw cctx failed with status %s, message %s, cctx index %s",
					i,
					cctx.CctxStatus.Status,
					cctx.CctxStatus.StatusMessage,
					cctx.Index,
				)
			}

			timeToComplete := time.Since(startTime)
			r.Logger.Print("Index %d: Withdraw SOL cctx success in %s", i, timeToComplete.String())

			durationsChan <- timeToComplete.Seconds()
			return nil
		})
	}

	err = eg.Wait()
	require.NoError(r, err)

	var withdrawDurations []float64
	for duration := range durationsChan {
		withdrawDurations = append(withdrawDurations, duration)
	}

	desc, descErr := stats.Describe(withdrawDurations, false, &[]float64{50.0, 75.0, 90.0, 95.0})
	if descErr != nil {
		r.Logger.Print("❌ Failed to calculate latency report: %v", descErr)
	}

	r.Logger.Print("Latency report:")
	r.Logger.Print("Min:  %.2f", desc.Min)
	r.Logger.Print("Max:  %.2f", desc.Max)
	r.Logger.Print("Mean: %.2f", desc.Mean)
	r.Logger.Print("Std:  %.2f", desc.Std)
	for _, p := range desc.DescriptionPercentiles {
		r.Logger.Print("P%.0f:  %.2f", p.Percentile, p.Value)
	}

	r.Logger.Print("All SOL withdrawals completed")
}
