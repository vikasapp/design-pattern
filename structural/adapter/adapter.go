package adapter

func TestAdapter() {
	client := &Client{}
	apple := &Apple{}
	client.ChargePhone(apple)
	android := &Android{}
	androidAdapter := &AndroidAdapter{
		Android: android,
	}
	client.ChargePhone(androidAdapter)
}
