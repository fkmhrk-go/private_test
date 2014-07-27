package moke

func Hello() string {
	return internalHello("Moke")
}

// パッケージプライベートな関数
func internalHello(name string) string {
	return "My name is " + name
}
