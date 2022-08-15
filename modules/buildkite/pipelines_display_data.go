package buildkite

import (
	"fmt"
	"sort"
	"strings"
)

type pipelinesDisplayData struct {
	buildsForPipeline map[string][]Build
	orderedPipelines  []string
}

func (data *pipelinesDisplayData) Content() string {
	maxPipelineLength := getLongestLength(data.orderedPipelines)

	str := ""
	for _, pipeline := range data.orderedPipelines {
		str += fmt.Sprintf("[white]%s", padRight(pipeline, maxPipelineLength))
		for _, build := range data.buildsForPipeline[pipeline] {
			str += fmt.Sprintf("  [%s]%s[white]", buildColor(build.State), build.Branch)
		}
		str += "\n"
	}

	return str
}

func newPipelinesDisplayData(builds []Build) pipelinesDisplayData {
	grouped := make(map[string][]Build)

	for _, build := range builds {
		if _, ok := grouped[build.Pipeline.Slug]; ok {
			grouped[build.Pipeline.Slug] = append(grouped[build.Pipeline.Slug], build)
		} else {
			grouped[build.Pipeline.Slug] = []Build{}
			grouped[build.Pipeline.Slug] = append(grouped[build.Pipeline.Slug], build)
		}
	}

	orderedPipelines := make([]string, len(grouped))
	i := 0
	for pipeline := range grouped {
		orderedPipelines[i] = pipeline
		i++
	}
	sort.Strings(orderedPipelines)

	name := func(b1, b2 *Build) bool {
		return b1.Branch < b2.Branch
	}
	for _, builds := range grouped {
		ByBuild(name).Sort(builds)
	}

	return pipelinesDisplayData{
		buildsForPipeline: grouped,
		orderedPipelines:  orderedPipelines,
	}
}

type ByBuild func(b1, b2 *Build) bool

func (by ByBuild) Sort(builds []Build) {
	sorter := &buildSorter{
		builds: builds,
		by:     by,
	}
	sort.Sort(sorter)
}

type buildSorter struct {
	builds []Build
	by     func(b1, b2 *Build) bool
}

func (bs *buildSorter) Len() int {
	return len(bs.builds)
}

func (bs *buildSorter) Swap(i, j int) {
	bs.builds[i], bs.builds[j] = bs.builds[j], bs.builds[i]
}

func (bs *buildSorter) Less(i, j int) bool {
	return bs.by(&bs.builds[i], &bs.builds[j])
}

func getLongestLength(strs []string) int {
	longest := 0

	for _, str := range strs {
		if len(str) > longest {
			longest = len(str)
		}
	}

	return longest
}

func padRight(text string, length int) string {
	padLength := length - len(text)

	if padLength <= 0 {
		return text[:length]
	}

	return text + strings.Repeat(" ", padLength)
}

func buildColor(state string) string {
	switch state {
	case "passed":
		return "green"
	case "failed":
		return "red"
	default:
		return "yellow"
	}
}
