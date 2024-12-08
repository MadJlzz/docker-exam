package main

func main() {
	cfg := NewAppConfiguration()

	sl := NewLogger(cfg.Logger)
	defer sl.Sync()

	sl.Info("hello, how are you? I am under the water")
}
