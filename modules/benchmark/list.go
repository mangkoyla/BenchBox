package benchmark

type BenchmarkListType struct {
	Name           string
	Domain         string
	Port           uint
	ExpectedStatus int
}

var BenchmarkList = []BenchmarkListType{
	{
		Name:           "TCP",
		Domain:         "https://cp.cloudflare.com/generate_204",
		Port:           443,
		ExpectedStatus: 204,
	},
	{
		Name:           "UDP",
		Domain:         "https://twilio.com",
		Port:           443,
		ExpectedStatus: 200,
	},
	{
		Name:           "Google",
		Domain:         "https://google.com",
		Port:           443,
		ExpectedStatus: 200,
	},
	{
		Name:           "Netflix",
		Domain:         "https://netflix.com",
		Port:           443,
		ExpectedStatus: 200,
	},
	{
		Name:           "Youtube",
		Domain:         "https://youtube.com",
		Port:           443,
		ExpectedStatus: 200,
	},
	{
		Name:           "Shopee",
		Domain:         "https://shopee.com",
		Port:           443,
		ExpectedStatus: 200,
	},
	{
		Name:           "Tokopedia",
		Domain:         "https://tokopedia.com",
		Port:           443,
		ExpectedStatus: 200,
	},
}
