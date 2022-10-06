/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/mikuta0407/misskey-cli/misskey"
	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("note called")

		client := misskey.NewClient(instanceName, cfgFile)

		if deleteId == "" && replyId == "" {
			if len(args) > 1 {
				fmt.Println("too many args")
				return
			}
			if len(args) == 0 {
				fmt.Println("Please write note")
				return
			}

			if err := client.CreateNote(args[0]); err != nil {
				fmt.Println(err)
			}
		} else if deleteId != "" && replyId == "" {
			if err := client.DeleteNote(deleteId); err != nil {
				fmt.Println(err)
			}
		} else if deleteId == "" && replyId != "" {
			if len(args) > 1 {
				fmt.Println("too many args")
				return
			}
			if len(args) == 0 {
				fmt.Println("Please write note")
				return
			}
			if err := client.ReplyNote(replyId, args[0]); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Please one Option")
			return
		}
	},
}

var (
	replyId  string
	deleteId string
)

func init() {
	rootCmd.AddCommand(noteCmd)

	// 削除
	noteCmd.Flags().StringVarP(&deleteId, "delete", "d", "", "Delete notes (id)")

	// リプライ
	noteCmd.Flags().StringVarP(&replyId, "reply", "r", "", "reply note")

	// 公開範囲の話
	//noteCmd.Flags().StringVarP(&reply, "", "", "", "")

}
