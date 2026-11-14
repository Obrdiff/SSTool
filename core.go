package main

import (
	"fmt"
	"time"
)

func typeWriter(text string) {
	for _, ch := range text {
		fmt.Printf("%c", ch)
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Println()
}

func fakeProgress(task string, steps int) {
	fmt.Printf("\n[*] %s\n", task)
	for i := 1; i <= steps; i++ {
		pct := (i * 100) / steps
		bar := ""
		for j := 0; j < 20; j++ {
			if j < (pct / 5) {
				bar += "█"
			} else {
				bar += "░"
			}
		}
		fmt.Printf("\r    [%s] %d%%", bar, pct)
		time.Sleep(80 * time.Millisecond)
	}
	fmt.Println(" ✓ Done!")
}

func main() {
	typeWriter("╔══════════════════════════════════╗")
	typeWriter("║   SUPER MEGA HACKER 3000  v0.1   ║")
	typeWriter("║      (totally not fake)           ║")
	typeWriter("╚══════════════════════════════════╝")

	time.Sleep(500 * time.Millisecond)

	fakeProgress("Scanning for WiFi networks...", 30)
	fakeProgress("Downloading more RAM...", 20)
	fakeProgress("Rerouting through Minecraft servers...", 40)
	fakeProgress("Asking nicely to enter the mainframe...", 25)
	fakeProgress("Compiling excuses...", 15)

	fmt.Println()
	typeWriter("⚠  WARNING: Coffee levels critically low")
	typeWriter("⚠  WARNING: Keyboard not found — press any key to continue")
	typeWriter("⚠  WARNING: DNS gremlins detected in sector 7G")

	time.Sleep(500 * time.Millisecond)

	fmt.Println()
	typeWriter("✅ Status: Successfully did absolutely nothing")
	typeWriter("💀 Systems: Still running somehow")
	typeWriter("🤝 Handshake: Completed with the void")
	fmt.Println()
	fmt.Println("Have a great day! (no systems were harmed)")
}
