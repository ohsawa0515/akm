package main

type region struct {
	Name string
	Code string
}

var regions = []region{
	{Name: "None", Code: ""},
	{Name: "Asia Pacific (Tokyo)", Code: "ap-northeast-1"},
	{Name: "Asia Pacific (Seoul)", Code: "ap-northeast-2"},
	{Name: "Asia Pacific (Mumbai)", Code: "ap-south-1"},
	{Name: "Asia Pacific (Singapore)", Code: "ap-southeast-1"},
	{Name: "Asia Pacific (Sydney)", Code: "ap-southeast-2"},
	{Name: "Canada (Central)", Code: "ca-central-1"},
	{Name: "China (Beijing)", Code: "cn-north-1"},
	{Name: "China (Ningxia)", Code: "cn-northwest-1"},
	{Name: "EU (Frankfurt)", Code: "eu-central-1"},
	{Name: "EU (Ireland)", Code: "eu-west-1"},
	{Name: "EU (London)", Code: "eu-west-2"},
	{Name: "EU (Paris)", Code: "eu-west-3"},
	{Name: "South America (Sao Paulo)", Code: "sa-east-1"},
	{Name: "US East (N. Virginia)", Code: "us-east-1"},
	{Name: "US East (Ohio)", Code: "us-east-2"},
	{Name: "US West (N. California)", Code: "us-west-1"},
	{Name: "US West (Oregon)", Code: "us-west-2"},
}
