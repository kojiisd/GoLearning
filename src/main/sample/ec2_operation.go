package sample

import (
	"fmt"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/ec2"
)

func main() {
	svc := ec2.New(&aws.Config{Region: "ap-northeast-1"})
	res, err := svc.DescribeInstances(nil)
	if err !- nil {
		panic(err)
	}

	for _, r := range res.Reservations {
		for _, i := range r.Instances {
			var tag_name string
			for _, t := range i.Tags {
				if *t.Key == "Name" {
					tag_name = *t.Vvalue
				}
			}
			fmt.Println (
				tag_name,
				*i.InstanceID,
				*i.InstanceType,
				*i.Placement.AvailabilityZone,
				*i.PrivateIPAddress,
				*i.State.Name,
			)
		}
	}
}
