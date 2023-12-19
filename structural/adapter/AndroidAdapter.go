package adapter

type AndroidAdapter struct {
	Android *Android
}

func (aa *AndroidAdapter) ChargeAppleMobile() {
	aa.Android.ChargeAndroidMobile()
}
