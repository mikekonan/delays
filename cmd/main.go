package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mikekonan/delays"
)

func main() {
	strategy := flag.String("strategy", "exponential", "Delay strategy (exponential, linear)")
	totalDuration := flag.Duration("totalDuration", 30*time.Minute, "Total duration")
	numAttempts := flag.Int("numAttempts", 10, "Number of attempts")
	exponent := flag.Float64("exponent", 2.0, "Exponent for the exponential strategy (only used in exponential strategy)")
	htmlOutput := flag.String("htmlOutput", "", "File path to save the HTML output")

	flag.Parse()

	var delayStrategy delays.DelayStrategy
	switch *strategy {
	case "exponential":
		delayStrategy = delays.Exponential(*totalDuration, *numAttempts, *exponent)
	default:
		fmt.Println("Unknown strategy:", *strategy)
		return
	}

	delay := delays.New(delayStrategy)

	if *htmlOutput != "" {
		generateHTMLOutput(*htmlOutput, delay.Plan())
	} else {
		printPlanWithElapsed(delay.Plan())
	}
}

func printPlanWithElapsed(plan []time.Duration) {
	fmt.Println("Plan of Delays:")
	fmt.Println("================")
	elapsed := time.Duration(0)
	for i, delay := range plan {
		elapsed += delay
		fmt.Printf("Attempt %2d: wait %v. Elapsed %v.\n", i+1, delay, elapsed)
	}
	fmt.Println("================")
}

func generateHTMLOutput(filePath string, plan []time.Duration) {
	htmlContent := generateHTML(plan)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating HTML file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(htmlContent)
	if err != nil {
		fmt.Println("Error writing to HTML file:", err)
	}
}

func generateHTML(plan []time.Duration) string {
	var attempts []string
	var waits []string
	var elapsedTimes []string

	elapsed := 0.0
	for i, delay := range plan {
		attempts = append(attempts, fmt.Sprintf("%d", i+1))
		wait := delay.Seconds()
		waits = append(waits, fmt.Sprintf("%.2f", wait))
		elapsedTimes = append(elapsedTimes, fmt.Sprintf("%.2f", elapsed))
		elapsed += wait
	}

	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>Delays Visualization</title>
    <script src="https://cdn.jsdelivr.net/npm/echarts/dist/echarts.min.js"></script>
</head>
<body>
    <div id="chart" style="width: 100%%; height: 600px;"></div>
    <script type="text/javascript">
        var chart = echarts.init(document.getElementById('chart'));
        var option = {
            title: {
                text: 'Delays Visualization'
            },
            tooltip: {
                trigger: 'axis',
                formatter: function (params) {
                    var content = params[0].axisValueLabel + '<br/>';
                    params.forEach(function (item) {
                        var value = item.data;
                        var seconds = value.toFixed(2) + ' s';
                        var minutes = (value / 60).toFixed(2) + ' min';
                        var hours = (value / 3600).toFixed(2) + ' h';
                        content += item.marker + item.seriesName + ': ' + seconds + ' / ' + minutes + ' / ' + hours + '<br/>';
                    });
                    return content;
                }
            },
            legend: {
                data: ['Wait Time', 'Elapsed Time']
            },
            xAxis: {
                type: 'category',
                data: [%s]
            },
            yAxis: {
                type: 'value',
                axisLabel: {
                    formatter: '{value} s'
                }
            },
            series: [
                {
                    name: 'Wait Time',
                    type: 'bar',
                    data: [%s]
                },
                {
                    name: 'Elapsed Time',
                    type: 'line',
                    data: [%s]
                }
            ]
        };
        chart.setOption(option);
    </script>
</body>
</html>
`, strings.Join(attempts, ","), strings.Join(waits, ","), strings.Join(elapsedTimes, ","))
}
