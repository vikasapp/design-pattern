package adapter

type Client struct{}

func (c *Client) ChargePhone(mobile IMobile) {
	mobile.ChargeAppleMobile()
}
