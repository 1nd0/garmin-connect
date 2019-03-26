package main

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	connectionsCmd := &cobra.Command{
		Use: "connections",
	}
	rootCmd.AddCommand(connectionsCmd)

	connectionsListCmd := &cobra.Command{
		Use:   "list [display name]",
		Short: "List all connections",
		Run:   connectionsList,
		Args:  cobra.RangeArgs(0, 1),
	}
	connectionsCmd.AddCommand(connectionsListCmd)

	connectionsPendingCmd := &cobra.Command{
		Use:   "pending",
		Short: "List pending connections",
		Run:   connectionsPending,
		Args:  cobra.NoArgs,
	}
	connectionsCmd.AddCommand(connectionsPendingCmd)

	connectionsRemoveCmd := &cobra.Command{
		Use:   "remove <display name>",
		Short: "Remove a connection",
		Run:   connectionsRemove,
		Args:  cobra.ExactArgs(1),
	}
	connectionsCmd.AddCommand(connectionsRemoveCmd)

	connectionsSearchCmd := &cobra.Command{
		Use:   "search <keyword>",
		Short: "Search Garmin wide for a person",
		Run:   connectionsSearch,
		Args:  cobra.ExactArgs(1),
	}
	connectionsCmd.AddCommand(connectionsSearchCmd)

	connectionsAcceptCmd := &cobra.Command{
		Use:   "accept <request id>",
		Short: "Accept a connection request",
		Run:   connectionsAccept,
		Args:  cobra.ExactArgs(1),
	}
	connectionsCmd.AddCommand(connectionsAcceptCmd)

	connectionsRequestCmd := &cobra.Command{
		Use:   "request <display name>",
		Short: "Request connectio from another user",
		Run:   connectionsRequest,
		Args:  cobra.ExactArgs(1),
	}
	connectionsCmd.AddCommand(connectionsRequestCmd)
}

func connectionsList(_ *cobra.Command, args []string) {
	displayName := ""
	if len(args) == 1 {
		displayName = args[0]
	}

	connections, err := client.Connections(displayName)
	bail(err)

	t := NewTable()
	t.AddHeader("Display Name", "Name", "Location", "Profile Image")
	for _, c := range connections {
		t.AddRow(c.DisplayName, c.Fullname, c.Location, c.ProfileImageURLMedium)
	}
	t.Output(os.Stdout)
}

func connectionsPending(_ *cobra.Command, _ []string) {
	connections, err := client.PendingConnections()
	bail(err)

	t := NewTable()
	t.AddHeader("RequestID", "Display Name", "Name", "Location", "Profile Image")
	for _, c := range connections {
		t.AddRow(c.ConnectionRequestID, c.DisplayName, c.Fullname, c.Location, c.ProfileImageURLMedium)
	}
	t.Output(os.Stdout)
}

func connectionsRemove(_ *cobra.Command, args []string) {
	connectionRequestID, _ := strconv.Atoi(args[0])
	err := client.RemoveConnection(connectionRequestID)
	bail(err)
}

func connectionsSearch(_ *cobra.Command, args []string) {
	keyword := args[0]
	connections, err := client.SearchConnections(keyword)
	bail(err)

	t := NewTabular()
	for _, c := range connections {
		t.AddValue(c.DisplayName, c.Fullname)
	}
	t.Output(os.Stdout)
}

func connectionsAccept(_ *cobra.Command, args []string) {
	connectionRequestID, err := strconv.Atoi(args[0])
	bail(err)

	err = client.AcceptConnection(connectionRequestID)
	bail(err)
}

func connectionsRequest(_ *cobra.Command, args []string) {
	displayName := args[0]

	err := client.RequestConnection(displayName)
	bail(err)
}
