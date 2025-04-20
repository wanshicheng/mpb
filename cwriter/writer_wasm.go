package cwriter

// WebAssembly环境下的Writer实现

// clearLines在WebAssembly环境中使用ANSI转义序列清除行
func (w *Writer) clearLines() error {
	// 在WebAssembly环境中，我们只能使用ANSI转义序列
	return w.ansiCuuAndEd()
}

// GetSize在WebAssembly环境中无法获取终端大小，返回默认值
func GetSize(fd int) (width, height int, err error) {
	// WebAssembly环境无法获取真实终端大小
	// 返回一个合理的默认值或零值
	return 0, 0, nil
}

// IsTerminal在WebAssembly环境中总是返回true
// 这样可以确保使用ANSI转义序列
func IsTerminal(fd int) bool {
	return true
}