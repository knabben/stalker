package tui

import (
	"fmt"
	"github.com/knabben/stalker/pkg/testgrid"
	"regexp"
	"strings"
	"time"
)

func (d *DashboardIssue) RenderVisual(counter int) error {
	var (
		failingThreshold, flakeThreshold = 2, 3
		flakeStatus                      = make(map[string]string)
		isFlake                          bool
	)

	icon, state := "🟪", "Flaking"
	if d.Dashboard.OverallStatus == testgrid.FAILING_STATUS {
		icon, state = "🟥", "Failing"

	}
	testAgg := fmt.Sprintf("%s#%s", d.Dashboard.DashboardName, d.Tab)
	boardLink := fmt.Sprintf("https://testgrid.k8s.io/%s&exclude-non-failed-tests=", testAgg)

	for _, test := range d.Table.Tests {
		lastTimestmap := d.Table.Timestamps[0]

		prowURL := fmt.Sprintf("https://prow.k8s.io/view/gs/%s/%s", d.Table.Query, d.Table.Changelists[0])

		_, num := RenderStatuses(&test, d.Table.Timestamps)
		if (num >= failingThreshold && d.Dashboard.OverallStatus == testgrid.FAILING_STATUS) || (num >= flakeThreshold && d.Dashboard.OverallStatus == testgrid.FLAKY_STATUS) {
			testName := test.Name
			if strings.Contains(test.Name, "Kubernetes e2e suite.[It]") {
				params := getParameter(testRegex, testName)
				testName = params["TEST"]
			}
			triageLink := fmt.Sprintf("https://storage.googleapis.com/k8s-triage/index.html?test=%s", testName)
			unixTimeUTC := time.Unix(lastTimestmap/1000, 0)
			item := fmt.Sprintf("%s %s on [%s](%s): `%s` [Prow](%s), [Triage](%s), last failure on %s\n", icon, state, testAgg, boardLink, test.Name, prowURL, triageLink, unixTimeUTC.Format(time.RFC3339))
			//https: //prow.k8s.io/view/gs/kubernetes-jenkins/logs/ci-kubernetes-e2e-ubuntu-gce-containerd/1841461726814408704
			flakeStatus[item] = ""
			isFlake = true
		}
	}

	if isFlake {
		params := getParameter(summaryRegex, d.Dashboard.Status)
		fmt.Println(style.Render(fmt.Sprintf("%d) %s %s \t\t\t %s", counter, d.Dashboard.DashboardName, params["PERCENT"], boardLink)))
		fmt.Println("")
		for item, _ := range flakeStatus {
			fmt.Println(bold.Render(item))
		}
	}

	return nil
}

func RenderStatuses(test *testgrid.Test, timestamps []int64) (text string, num int) {
	for i, t := range test.ShortTexts {
		if t != "" {
			tm := time.Unix(timestamps[i]/1000, 0)
			msg := test.Messages[i]
			text += fmt.Sprintf("\t%s %s %s\n", keyStyle.Render(t), tm, msg)
			num += 1
		}
	}
	return
}

func getParameter(regEx, value string) (paramsMap map[string]string) {
	var r = regexp.MustCompile(regEx)
	match := r.FindStringSubmatch(value)
	paramsMap = make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}
