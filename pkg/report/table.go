package report

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/liamg/tml"
	"github.com/samber/lo"
	"github.com/xlab/treeprint"
	"golang.org/x/exp/slices"

	"github.com/aquasecurity/table"
	dbTypes "github.com/aquasecurity/trivy-db/pkg/types"
	ftypes "github.com/aquasecurity/trivy/pkg/fanal/types"
	"github.com/aquasecurity/trivy/pkg/log"
	"github.com/aquasecurity/trivy/pkg/types"
)

var (
	SeverityColor = []func(a ...interface{}) string{
		color.New(color.FgCyan).SprintFunc(),   // UNKNOWN
		color.New(color.FgBlue).SprintFunc(),   // LOW
		color.New(color.FgYellow).SprintFunc(), // MEDIUM
		color.New(color.FgHiRed).SprintFunc(),  // HIGH
		color.New(color.FgRed).SprintFunc(),    // CRITICAL
	}
)

// TableWriter implements Writer and output in tabular form
type TableWriter struct {
	Severities []dbTypes.Severity
	Output     io.Writer

	// Show dependency origin tree
	Tree bool

	// We have to show a message once about using the '-format json' subcommand to get the full pkgPath
	ShowMessageOnce *sync.Once

	// For misconfigurations
	IncludeNonFailures bool
	Trace              bool
}

// Write writes the result on standard output
func (tw TableWriter) Write(report types.Report) error {
	for _, result := range report.Results {
		// Not display a table of custom resources
		if result.Class == types.ClassCustom {
			continue
		}
		tw.write(result)
	}
	return nil
}

func (tw TableWriter) isOutputToTerminal() bool {
	if tw.Output != os.Stdout {
		return false
	}
	o, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (o.Mode() & os.ModeCharDevice) == os.ModeCharDevice
}

func (tw TableWriter) write(result types.Result) {
	tableWriter := table.New(tw.Output)
	if tw.isOutputToTerminal() { // use ansi output if we're not piping elsewhere
		tableWriter.SetHeaderStyle(table.StyleBold)
		tableWriter.SetLineStyle(table.StyleDim)
	}
	tableWriter.SetBorders(true)
	tableWriter.SetAutoMerge(true)
	tableWriter.SetRowLines(true)

	severityCount := tw.countSeverities(result)

	switch {
	case len(result.Vulnerabilities) > 0:
		tw.writeVulnerabilities(tableWriter, result.Vulnerabilities)
	case len(result.Secrets) > 0:
		tw.writeSecrets(tableWriter, result.Secrets)
	}

	total, summaries := tw.summary(severityCount)

	target := result.Target
	if result.Class == types.ClassConfig && len(result.Misconfigurations) == 0 {
		return
	} else if result.Class == types.ClassSecret {
		if len(result.Secrets) == 0 {
			return
		}
		target += " (secrets)"
	} else if result.Class != types.ClassOSPkg {
		target += fmt.Sprintf(" (%s)", result.Type)
	}

	if tw.isOutputToTerminal() {
		// nolint
		_ = tml.Printf("\n<underline><bold>%s</bold></underline>\n\n", target)
	} else {
		fmt.Printf("\n%s\n", target)
		fmt.Println(strings.Repeat("=", len(target)))
	}
	if result.Class == types.ClassConfig {
		// for misconfigurations
		summary := result.MisconfSummary
		fmt.Printf("Tests: %d (SUCCESSES: %d, FAILURES: %d, EXCEPTIONS: %d)\n",
			summary.Successes+summary.Failures+summary.Exceptions, summary.Successes, summary.Failures, summary.Exceptions)
		fmt.Printf("Failures: %d (%s)\n\n", total, strings.Join(summaries, ", "))
	} else {
		// for vulnerabilities and secrets
		fmt.Printf("Total: %d (%s)\n\n", total, strings.Join(summaries, ", "))
	}

	tableWriter.Render()

	if len(result.Misconfigurations) > 0 {
		_, _ = fmt.Fprint(tw.Output, NewMisconfigRenderer(result.Target, result.Misconfigurations, tw.IncludeNonFailures, tw.isOutputToTerminal()).Render())
	}

	if tw.Tree {
		tw.renderDependencyTree(result)
	}

	// For debugging
	if tw.Trace {
		tw.outputTrace(result)
	}

	return
}

func (tw TableWriter) summary(severityCount map[string]int) (int, []string) {
	var total int
	var severities []string
	for _, sev := range tw.Severities {
		severities = append(severities, sev.String())
	}

	var summaries []string
	for _, severity := range dbTypes.SeverityNames {
		if !slices.Contains(severities, severity) {
			continue
		}
		count := severityCount[severity]
		r := fmt.Sprintf("%s: %d", severity, count)
		summaries = append(summaries, r)
		total += count
	}

	return total, summaries
}

func (tw TableWriter) writeVulnerabilities(tableWriter *table.Table, vulns []types.DetectedVulnerability) {
	header := []string{"Library", "Vulnerability", "Severity", "Installed Version", "Fixed Version", "Title"}
	tableWriter.SetHeaders(header...)
	tw.setVulnerabilityRows(tableWriter, vulns)
}

func (tw TableWriter) setVulnerabilityRows(tableWriter *table.Table, vulns []types.DetectedVulnerability) {
	for _, v := range vulns {
		lib := v.PkgName
		if v.PkgPath != "" {
			fileName := filepath.Base(v.PkgPath)
			lib = fmt.Sprintf("%s (%s)", v.PkgName, fileName)
			tw.ShowMessageOnce.Do(func() {
				log.Logger.Infof("Table result includes only package filenames. Use '--format json' option to get the full path to the package file.")
			})
		}

		title := v.Title
		if title == "" {
			title = v.Description
		}
		splitTitle := strings.Split(title, " ")
		if len(splitTitle) >= 12 {
			title = strings.Join(splitTitle[:12], " ") + "..."
		}

		if len(v.PrimaryURL) > 0 {
			if tw.isOutputToTerminal() {
				title = tml.Sprintf("%s\n<blue>%s</blue>", title, v.PrimaryURL)
			} else {
				title = fmt.Sprintf("%s\n%s", title, v.PrimaryURL)
			}
		}

		var row []string
		if tw.isOutputToTerminal() {
			row = []string{lib, v.VulnerabilityID, ColorizeSeverity(v.Severity, v.Severity),
				v.InstalledVersion, v.FixedVersion, strings.TrimSpace(title)}
		} else {
			row = []string{lib, v.VulnerabilityID, v.Severity, v.InstalledVersion, v.FixedVersion, strings.TrimSpace(title)}
		}

		tableWriter.AddRow(row...)
	}
}
func (tw TableWriter) renderDependencyTree(result types.Result) {
	// Get parents of each dependency
	parents := reverseDeps(result.Packages)
	if len(parents) == 0 {
		return
	}

	root := treeprint.NewWithRoot(fmt.Sprintf(`
Dependency Origin Tree
======================
%s`, result.Target))

	// This count is next to the package ID.
	// e.g. node-fetch@1.7.3 (MEDIUM: 2, HIGH: 1, CRITICAL: 3)
	pkgSeverityCount := map[string]map[string]int{}
	for _, vuln := range result.Vulnerabilities {
		cnts, ok := pkgSeverityCount[vuln.PkgID]
		if !ok {
			cnts = map[string]int{}
		}

		cnts[vuln.Severity]++
		pkgSeverityCount[vuln.PkgID] = cnts
	}

	// Render tree
	seen := map[string]struct{}{}
	for _, vuln := range result.Vulnerabilities {
		if _, ok := seen[vuln.PkgID]; ok {
			continue
		}

		_, summaries := tw.summary(pkgSeverityCount[vuln.PkgID])
		topLvlID := fmt.Sprintf("%s, (%s)", vuln.PkgID, strings.Join(summaries, ", "))
		if tw.isOutputToTerminal() {
			topLvlID = color.HiRedString(topLvlID)
		}

		seen[vuln.PkgID] = struct{}{}
		branch := root.AddBranch(topLvlID)
		addParents(branch, vuln.PkgID, parents)

	}
	tw.Println(root.String())
}

func addParents(topItem treeprint.Tree, pkgID string, parentMap map[string][]string) {
	parents, ok := parentMap[pkgID]
	if !ok {
		return
	}
	for _, parent := range parents {
		branch := topItem.AddBranch(parent)
		addParents(branch, parent, parentMap)
	}
}

func reverseDeps(libs []ftypes.Package) map[string][]string {
	reversed := make(map[string][]string)
	for _, lib := range libs {
		for _, dependOn := range lib.DependsOn {
			items, ok := reversed[dependOn]
			if !ok {
				reversed[dependOn] = []string{lib.ID}
			} else {
				reversed[dependOn] = append(items, lib.ID)
			}
		}
	}

	for k, v := range reversed {
		reversed[k] = lo.Uniq(v)
	}
	return reversed
}

func (tw TableWriter) outputTrace(result types.Result) {
	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintfFunc()
	red := color.New(color.FgRed).SprintfFunc()

	for _, misconf := range result.Misconfigurations {
		if len(misconf.Traces) == 0 {
			continue
		}

		c := green
		if misconf.Status == types.StatusFailure {
			c = red
		}

		tw.Println(c("\nID: %s", misconf.ID))
		tw.Println(c("File: %s", result.Target))
		tw.Println(c("Namespace: %s", misconf.Namespace))
		tw.Println(c("Query: %s", misconf.Query))
		tw.Println(c("Message: %s", misconf.Message))
		for _, t := range misconf.Traces {
			tw.Println(blue("TRACE "), t)
		}
		tw.Println()
	}
}

func (tw TableWriter) writeSecrets(tableWriter *table.Table, secrets []ftypes.SecretFinding) {

	alignment := []table.Alignment{table.AlignCenter, table.AlignCenter, table.AlignCenter,
		table.AlignCenter, table.AlignLeft}
	header := []string{"Category", "Description", "Severity", "Line No", "Match"}

	tableWriter.SetAlignment(alignment...)
	tableWriter.SetHeaders(header...)
	tw.setSecretRows(tableWriter, secrets)
}

func (tw TableWriter) setSecretRows(tableWriter *table.Table, secrets []ftypes.SecretFinding) {
	for _, secret := range secrets {
		severity := secret.Severity
		if tw.isOutputToTerminal() {
			severity = ColorizeSeverity(severity, severity)
		}
		row := []string{string(secret.Category), secret.Title, severity,
			fmt.Sprint(secret.StartLine), // multi-line is not supported for now.
			secret.Match}

		tableWriter.AddRow(row...)
	}
}

func (tw TableWriter) Println(a ...interface{}) {
	_, _ = fmt.Fprintln(tw.Output, a...)
}

func (tw TableWriter) countSeverities(result types.Result) map[string]int {
	severityCount := map[string]int{}
	for _, misconf := range result.Misconfigurations {
		if misconf.Status == types.StatusFailure {
			severityCount[misconf.Severity]++
		}
	}
	for _, secret := range result.Secrets {
		severity := secret.Severity
		severityCount[severity]++
	}
	for _, v := range result.Vulnerabilities {
		severityCount[v.Severity]++
	}
	return severityCount
}

func ColorizeSeverity(value, severity string) string {
	for i, name := range dbTypes.SeverityNames {
		if severity == name {
			return SeverityColor[i](value)
		}
	}
	return color.New(color.FgBlue).SprintFunc()(severity)
}
