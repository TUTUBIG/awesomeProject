package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"log"
)

func TestGetDNSRecords() {
	mySession := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(),
	}))

	// Create a Route53 client with additional configuration
	svc := route53.New(mySession, aws.NewConfig().WithRegion(route53.ResourceRecordSetRegionCnNorth1))

	out, err := svc.ListResourceRecordSets(&route53.ListResourceRecordSetsInput{HostedZoneId: aws.String("Z07850576G00VDWKXAFO")})
	if err != nil {
		log.Fatal(err)
	}
	for _, rs := range out.ResourceRecordSets {
		log.Println(rs.Name, rs.Type)
	}

	out1, e := svc.ChangeResourceRecordSets(&route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String(route53.ChangeActionCreate),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name:   aws.String("test"),
						Region: aws.String(route53.ResourceRecordSetRegionCnNorth1),
						ResourceRecords: []*route53.ResourceRecord{
							{
								Value: aws.String(""),
							},
							{
								Value: aws.String(""),
							},
							{
								Value: aws.String(""),
							},
						},
						Type: aws.String(route53.RRTypeA),
					},
				},
			},
			Comment: aws.String("test"),
		},
		HostedZoneId: aws.String("Z07850576G00VDWKXAFO"),
	})
	if e != nil {
		log.Fatal(err)
	}
	log.Println(out1.ChangeInfo.Status)

	out, err = svc.ListResourceRecordSets(&route53.ListResourceRecordSetsInput{HostedZoneId: aws.String("Z07850576G00VDWKXAFO")})
	if err != nil {
		log.Fatal(err)
	}
	for _, rs := range out.ResourceRecordSets {
		log.Println(rs.Name, rs.Type)
	}
}
