package crabbot

import (
	"bytes"
	"fmt"
	"github.com/Necroforger/dgrouter/exrouter"
	"github.com/olekukonko/tablewriter"
	"log"
)

type Help struct {
	router *exrouter.Route
}

func (h *Help) Handle(ctx *exrouter.Context) {
	var helpMessage string

	if h.router.Name != "" {
		helpMessage = fmt.Sprintf("## Help\n\nBelow are the commands for: %s\n\n", h.router.Name)

	} else {
		helpMessage = "## Help\n\nBelow are the root commands \n\n"

	}

	table := h.renderMarkDownTable()

	_, err := ctx.Reply("```" + helpMessage + table + "```")
	if err != nil {
		log.Print("Something went wrong when handling the help request", err)
	}
}

func (h *Help) GetRouteCommand() string {
	return "help"
}

func (h *Help) GetDescription() string {
	return "prints this help menu"
}

func (h *Help) Register(router *exrouter.Route) *exrouter.Route {
	h.router = router
	return router.On(h.GetRouteCommand(), h.Handle)
}

func (h *Help) renderMarkDownTable() string {
	var tableData [][]string

	// Iterate over the myRoutes and add them to the table tableData
	for _, v := range h.router.Routes {
		row := []string{v.Name, v.Description}
		tableData = append(tableData, row)
	}

	// Create a buffer so we can capture table output to a string later
	buffer := new(bytes.Buffer)

	table := tablewriter.NewWriter(buffer)
	table.SetHeader([]string{"Command Name", "Command Description"})
	table.SetColWidth(40)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(tableData) // Add Bulk Data
	table.Render()

	return buffer.String()
}

func (h *Help) GetSubRoutes() []RouteHandler {
	return nil
}

func NewHelp() *Help {
	return &Help{}
}
