package describe

import (
	"github.com/deviceinsight/kafkactl/operations/consumergroups"
	"github.com/spf13/cobra"
)

var consumerGroupFlags consumergroups.DescribeConsumerGroupFlags

var cmdDescribeConsumerGroup = &cobra.Command{
	Use:     "consumer-group GROUP",
	Aliases: []string{"cg"},
	Short:   "describe a consumerGroup",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		(&consumergroups.ConsumerGroupOperation{}).DescribeConsumerGroup(consumerGroupFlags, args[0])
	},
}

func init() {
	cmdDescribeConsumerGroup.Flags().BoolVarP(&consumerGroupFlags.OnlyPartitionsWithLag, "only-with-lag", "l", false, "show only partitions that have a lag")
	cmdDescribeConsumerGroup.Flags().StringVarP(&consumerGroupFlags.FilterTopic, "topic", "t", "", "show group details for given topic only")
	cmdDescribeConsumerGroup.Flags().StringVarP(&consumerGroupFlags.OutputFormat, "output", "o", flags.OutputFormat, "output format. One of: json|yaml|wide")
	cmdDescribeConsumerGroup.Flags().BoolVarP(&consumerGroupFlags.PrintTopics, "print-topics", "T", true, "print topic details")
	cmdDescribeConsumerGroup.Flags().BoolVarP(&consumerGroupFlags.PrintMembers, "print-members", "m", true, "print group members")
}
