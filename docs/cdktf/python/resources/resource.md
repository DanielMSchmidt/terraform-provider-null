---

<!-- Please do not edit this file, it is generated. -->
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "null_resource Resource - terraform-provider-null"
subcategory: ""
description: |-
  The null_resource resource implements the standard resource lifecycle but takes no further action. On Terraform 1.4 and later, use the terraform_data resource type https://developer.hashicorp.com/terraform/language/resources/terraform-data instead.
  The triggers argument allows specifying an arbitrary set of values that, when changed, will cause the resource to be replaced.
---

# null_resource

The `null_resource` resource implements the standard resource lifecycle but takes no further action. On Terraform 1.4 and later, use the [terraform_data resource type](https://developer.hashicorp.com/terraform/language/resources/terraform-data) instead.

The `triggers` argument allows specifying an arbitrary set of values that, when changed, will cause the resource to be replaced.

## Example Usage

```python
import constructs as constructs
import cdktf as cdktf
# Provider bindings are generated by running cdktf get.
# See https://cdk.tf/provider-generation for more details.
import ...gen.providers.aws as aws
import ...gen.providers.null as NullProvider
class MyConvertedCode(cdktf.TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        # In most cases loops should be handled in the programming language context and
        #     not inside of the Terraform context. If you are looping over something external, e.g. a variable or a file input
        #     you should consider using a for loop. If you are looping over something only known to Terraform, e.g. a result of a data source
        #     you need to keep this like it is.
        aws_instance_cluster_count = cdktf.TerraformCount.of(
            cdktf.Token.as_number("3"))
        aws_instance_cluster = aws.instance.Instance(self, "cluster",
            ami="ami-0dcc1e21636832c5d",
            instance_type="m5.large",
            count=aws_instance_cluster_count
        )
        null_provider_resource_cluster = NullProvider.resource.Resource(self, "cluster_1",
            connection=cdktf.SSHProvisionerConnection(
                host=cdktf.Fn.element(
                    cdktf.property_access(aws_instance_cluster, ["*", "public_ip"]), 0)
            ),
            triggers=[{
                "cluster_instance_ids": cdktf.Fn.join(",",
                    cdktf.Token.as_list(
                        cdktf.property_access(aws_instance_cluster, ["*", "id"])))
            }
            ],
            provisioners=[cdktf.FileProvisioner(
                type="remote-exec",
                inline=["bootstrap-cluster.sh " +
                    cdktf.Token.as_string(
                        cdktf.Fn.join(" ",
                            cdktf.Token.as_list(
                                cdktf.property_access(aws_instance_cluster, ["*", "private_ip"
                                ]))))
                ]
            )
            ]
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        null_provider_resource_cluster.override_logical_id("cluster")
```

<!-- schema generated by tfplugindocs -->

## Schema

### Optional

- `triggers` (Map of String) A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.

### Read-Only

- `id` (String) This is set to a random value at create time.

<!-- cache-key: cdktf-0.17.0-pre.15 input-64734d251545a8102312ebf4d3b3acb298b9d9b9070729db262d9ba29176dd3d -->