// Copyright 2016 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/FactomProject/factom"
)

// newtx creates a new transaction in the wallet.
var newtx = func() *fctCmd {
	cmd := new(fctCmd)
	cmd.helpMsg = "factom-cli newtx TXNAME"
	cmd.description = "Create a new transaction in the wallet"
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()

		if len(args) != 1 {
			fmt.Println(cmd.helpMsg)
			return
		}
		if err := factom.NewTransaction(args[0]); err != nil {
			errorln(err)
			return
		}
	}
	help.Add("newtx", cmd)
	return cmd
}()

// rmtx removes a transaction in the wallet.
var rmtx = func() *fctCmd {
	cmd := new(fctCmd)
	cmd.helpMsg = "factom-cli rmtx TXNAME"
	cmd.description = "Remove a transaction in the wallet"
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()

		if len(args) != 1 {
			fmt.Println(cmd.helpMsg)
			return
		}
		if err := factom.DeleteTransaction(args[0]); err != nil {
			errorln(err)
			return
		}
	}
	help.Add("rmtx", cmd)
	return cmd
}()

// listtxs lists the working transactions in the wallet.
var listtxs = func() *fctCmd {
	cmd := new(fctCmd)
	cmd.helpMsg = "factom-cli listtxs"
	cmd.description = "List current working transactions in the wallet"
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()

		txs, err := factom.ListTransactions()
		if err != nil {
			errorln(err)
			return
		}
		for _, tx := range txs {
			fmt.Println(tx)
		}
	}
	help.Add("listtxs", cmd)
	return cmd
}()

// addtxinput adds a factoid input to a transaction in the wallet.
var addtxinput = func() *fctCmd {
	cmd := new(fctCmd)
	cmd.helpMsg = "factom-cli addtxinput TXNAME ADDRESS AMOUNT"
	cmd.description = "Add a Factoid input to a transaction in the wallet"
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()

		if len(args) != 3 {
			fmt.Println(cmd.helpMsg)
			return
		}
		var amt uint64
		if i, err := strconv.Atoi(args[2]); err != nil {
			errorln(err)
		} else if i < 0 {
			errorln("AMMOUNT may not be less than 0")
		} else {
			amt = uint64(i)
		}
		if err := factom.AddTransactionInput(args[0], args[1], amt); err != nil {
			errorln(err)
			return
		}
	}
	help.Add("addtxinput", cmd)
	return cmd
}()

// addtxoutput adds a factoid output to a transaction in the wallet.
var addtxoutput = func() *fctCmd {
	cmd := new(fctCmd)
	cmd.helpMsg = "factom-cli addtxoutput TXNAME ADDRESS AMOUNT"
	cmd.description = "Add a Factoid output to a transaction in the wallet"
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()

		if len(args) != 3 {
			fmt.Println(cmd.helpMsg)
			return
		}
		var amt uint64
		if i, err := strconv.Atoi(args[2]); err != nil {
			errorln(err)
		} else if i < 0 {
			errorln("AMMOUNT may not be less than 0")
		} else {
			amt = uint64(i)
		}
		if err := factom.AddTransactionOutput(args[0], args[1], amt); err != nil {
			errorln(err)
			return
		}
	}
	help.Add("addtxoutput", cmd)
	return cmd
}()

// addtxecoutput adds an entry credit output to a transaction in the wallet.
var addtxecoutput = func() *fctCmd {
	cmd := new(fctCmd)
	cmd.helpMsg = "factom-cli addtxecoutput TXNAME ADDRESS AMOUNT"
	cmd.description = "Add an Entry Credit output to a transaction in the wallet"
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()

		if len(args) != 3 {
			fmt.Println(cmd.helpMsg)
			return
		}
		var amt uint64
		if i, err := strconv.Atoi(args[2]); err != nil {
			errorln(err)
		} else if i < 0 {
			errorln("AMMOUNT may not be less than 0")
		} else {
			amt = uint64(i)
		}
		if err := factom.AddTransactionECOutput(args[0], args[1], amt); err != nil {
			errorln(err)
			return
		}
	}
	help.Add("addtxecoutput", cmd)
	return cmd
}()

// addtxfee adds an entry credit output to a transaction in the wallet.
var addtxfee = func() *fctCmd {
	cmd := new(fctCmd)
	cmd.helpMsg = "factom-cli addtxfee TXNAME ADDRESS"
	cmd.description = "Add the transaction fee to an input of a transaction in the wallet"
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()

		if len(args) != 2 {
			fmt.Println(cmd.helpMsg)
			return
		}
		if err := factom.AddTransactionFee(args[0], args[1]); err != nil {
			errorln(err)
			return
		}
	}
	help.Add("addtxfee", cmd)
	return cmd
}()

// subtxfee adds an entry credit output to a transaction in the wallet.
var subtxfee = func() *fctCmd {
	cmd := new(fctCmd)
	cmd.helpMsg = "factom-cli subtxfee TXNAME ADDRESS"
	cmd.description = "Subtranct the transaction fee to an input of a transaction in the wallet"
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()

		if len(args) != 2 {
			fmt.Println(cmd.helpMsg)
			return
		}
		if err := factom.SubTransactionFee(args[0], args[1]); err != nil {
			errorln(err)
			return
		}
	}
	help.Add("subtxfee", cmd)
	return cmd
}()

// signtx signs a transaction in the wallet.
var signtx = func() *fctCmd {
	cmd := new(fctCmd)
	cmd.helpMsg = "factom-cli signtx TXNAME"
	cmd.description = "Sign a transaction in the wallet"
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()

		if len(args) != 1 {
			fmt.Println(cmd.helpMsg)
			return
		}
		if err := factom.SignTransaction(args[0]); err != nil {
			errorln(err)
			return
		}
	}
	help.Add("signtx", cmd)
	return cmd
}()

// composetx composes the signed json rpc object to make a transaction against factomd
var composetx = func() *fctCmd {
	cmd := new(fctCmd)
	cmd.helpMsg = "factom-cli composetx TXNAME"
	cmd.description = "Compose a wallet transaction into a json rpc object"
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()

		if len(args) != 1 {
			fmt.Println(cmd.helpMsg)
			return
		}
		p, err := factom.ComposeTransaction(args[0])
		if err != nil {
			errorln(err)
			return
		}
		fmt.Println(string(p))
	}
	help.Add("composetx", cmd)
	return cmd
}()
