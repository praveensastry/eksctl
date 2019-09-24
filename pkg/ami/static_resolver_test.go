package ami_test

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/weaveworks/eksctl/pkg/ami"
)

var _ = Describe("AMI Static Resolution", func() {
	type ResolveCase struct {
		Region       string
		Version      string
		InstanceType string
		ImageFamily  string
		ExpectedAMI  string
		ExpectError  bool
	}

	DescribeTable("When resolving an AMI using the default resolvers",
		func(c ResolveCase) {
			actualAmi, err := ami.Resolve(c.Region, c.Version, c.InstanceType, c.ImageFamily)
			Expect(actualAmi).Should(Equal(c.ExpectedAMI))
			if c.ExpectError {
				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError(ami.NewErrFailedResolution(c.Region, c.Version, c.InstanceType, c.ImageFamily)))
				errorType := reflect.TypeOf(err).Elem().Name()
				Expect(errorType).To(Equal("ErrFailedResolution"))
			} else {
				Expect(err).ShouldNot(HaveOccurred())
			}
		},
		Entry("with non-gpu instance and us-west-2", ResolveCase{
			Region:       "us-west-2",
			Version:      "1.12",
			InstanceType: "t2.medium",
			ImageFamily:  "AmazonLinux2",
			ExpectedAMI:  "ami-0315dd35bf204311d",
			ExpectError:  false,
		}),
		Entry("with non-gpu instance and us-east-1", ResolveCase{
			Region:       "us-east-1",
			Version:      "1.12",
			InstanceType: "t2.medium",
			ImageFamily:  "AmazonLinux2",
			ExpectedAMI:  "ami-0259ce67309f76e0b",
			ExpectError:  false,
		}),
		Entry("with non-gpu instance and eu-west-1", ResolveCase{
			Region:       "eu-west-1",
			Version:      "1.12",
			InstanceType: "t2.medium",
			ImageFamily:  "AmazonLinux2",
			ExpectedAMI:  "ami-0a6be9528ebb8999d",
			ExpectError:  false,
		}),
		Entry("with non-gpu instance and non-eks enabled region", ResolveCase{
			Region:       "sa-east-1",
			Version:      "1.12",
			InstanceType: "t2.medium",
			ImageFamily:  "AmazonLinux2",
			ExpectedAMI:  "",
			ExpectError:  true,
		}),
		Entry("with gpu (p2) instance and us-west-2", ResolveCase{
			Region:       "us-west-2",
			Version:      "1.12",
			InstanceType: "p2.xlarge",
			ImageFamily:  "AmazonLinux2",
			ExpectedAMI:  "ami-04f40a6dd1cd12b3c",
			ExpectError:  false,
		}),
		Entry("with gpu (p3) instance and us-east-1", ResolveCase{
			Region:       "us-east-1",
			Version:      "1.12",
			InstanceType: "p3.2xlarge",
			ImageFamily:  "AmazonLinux2",
			ExpectedAMI:  "ami-0585963de7ab8b964",
			ExpectError:  false,
		}),
		Entry("with gpu (p2) instance and eu-west-1", ResolveCase{
			Region:       "eu-west-1",
			Version:      "1.12",
			InstanceType: "p2.xlarge",
			ImageFamily:  "AmazonLinux2",
			ExpectedAMI:  "ami-049574de1981b69dc",
			ExpectError:  false,
		}),
		Entry("with gpu (g3) instance and eu-west-1", ResolveCase{
			Region:       "eu-west-1",
			Version:      "1.12",
			InstanceType: "g3.4xlarge",
			ImageFamily:  "AmazonLinux2",
			ExpectedAMI:  "ami-049574de1981b69dc",
			ExpectError:  false,
		}),
		Entry("with gpu (g4) instance and eu-west-1", ResolveCase{
			Region:       "eu-west-1",
			Version:      "1.12",
			InstanceType: "g4dn.xlarge",
			ImageFamily:  "AmazonLinux2",
			ExpectedAMI:  "ami-049574de1981b69dc",
			ExpectError:  false,
		}),
		Entry("with gpu (p3) instance and non-existent region", ResolveCase{
			Region:       "eu-east-1",
			Version:      "1.12",
			InstanceType: "p3.2xlarge",
			ImageFamily:  "AmazonLinux2",
			ExpectedAMI:  "",
			ExpectError:  true,
		}),
		Entry("with non-gpu instance, us-west-2 and Ubuntu image", ResolveCase{
			Region:       "us-west-2",
			Version:      "1.12",
			InstanceType: "t2.medium",
			ImageFamily:  "Ubuntu1804",
			ExpectedAMI:  "ami-0551b6dc0b4079e1d",
			ExpectError:  false,
		}),
		Entry("with non-gpu instance, us-east-1 and Ubuntu image", ResolveCase{
			Region:       "us-east-1",
			Version:      "1.12",
			InstanceType: "t2.medium",
			ImageFamily:  "Ubuntu1804",
			ExpectedAMI:  "ami-058c6e6f18b1cca2c",
			ExpectError:  false,
		}),
		Entry("with non-gpu instance, eu-west-1 and Ubuntu image", ResolveCase{
			Region:       "eu-west-1",
			Version:      "1.12",
			InstanceType: "t2.medium",
			ImageFamily:  "Ubuntu1804",
			ExpectedAMI:  "ami-0ac97cadeed862a54",
			ExpectError:  false,
		}),
		Entry("with non-gpu instance, non-eks enabled region and Ubuntu image", ResolveCase{
			Region:       "sa-east-1",
			Version:      "1.12",
			InstanceType: "t2.medium",
			ImageFamily:  "Ubuntu1804",
			ExpectedAMI:  "",
			ExpectError:  true,
		}),
		Entry("with gpu instance, any region and Ubuntu image", ResolveCase{
			Region:       "us-east-1",
			Version:      "1.12",
			InstanceType: "p2.xlarge",
			ImageFamily:  "Ubuntu1804",
			ExpectedAMI:  "",
			ExpectError:  true,
		}),
	)
})
